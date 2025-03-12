package main

import (
	"bytes"
	"context"
	pb "kvstore/proto"

	"fmt"
	"log"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	"google.golang.org/protobuf/proto"
)

type KVStoreApplication struct {
	balances      map[string]uint64
	currentHeight int
	pendingTxs    []QueuedTx
}

type QueuedTx struct {
	tx         []byte
	execHeight int
}

var _ abcitypes.Application = (*KVStoreApplication)(nil)

func NewKVStoreApplication() *KVStoreApplication {
    return &KVStoreApplication{
        balances: map[string]uint64{
			"Alice":   1000,
			"Bob":     500,
			"Clark":  300,
		},
    }
}

func (app *KVStoreApplication) Info(_ context.Context, info *abcitypes.InfoRequest) (*abcitypes.InfoResponse, error) {
    return &abcitypes.InfoResponse{}, nil
}

func (app *KVStoreApplication) Query(_ context.Context, req *abcitypes.QueryRequest) (*abcitypes.QueryResponse, error) {
	address := string(req.Data)
	balance := app.balances[address]
	return &abcitypes.QueryResponse{Value: []byte(fmt.Sprintf("%d", balance))}, nil
}

func (app *KVStoreApplication) CheckTx(_ context.Context, check *abcitypes.CheckTxRequest) (*abcitypes.CheckTxResponse, error) {
    return &abcitypes.CheckTxResponse{Code: 0}, nil
}

func (app *KVStoreApplication) InitChain(_ context.Context, chain *abcitypes.InitChainRequest) (*abcitypes.InitChainResponse, error) {
    return &abcitypes.InitChainResponse{}, nil
}

func (app *KVStoreApplication) PrepareProposal(_ context.Context, req *abcitypes.PrepareProposalRequest) (*abcitypes.PrepareProposalResponse, error) {
	var newTxs [][]byte

	for _, tx := range req.Txs {
        var decodedTx pb.Transaction
		if err := proto.Unmarshal(tx, &decodedTx); err != nil {
			log.Println("Failed to decode transaction:", err)
			continue
		}

        // Handle Transfer Transactions
		if transfer := decodedTx.GetTransfer(); transfer != nil {
			if transfer.Sender == "Alice" {
				// Check if Alice's transaction is already pending to avoid duplicate adds
				exists := false
				for _, pendingTx := range app.pendingTxs {
					if bytes.Equal(pendingTx.tx, tx) {
						exists = true
						break
					}
				}

				if !exists {
					log.Println("Delaying Alice's transaction, will execute at block %d\n", req.Height+2)
					app.pendingTxs = append(app.pendingTxs, QueuedTx{tx: tx, execHeight: int(req.Height + 2)})
				}
				continue
			}
		}

		newTxs = append(newTxs, tx)
	}

	var remainingQueue []QueuedTx
	for _, queuedTx := range app.pendingTxs {
		if int(req.Height) == queuedTx.execHeight {
			log.Println("Adding Alice's transaction %s, at block %d\n", queuedTx.tx, req.Height+2)
			newTxs = append(newTxs, queuedTx.tx)
		} else {
			remainingQueue = append(remainingQueue, queuedTx)
		}
	}
	app.pendingTxs = remainingQueue

	return &abcitypes.PrepareProposalResponse{Txs: newTxs}, nil
}

func (app *KVStoreApplication) ProcessProposal(_ context.Context, proposal *abcitypes.ProcessProposalRequest) (*abcitypes.ProcessProposalResponse, error) {
    return &abcitypes.ProcessProposalResponse{Status: abcitypes.PROCESS_PROPOSAL_STATUS_ACCEPT}, nil
}

func (app *KVStoreApplication) FinalizeBlock(_ context.Context, req *abcitypes.FinalizeBlockRequest) (*abcitypes.FinalizeBlockResponse, error) {
    var txs = make([]*abcitypes.ExecTxResult, 0)

    for _, tx := range req.Txs {
        var decodedTx pb.Transaction
		if err := proto.Unmarshal(tx, &decodedTx); err != nil {
			log.Println("Failed to decode transaction:", err)
			txs = append(txs, &abcitypes.ExecTxResult{Code: 1})
			continue
		}

		if transfer := decodedTx.GetTransfer(); transfer != nil {
			sender := transfer.Sender
			receiver := transfer.Receiver
			amount := transfer.Amount

			if app.balances[sender] < amount {
				log.Println("Insufficient balance for %s\n", sender)
				txs = append(txs, &abcitypes.ExecTxResult{Code: 1})
				continue
			}

			app.balances[sender] -= amount
			app.balances[receiver] += amount
			log.Println("%s sent %d tokens to %s\n", sender, amount, receiver)

			txs = append(txs, &abcitypes.ExecTxResult{Code: 0})
		}
    }

    return &abcitypes.FinalizeBlockResponse{
        TxResults:        txs,
    }, nil
}



func (app KVStoreApplication) Commit(_ context.Context, commit *abcitypes.CommitRequest) (*abcitypes.CommitResponse, error) {
    return &abcitypes.CommitResponse{}, nil
}

func (app *KVStoreApplication) ListSnapshots(_ context.Context, snapshots *abcitypes.ListSnapshotsRequest) (*abcitypes.ListSnapshotsResponse, error) {
    return &abcitypes.ListSnapshotsResponse{}, nil
}

func (app *KVStoreApplication) OfferSnapshot(_ context.Context, snapshot *abcitypes.OfferSnapshotRequest) (*abcitypes.OfferSnapshotResponse, error) {
    return &abcitypes.OfferSnapshotResponse{}, nil
}

func (app *KVStoreApplication) LoadSnapshotChunk(_ context.Context, chunk *abcitypes.LoadSnapshotChunkRequest) (*abcitypes.LoadSnapshotChunkResponse, error) {
    return &abcitypes.LoadSnapshotChunkResponse{}, nil
}

func (app *KVStoreApplication) ApplySnapshotChunk(_ context.Context, chunk *abcitypes.ApplySnapshotChunkRequest) (*abcitypes.ApplySnapshotChunkResponse, error) {
    return &abcitypes.ApplySnapshotChunkResponse{Result: abcitypes.APPLY_SNAPSHOT_CHUNK_RESULT_ACCEPT}, nil
}

func (app KVStoreApplication) ExtendVote(_ context.Context, extend *abcitypes.ExtendVoteRequest) (*abcitypes.ExtendVoteResponse, error) {
    return &abcitypes.ExtendVoteResponse{}, nil
}

func (app *KVStoreApplication) VerifyVoteExtension(_ context.Context, verify *abcitypes.VerifyVoteExtensionRequest) (*abcitypes.VerifyVoteExtensionResponse, error) {
    return &abcitypes.VerifyVoteExtensionResponse{}, nil
}

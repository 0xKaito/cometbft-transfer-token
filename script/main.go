package main

import (
	"fmt"
	pb "kvstore/proto"
	"google.golang.org/protobuf/proto"
	"net/http"
	"time"
	"encoding/hex"
)

var httpClient = &http.Client{Timeout: 60 * time.Second}

func main() {
	txBroadcastAddress := "http://localhost:26657"
	tx := &pb.Transaction{}
	tx.TxType = &pb.Transaction_Transfer{
		Transfer: &pb.TransferTransaction{
			Sender: "Alice",
			Receiver: "Bob",
			Amount: uint64(10),
			Timestamp: int64(time.Now().UnixMicro()),
		},
	}

	data, err := proto.Marshal(tx)

	if err != nil {
		fmt.Println(err)
	}
	hexData := hex.EncodeToString(data)

	fmt.Println("Broadcasting order ")
	res, err := httpClient.Get(fmt.Sprintf("%s/broadcast_tx_commit?tx=0x%s", txBroadcastAddress, hexData))

	if err != nil {
		fmt.Println("Error broadcasting order (trader/salt", err)
	}

	fmt.Println("res", res)
}
package main

import (
    "flag"
    "fmt"
    abciserver "github.com/cometbft/cometbft/abci/server"
    "os"
    "os/signal"
    "syscall"

    cmtlog "github.com/cometbft/cometbft/libs/log"
)

var socketAddr string

func init() {
    flag.StringVar(&socketAddr, "socket-addr", "unix://example.sock", "Unix domain socket address (if empty, uses \"unix://example.sock\"")
}

func main() {
    app := NewKVStoreApplication()
    logger := cmtlog.NewTMLogger(cmtlog.NewSyncWriter(os.Stdout))

    server := abciserver.NewSocketServer("tcp://127.0.0.1:26655", app)
    server.SetLogger(logger)

    if err := server.Start(); err != nil {
        fmt.Fprintf(os.Stderr, "error starting socket server: %v", err)

        os.Exit(1)
    }
    defer server.Stop()

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    <-c
}

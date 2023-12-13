package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joeldejesus1/wasm-go-1/web"
)

func main() {
	signalC := make(chan os.Signal, 1)
	signal.Notify(signalC, syscall.SIGTERM, syscall.SIGINT)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loopSignal(ctx, cancel, signalC)

	err := web.Run(ctx, &web.Configuration{
		Port: 8080,
	})
	if err != nil {
		panic(err)
	}
}

func loopSignal(ctx context.Context, cancel context.CancelFunc, signalC <-chan os.Signal) {
	defer cancel()
	doneC := ctx.Done()
	select {
	case <-doneC:
	case s := <-signalC:
		os.Stderr.WriteString(fmt.Sprintf("%s\n", s.String()))
	}
}

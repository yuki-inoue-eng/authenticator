package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yuki-inoue-eng/authenticator/configs"
	"github.com/yuki-inoue-eng/authenticator/server"
)

func main(){
	ctx, cancel := context.WithCancel(context.Background())
	cfgs, err := configs.New()
	if err != nil {
		panic(err)
	}

	eServer := server.New(ctx, cfgs)
	errC := make(chan error)
	go eServer.Start(errC)

	quitC := make(chan os.Signal, 1)
	signal.Notify(quitC, syscall.SIGTERM, os.Interrupt)

	select {
	case err := <-errC:
		panic(err)
	case <-quitC:
		if err := eServer.Shutdown(ctx); err != nil {
			errC <- err
		}
		cancel()
		time.Sleep(1 * time.Second)
	}
}
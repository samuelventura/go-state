package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	"github.com/samuelventura/go-state"
	"github.com/samuelventura/go-tree"
)

func run(launch func(tree.Node)) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	ctrlc := make(chan os.Signal, 1)
	signal.Notify(ctrlc, os.Interrupt)

	slog := &state.Log{}
	slog.Info = func(args ...interface{}) {
		log.Println("info", args)
	}
	slog.Warn = func(args ...interface{}) {
		log.Println("warn", args)
	}
	slog.Recover = func(ss string, args ...interface{}) {
		log.Println("recover", args, ss)
	}
	root := tree.NewRoot("root", &slog.Log)
	root.SetValue("log", slog)
	defer root.WaitDisposed()
	defer root.Recover()
	//async launcher must close root on error
	//and cleanup on root closed channel.
	go launch(root)

	stdin := make(chan interface{})
	go func() {
		defer close(stdin)
		ioutil.ReadAll(os.Stdin)
	}()
	select {
	case <-root.Closed():
	case <-ctrlc:
	case <-stdin:
	}
}

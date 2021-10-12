package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	"github.com/samuelventura/go-tree"
)

func run(cb func(tree.Node)) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	ctrlc := make(chan os.Signal, 1)
	signal.Notify(ctrlc, os.Interrupt)

	tl := &tree.Log{}
	tl.Warn = func(args ...interface{}) {
		log.Println("warn", args)
	}
	tl.Recover = func(ss string, args ...interface{}) {
		log.Println("recover", args, ss)
	}
	root := tree.NewRoot("root", tl)
	root.SetValue("log", log.Default())
	defer root.WaitDisposed()
	defer root.Recover()
	go cb(root)

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

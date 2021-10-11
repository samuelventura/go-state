package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	"github.com/samuelventura/go-state"
	"github.com/samuelventura/go-tree"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	ctrlc := make(chan os.Signal, 1)
	signal.Notify(ctrlc, os.Interrupt)

	http := tree.NewRoot("http", nil)
	mux := state.NewMux()
	state.AddPProfHandlers(mux)
	state.AddNodeHandlers(mux, http)
	state.AddEnvironHandlers(mux)
	http.SetValue("mux", mux)
	http.SetValue("path", "/tmp/go-state-to.sock")
	state.Serve(http)
	defer http.Close()

	stdin := make(chan interface{})
	go func() {
		defer close(stdin)
		ioutil.ReadAll(os.Stdin)
	}()
	select {
	case <-http.Closed():
	case <-ctrlc:
	case <-stdin:
	}
}

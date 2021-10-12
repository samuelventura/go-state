package main

import (
	"log"

	"github.com/samuelventura/go-state"
	"github.com/samuelventura/go-tree"
)

func main() {
	run(func(root tree.Node) {
		path := state.SingletonPath("/tmp")
		log := root.GetValue("log").(*log.Logger)
		log.Println("path", path)
		mux := state.NewMux()
		state.AddPProfHandlers(mux)
		state.AddNodeHandlers(mux, root)
		state.AddEnvironHandlers(mux)
		root.SetValue("mux", mux)
		root.SetValue("path", path)
		state.Serve(root) //ignore error
	})
}

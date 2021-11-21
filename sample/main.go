package main

import (
	"log"
	"os"

	"github.com/samuelventura/go-state"
	"github.com/samuelventura/go-tools"
	"github.com/samuelventura/go-tree"
)

func main() {
	tools.SetupLog()
	ctrlc := tools.SetupCtrlc()
	stdin := tools.SetupStdinWord("exit")

	rnode := tree.NewRoot("root", log.Println)
	defer rnode.WaitDisposed()
	//recover closes as well
	defer rnode.Recover()

	spath := state.SingletonPath()
	snode := state.Serve(rnode, spath)
	defer snode.WaitDisposed()
	defer snode.Close()
	slink := "/tmp/sample.state"
	os.Remove(slink) //ignore error
	err := os.Symlink(spath, slink)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("socket", spath)
	log.Println("link", slink)

	select {
	case <-rnode.Closed():
	case <-snode.Closed():
	case <-ctrlc:
	case <-stdin:
	}
}

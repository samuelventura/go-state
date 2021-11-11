package main

import (
	"bufio"
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

	stdin := make(chan interface{})
	go func() {
		defer close(stdin)
		//ioutil.ReadAll(os.Stdin)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == "exit" {
				return
			}
		}
	}()
	select {
	case <-rnode.Closed():
	case <-snode.Closed():
	case <-ctrlc:
	case <-stdin:
	}
}

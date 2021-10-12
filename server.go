package state

import (
	"net"
	"net/http"
	"os"

	"github.com/samuelventura/go-tree"
)

type Log struct {
	tree.Log
	Info func(args ...interface{})
}

func Serve(node tree.Node) error {
	mux := node.GetValue("mux").(Mux)
	path := node.GetValue("path").(string)
	log := node.GetValue("log").(*Log)
	os.Remove(path)
	//auto removed on close on macos
	listen, err := net.Listen("unix", path)
	if err != nil {
		node.Close()
		return err
	}
	node.AddCloser("listen", listen.Close)
	server := &http.Server{Handler: mux}
	node.AddProcess("serve", func() {
		err := server.Serve(listen)
		if err != nil {
			log.Warn(err)
		}
	})
	return nil
}

package state

import (
	"net"
	"net/http"
	"os"

	"github.com/samuelventura/go-tree"
)

func Serve(node tree.Node) error {
	mux := node.GetValue("mux").(Mux)
	path := node.GetValue("path").(string)
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
		server.Serve(listen)
	})
	return nil
}

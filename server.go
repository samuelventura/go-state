package state

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/samuelventura/go-tree"
)

func Serve(root tree.Node, path string) tree.Node {
	//auto removed on close on macos
	os.Remove(path) //ignore error
	listen, err := net.Listen("unix", path)
	if err != nil {
		log.Fatal(err)
	}

	slog := tree.NewLog()
	snode := tree.NewRoot("state", slog)
	smux := NewMux()
	AddPProfHandlers(smux)
	AddEnvironHandlers(smux)
	AddNodeHandlers(smux, snode)
	if root != nil {
		AddNodeHandlers(smux, root)
	}

	server := &http.Server{Handler: smux}
	snode.AddCloser("listen", listen.Close)
	snode.AddProcess("serve", func() {
		err := server.Serve(listen)
		if err != nil {
			log.Println(path, err)
		}
	})
	return snode
}

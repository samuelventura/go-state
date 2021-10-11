package state

import (
	"fmt"
	"net/http"

	"github.com/samuelventura/go-tree"
)

func AddNodeHandlers(mux Mux, node tree.Node) {
	path := fmt.Sprintf("/node/%s/", node.Name())
	handle := func(w http.ResponseWriter, s *tree.State, path string) {}
	handle = func(w http.ResponseWriter, s *tree.State, path string) {
		fmt.Fprintf(w, "<h1>Node /%s</h1>\n", path)
		fmt.Fprint(w, "<ul>\n")
		fmt.Fprintf(w, "<li>Closed %v\n", s.Closed)
		fmt.Fprintf(w, "<li>Disposed %v\n", s.Disposed)
		fmt.Fprintf(w, "<li>Actions %d\n", len(s.Actions))
		fmt.Fprint(w, "<ul>\n")
		for _, n := range s.Actions {
			fmt.Fprintf(w, "<li>%s\n", n)
		}
		fmt.Fprint(w, "</ul>\n")
		fmt.Fprintf(w, "<li>Agents %d\n", len(s.Agents))
		fmt.Fprint(w, "<ul>\n")
		for _, n := range s.Agents {
			fmt.Fprintf(w, "<li>%s\n", n)
		}
		fmt.Fprint(w, "</ul>\n")
		fmt.Fprintf(w, "<li>Children %d\n", len(s.Children))
		fmt.Fprint(w, "</ul>\n")
		for _, n := range s.Children {
			s = n.State()
			handle(w, s, path+"/"+s.Name)
		}
	}
	mux.AddHandler(path, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, "<html>\n")
		s := node.State()
		handle(w, s, s.Name)
		fmt.Fprint(w, "</html>\n")
	})
}

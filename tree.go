package state

import (
	"fmt"
	"net/http"

	"github.com/samuelventura/go-tree"
)

func AddNodeHandlers(mux Mux, node tree.Node) {
	path := fmt.Sprintf("/node/%s/", node.Name())
	handle := func(w http.ResponseWriter, s *tree.State, values bool, path string) {}
	handle = func(w http.ResponseWriter, s *tree.State, values bool, path string) {
		fmt.Fprintf(w, "%s\n", path)
		fmt.Fprintf(w, "- Closed %v\n", s.Closed)
		fmt.Fprintf(w, "- Disposed %v\n", s.Disposed)
		if values {
			fmt.Fprintf(w, "- Values %d\n", len(s.Values))
			for n, v := range s.Values {
				fmt.Fprintf(w, "-- %s=%s\n", n, v)
			}
		}
		fmt.Fprintf(w, "- Actions %d\n", len(s.Actions))
		for _, n := range s.Actions {
			fmt.Fprintf(w, "-- %s\n", n)
		}
		fmt.Fprintf(w, "- Agents %d\n", len(s.Agents))
		for _, n := range s.Agents {
			fmt.Fprintf(w, "-- %s\n", n)
		}
		fmt.Fprintf(w, "- Children %d\n", len(s.Children))
		for _, n := range s.Children {
			s = n.State()
			handle(w, s, values, path+"/"+s.Name)
		}
	}
	mux.AddHandler(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		values := r.URL.Query().Get("values") == "true"
		s := node.State()
		handle(w, s, values, s.Name)
	}))
}

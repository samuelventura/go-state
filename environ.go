package state

import (
	"fmt"
	"net/http"
	"os"
)

func AddEnvironHandlers(mux Mux) {
	mux.AddHandler("/environ/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		for _, n := range os.Environ() {
			fmt.Fprintf(w, "%s\n", n)
		}
	}))
}

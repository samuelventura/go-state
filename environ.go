package state

import (
	"fmt"
	"net/http"
	"os"
)

func AddEnvironHandlers(mux Mux) {
	mux.AddHandler("/environ/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, "<html>\n")
		for _, n := range os.Environ() {
			fmt.Fprintf(w, "<li>%s\n", n)
		}
		fmt.Fprint(w, "</html>\n")
	})
}

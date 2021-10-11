package state

import (
	"fmt"
	"net/http"
)

type Mux interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	AddHandler(path string, handler http.HandlerFunc)
}

type mux struct {
	index Sorted
}

func NewMux() Mux {
	mux := &mux{}
	mux.index = NewSorted()
	return mux
}

func (dso *mux) AddHandler(path string, handler http.HandlerFunc) {
	dso.index.Set(path, handler)
}

func (dso *mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.WriteHeader(200)
		fmt.Fprint(w, "<html>\n")
		fmt.Fprint(w, "<h1>Index</h1>\n")
		for _, n := range dso.index.Names() {
			fmt.Fprintf(w, "<a href='%s'>%s</a>\n", n, n)
		}
		fmt.Fprint(w, "</html>\n")
	} else {
		value := dso.index.Get(r.URL.Path)
		if value == nil {
			w.WriteHeader(404)
			fmt.Fprint(w, "NotFound 404")
		} else {
			handler := value.(http.HandlerFunc)
			handler(w, r)
		}
	}
}

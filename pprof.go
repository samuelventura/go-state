package state

import (
	"net/http"
	"net/http/pprof"
)

func AddPProfHandlers(mux Mux) {
	mux.AddHandler("/debug/pprof/", http.HandlerFunc(pprof.Index))
	mux.AddHandler("/debug/pprof/allocs", http.HandlerFunc(pprof.Cmdline))
	mux.AddHandler("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	mux.AddHandler("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	mux.AddHandler("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	mux.AddHandler("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	mux.AddHandler("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	mux.AddHandler("/debug/pprof/heap", pprof.Handler("heap"))
	mux.AddHandler("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	mux.AddHandler("/debug/pprof/block", pprof.Handler("block"))
}

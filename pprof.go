package state

import (
	"net/http"
	"net/http/pprof"
)

func AddPProfHandlers(mux Mux) {
	mux.AddHandler("/debug/pprof/", http.HandlerFunc(pprof.Index))
	mux.AddHandler("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	mux.AddHandler("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	mux.AddHandler("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	mux.AddHandler("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
}

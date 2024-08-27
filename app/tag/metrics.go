package main

import (
	"errors"
	"net/http"
	"sync/atomic"

	"github.com/VictoriaMetrics/metrics"
)

var (
	_live  atomic.Uint32
	_ready atomic.Uint32
)

func startMetrics() {
	setLiveOk()
	setReadyFail()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<h2>Server: core-pb/tag</h2><br/>
<a href="livez">livez</a> - liveness checking<br/>
<a href="readyz">readyz</a> - readiness checking<br/>
<a href="metrics">metrics</a> - available service metrics<br/>
`))
	})

	mux.HandleFunc("/metrics", func(w http.ResponseWriter, req *http.Request) {
		metrics.WritePrometheus(w, true)
	})

	mux.HandleFunc("/livez", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(int(_live.Load()))
	})
	mux.HandleFunc("/readyz", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(int(_ready.Load()))
	})

	go func() {
		if err := http.ListenAndServe(":80", mux); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
}

func setLiveOk()    { _live.Store(200) }
func setLiveFail()  { _live.Store(500) }
func setReadyOk()   { _ready.Store(200) }
func setReadyFail() { _ready.Store(500) }

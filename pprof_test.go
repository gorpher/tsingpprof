package tsingpprof_test

import (
	"github.com/dxvgef/tsing"
	"github.com/julienschmidt/httprouter"
	"github.com/wx11055/tsingpprof"
	"net/http"
	"testing"
	"time"
)

func newServer() *tsing.App {
	r := tsing.New()
	return r
}

var routers = map[string]string{
	"/debug/pprof/":             "IndexHandler",
	"/debug/pprof/heap":         "HeapHandler",
	"/debug/pprof/goroutine":    "GoroutineHandler",
	"/debug/pprof/allocs":       "AllocsHandler",
	"/debug/pprof/block":        "BlockHandler",
	"/debug/pprof/threadcreate": "ThreadCreateHandler",
	"/debug/pprof/cmdline":      "CmdlineHandler",
	"/debug/pprof/profile":      "ProfileHandler",
	"/debug/pprof/symbol":       "SymbolHandler",
	"/debug/pprof/trace":        "TraceHandler",
	"/debug/pprof/mutex":        "MutexHandler",
}

var url = "http://localhost:9999"

func TestRouters(t *testing.T) {
	httprouter.New()
	for k, v := range routers {
		t.Log(k, v)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Error(err)
			return
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err)
			return
		}
		if res.StatusCode != 200 {
			t.Errorf("missing router %s %s", k, v)
		}
	}
}

func TestWrap(t *testing.T) {
	app := newServer()
	tsingpprof.Wrap(app)
	go func() {
		t.Run("startup server", func(t *testing.T) {
			if err := http.ListenAndServe(":9999", app); err != nil {
				t.Error(err)
			}
		})
	}()

	t.Run("test routers ", func(t *testing.T) {
		time.Sleep(5 * time.Second)
		TestRouters(t)
	})
}

package tsingpprof_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dxvgef/tsing"
	"github.com/gorpher/tsingpprof"
)

func newServer() *tsing.App {
	r := tsing.New()
	return r
}

func TestWrap(t *testing.T) {
	app := newServer()
	tsingpprof.Wrap(app)

	req := httptest.NewRequest(http.MethodGet, "/debug/pprof", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("test failed : result code %d, want %d", w.Code, http.StatusOK)
	}
}

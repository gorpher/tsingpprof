package tsingpprof

import (
	"github.com/dxvgef/tsing"
	"net/http/pprof"
)

// 参考ginpprof : github.com/DeanThompson/ginpprof
// Wrap adds several routes from package `net/http/pprof` to *gin.Engine object
func Wrap(router *tsing.App) {
	WrapGroup(router.Router)
}

// Wrapper make sure we are backward compatible
var Wrapper = Wrap

// WrapGroup adds several routes from package `net/http/pprof` to *gin.RouterGroup object
func WrapGroup(router *tsing.RouterGroup) {
	routers := []struct {
		Method  string
		Path    string
		Handler tsing.Handler
	}{
		{"GET", "/debug/pprof/", IndexHandler()},
		{"GET", "/debug/pprof/heap", HeapHandler()},
		{"GET", "/debug/pprof/goroutine", GoroutineHandler()},
		{"GET", "/debug/pprof/allocs", AllocsHandler()},
		{"GET", "/debug/pprof/block", BlockHandler()},
		{"GET", "/debug/pprof/threadcreate", ThreadCreateHandler()},
		{"GET", "/debug/pprof/cmdline", CmdlineHandler()},
		{"GET", "/debug/pprof/profile", ProfileHandler()},
		{"GET", "/debug/pprof/symbol", SymbolHandler()},
		{"POST", "/debug/pprof/symbol", SymbolHandler()},
		{"GET", "/debug/pprof/trace", TraceHandler()},
		{"GET", "/debug/pprof/mutex", MutexHandler()},
	}

	//basePath := strings.TrimSuffix("", "/")
	//var prefix string
	//switch {
	//case basePath == "":
	//	prefix = ""
	//case strings.HasSuffix(basePath, "/debug"):
	//	prefix = "/debug"
	//case strings.HasSuffix(basePath, "/debug/pprof"):
	//	prefix = "/debug/pprof"
	//}

	for _, r := range routers {
		//router.Handle(r.Method, strings.TrimPrefix(r.Path, prefix), r.Handler)
		router.Handle(r.Method, r.Path, r.Handler)
	}
}

// IndexHandler will pass the call from /debug/pprof to pprof
func IndexHandler() tsing.Handler {
	return func(ctx *tsing.Context) error {
		pprof.Index(ctx.ResponseWriter, ctx.Request)
		return nil
	}
}

// HeapHandler will pass the call from /debug/pprof/heap to pprof
func HeapHandler() tsing.Handler {
	return func(ctx *tsing.Context) error {
		pprof.Handler("heap").ServeHTTP(ctx.ResponseWriter, ctx.Request)
		return nil
	}
}

// GoroutineHandler will pass the call from /debug/pprof/goroutine to pprof
func GoroutineHandler() tsing.Handler {
	return func(ctx *tsing.Context) error {
		pprof.Handler("goroutine").ServeHTTP(ctx.ResponseWriter, ctx.Request)
		return nil
	}
}

// AllocsHandler will pass the call from /debug/pprof/allocs to pprof
func AllocsHandler() tsing.Handler {
	return func(ctx *tsing.Context) error {
		pprof.Handler("allocs").ServeHTTP(ctx.ResponseWriter, ctx.Request)
		return nil
	}
}

// BlockHandler will pass the call from /debug/pprof/block to pprof
func BlockHandler() tsing.Handler {
	return func(ctx *tsing.Context) error {
		pprof.Handler("block").ServeHTTP(ctx.ResponseWriter, ctx.Request)
		return nil
	}
}

// ThreadCreateHandler will pass the call from /debug/pprof/threadcreate to pprof
func ThreadCreateHandler() tsing.Handler {
	return func(ctx *tsing.Context) error {
		pprof.Handler("threadcreate").ServeHTTP(ctx.ResponseWriter, ctx.Request)
		return nil
	}
}

// CmdlineHandler will pass the call from /debug/pprof/cmdline to pprof
func CmdlineHandler() tsing.Handler {
	return func(ctx *tsing.Context) error {
		pprof.Cmdline(ctx.ResponseWriter, ctx.Request)
		return nil
	}
}

// ProfileHandler will pass the call from /debug/pprof/profile to pprof
func ProfileHandler() tsing.Handler {
	return func(ctx *tsing.Context) error {
		pprof.Profile(ctx.ResponseWriter, ctx.Request)
		return nil
	}
}

// SymbolHandler will pass the call from /debug/pprof/symbol to pprof
func SymbolHandler() tsing.Handler {
	return func(ctx *tsing.Context) error {
		pprof.Symbol(ctx.ResponseWriter, ctx.Request)
		return nil
	}
}

// TraceHandler will pass the call from /debug/pprof/trace to pprof
func TraceHandler() tsing.Handler {
	return func(ctx *tsing.Context) error {
		pprof.Trace(ctx.ResponseWriter, ctx.Request)
		return nil
	}
}

// MutexHandler will pass the call from /debug/pprof/mutex to pprof
func MutexHandler() tsing.Handler {
	return func(ctx *tsing.Context) error {
		pprof.Handler("mutex").ServeHTTP(ctx.ResponseWriter, ctx.Request)
		return nil
	}
}

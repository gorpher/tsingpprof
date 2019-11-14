package tsingpprof

import (
	"net/http/pprof"

	"github.com/dxvgef/tsing"
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
	router = router.GROUP("/debug/pprof")
	router.Handle("GET", "/", IndexHandler())
	router.Handle("GET", "/heap", HeapHandler())
	router.Handle("GET", "/goroutine", GoroutineHandler())
	router.Handle("GET", "/allocs", AllocsHandler())
	router.Handle("GET", "/block", BlockHandler())
	router.Handle("GET", "/threadcreate", ThreadCreateHandler())
	router.Handle("GET", "/cmdline", CmdlineHandler())
	router.Handle("GET", "/profile", ProfileHandler())
	router.Handle("GET", "/symbol", SymbolHandler())
	router.Handle("POST", "/symbol", SymbolHandler())
	router.Handle("GET", "/trace", TraceHandler())
	router.Handle("GET", "/mutex", MutexHandler())
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

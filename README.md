### tsingpprof
[![GoDoc](https://godoc.org/github.com/gorpher/tsingpprof?status.svg)](https://godoc.org/github.com/gorpher/tsingpprof)
[![Build Status](https://travis-ci.org/gorpher/tsingpprof.svg?branch=master)](https://travis-ci.org/gorpher/tsingpprof)

A wrapper for golang web framework `tsing` to use   `net/http/pprof` easily.


### Install
First install tsingpprof to your GOPATH using go get:

> go get github.com/DeanThompson/tsingpprof

### usage
```go
package main

import (
	"log"
	"net/http"

	"github.com/dxvgef/tsing"
	"github.com/gorpher/tsingpprof"

)

func main()  {
	//获得框架实例
	var app = tsing.New()

    // 添加tsingpprof
    tsingpprof.Wrap(app)

	//注册路由
	app.Router.GET("/", func(ctx *tsing.Context) error {
		ctx.ResponseWriter.WriteHeader(200)
		_, err := ctx.ResponseWriter.Write([]byte("Hello world"))
        		//路由函数的出参为nil，表示路由控制器正常执行完毕
        return err
	})

	if err := http.ListenAndServe(":8080", app); err != nil {
		log.Fatal(err.Error())
	}
}

```

### test
```
=== RUN   TestWrap/test_routers_
    --- PASS: TestWrap/test_routers_ (5.02s)
        pprof_test.go:36: /debug/pprof/profile ProfileHandler
        pprof_test.go:36: /debug/pprof/trace TraceHandler
        pprof_test.go:36: /debug/pprof/goroutine GoroutineHandler
        pprof_test.go:36: /debug/pprof/heap HeapHandler
        pprof_test.go:36: /debug/pprof/allocs AllocsHandler
        pprof_test.go:36: /debug/pprof/block BlockHandler
        pprof_test.go:36: /debug/pprof/threadcreate ThreadCreateHandler
        pprof_test.go:36: /debug/pprof/cmdline CmdlineHandler
        pprof_test.go:36: /debug/pprof/symbol SymbolHandler
        pprof_test.go:36: /debug/pprof/mutex MutexHandler
        pprof_test.go:36: /debug/pprof/ IndexHandler
```
Now visit http://localhost:9999/debug/pprof/ and you'll see what you want.

Have Fun.

### reference
Thanks for free **JetBrains** Open Source license

<a href="https://www.jetbrains.com/?from=tsingpprof" target="_blank">
<img src="https://user-images.githubusercontent.com/1787798/69898077-4f4e3d00-138f-11ea-81f9-96fb7c49da89.png" height="200"/></a>
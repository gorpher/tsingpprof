package main

import (
	"log"
	"net/http"

	"github.com/dxvgef/tsing"
	"github.com/gorpher/tsingpprof"
)

func main() {
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

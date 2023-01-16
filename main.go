// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/qimengxingyuan/young_engine/biz/dal"
	"github.com/qimengxingyuan/young_engine/biz/handler"
)

func main() {
	Init()
	r := server.Default()

	r.GET("/ping", handler.Ping)

	g := r.Group("/api")
	g.POST("/engine/run", handler.HandleRunRule)
	g.POST("/engine/exp/new", handler.HandleAddExpression)
	g.GET("/engine/exp/list", handler.HandleGetAllExpression)
	g.DELETE("/engine/exp/:id", handler.HandleDeleteExpression)
	g.POST("/engine/exp/run", handler.HandleRunExpression)

	r.Spin()
}

func Init() {
	dal.InitDB()
}

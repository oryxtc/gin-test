package main

import (
	"vss/app"
	_ "vss/router" //初始化路由
)

func main() {
	engine := app.GetEngine()
	engine.Run(":8080")
}

package main

import (
	//_ "github.com/ennm/go-frame/boot"
	_ "github.com/ennm/go-frame/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}

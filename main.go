package main

import (
	_ "cli_gf_proj/boot"
	_ "cli_gf_proj/router"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	// g.Server().Run()

	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
        r.Response.Write("哈喽世界！")
    })

	s.Run()
}

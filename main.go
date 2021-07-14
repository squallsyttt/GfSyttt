package main

import (
	_ "cli_gf_proj/boot"
	_ "cli_gf_proj/router"

	"github.com/gogf/gf/frame/g"
)

func main() {

	// g.Server()方法获取到一个默认的 Server对象，该方法采用单例模式设计
	s := g.Server()
	// s.BindHandler("/", func(r *ghttp.Request) {
    //     r.Response.Write("goframe bindHandler！")
    // })

	s.SetIndexFolder(true)
	s.SetServerRoot("/usr/local/www/go_proj/cli_gf_proj/public")


	s.Run()
}

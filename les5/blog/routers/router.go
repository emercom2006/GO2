package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.BlogPostsController{})
	beego.Router("/post/:id([0-9]+)", &controllers.BlogPostsController{}, "get:GetOnePost")
}

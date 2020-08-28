package controllers

import (
	"blog/models"
	"log"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BlogPostsController struct {
	beego.Controller
}

func (c *BlogPostsController) Get() {
	beeOrm := orm.NewOrm()

	posts := models.Posts{}

	_, err := beeOrm.QueryTable("tablename").All(&posts)
	if err != nil {
		log.Fatal(err)
	}

	c.Data["Posts"] = posts
	c.Data["Title"] = posts.Title
	c.TplName = "index.tpl"
}

func (c *BlogPostsController) GetOnePost() {
	id := c.Ctx.Input.Param(":id")
	uint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte("Post id is incorrect"))
		return
	}

	beeOrm := orm.NewOrm()
	posts := models.Posts{Id: uint64}
	err = beeOrm.QueryTable("tablename").Filter("Id", uint64).One(&posts)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte("Post id is incorrect"))
		return
	}

	c.Data["Title"] = posts.Title
	c.Data["Short"] = posts.Short
	c.Data["Long"] = posts.Long
	c.TplName = "post.tpl"
}

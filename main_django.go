package main

import (
	"github.com/iris-contrib/template/django"
	"github.com/kataras/iris"
)

func main() {
	iris.UseTemplate(django.New()).Directory("./templates/django", ".html")
	iris.Get("/hi", hi)
	iris.Listen(":8080")
}

func hi(ctx *iris.Context) {
	ctx.Render("hi.html", map[string]interface{}{"Name": "iris"}, iris.RenderOptions{"gzip": true})
}

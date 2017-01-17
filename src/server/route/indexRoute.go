package route

import (
	"github.com/kataras/iris"
	"goshop/src/server/entity"
)

func setUpIndexRoute() {

	iris.Get("/", Index)
	iris.Get("/index/filter", indexFilter)
}

func Index(ctx *iris.Context) {

	ctx.Render("index.html", nil)
}

func indexFilter(ctx *iris.Context) {
	toUrl := ctx.URLParam("toUrl")
	ctx.Session().Set("toUrl", toUrl)
	ctx.JSON(iris.StatusOK, entity.Result{Status:0, Data:"", Msg:"success"})
}
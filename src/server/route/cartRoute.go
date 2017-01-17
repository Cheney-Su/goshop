package route

import (
	"github.com/kataras/iris"
	"goshop/src/server/entity"
	"goshop/src/server/service"
)

func setUpCartRoute() {

	iris.Post("/cart/add/:pid", addCart)
	iris.Get("/cart/delete/:pid", deleteCart)
	iris.Get("/cart/list/", listCart)
	iris.Get("/cart/clear", clearCart)
}

var (
	cartService = &service.CartService{}
)

func listCart(ctx *iris.Context) {
	sessionId := ctx.Session().ID()
	var userId string
	userId = ctx.Session().GetString("userId")
	cart := cartService.ListCart(userId, sessionId)
	ctx.JSON(iris.StatusOK, entity.Result{Status:0, Data:cart, Msg:"success"})
}

func addCart(ctx *iris.Context) {
	pid := ctx.Param("pid")
	count := ctx.URLParam("count")
	sessionId := ctx.Session().ID()
	if len(pid) == 0 || pid == "null" || len(count) == 0 || count == "null" {
		ctx.JSON(iris.StatusOK, entity.Result{Status:1005, Data:"", Msg:"参数非法"})
		return
	}
	var userId string
	userId = ctx.Session().GetString("userId")
	cart := cartService.AddCart(userId, sessionId, pid, count)
	ctx.JSON(iris.StatusOK, entity.Result{Status:0, Data:cart, Msg:"success"})
}

func deleteCart(ctx *iris.Context) {
	pid := ctx.Param("pid")
	sessionId := ctx.Session().ID()
	if len(pid) == 0 {
		ctx.JSON(iris.StatusOK, entity.Result{Status:1005, Data:"", Msg:"参数非法"})
		return
	}
	var userId string
	userId = ctx.Session().GetString("userId")
	cart := cartService.DeleteCart(userId, sessionId, pid)
	ctx.JSON(iris.StatusOK, entity.Result{Status:0, Data:cart, Msg:"success"})
}

func clearCart(ctx *iris.Context) {
	sessionId := ctx.Session().ID()
	var userId string
	userId = ctx.Session().GetString("userId")
	cart := cartService.ClearCart(userId, sessionId)
	ctx.JSON(iris.StatusOK, entity.Result{Status:0, Data:cart, Msg:"success"})
}
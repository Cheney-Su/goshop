package route

import (
	"github.com/kataras/iris"
	"goshop/src/server/entity"
	"goshop/src/server/manager/route"
)

func SetUpRoute() {
	setUpIndexRoute()
	setUpCategoryRoute()
	setUpProductRoute()
	setUpCartRoute()
	SetUpUserRoute()
	SetUpOrderRoute()

	//管理系统
	route.SetUpManagerRoute()

	iris.Listen(":8088")
}

func CheckLogin(ctx *iris.Context) {
	if len(ctx.Session().GetString("userId")) <= 0 {
		ctx.JSON(iris.StatusOK, entity.Result{Status:6003, Data:"", Msg:"亲，请登录后再进行提交订到操作"})
		return
	}
	ctx.Next()
}
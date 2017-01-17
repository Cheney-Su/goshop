package route

import (
	"github.com/kataras/iris"
	"go/src/server/entity"
	"goshop/src/server/service"
)

var (
	userService = &service.UserService{}
)

func SetUpUserRoute() {

	iris.Post("/user/login", userLogin)
	iris.Get("/user/userInfo", getUserInfoByUid)
}

func userLogin(ctx *iris.Context) {
	username := ctx.FormValueString("username")
	password := ctx.FormValueString("password")
	if username == "" || password == "" {
		ctx.JSON(iris.StatusOK, entity.Result{Status:1005, Data:"", Msg:"参数非法"})
		return
	}
	userLoginResult, toUrl := userService.UserLogin(username, password, ctx)
	if userLoginResult {
		ctx.JSON(iris.StatusOK, entity.Result{Status:0, Data:toUrl, Msg:"success"})
		return
	}
	ctx.JSON(iris.StatusOK, entity.Result{Status:-1, Data:toUrl, Msg:"error"})
}

func getUserInfoByUid(ctx *iris.Context) {
	userId := ctx.Session().GetString("userId")
	userInfo := userService.GetUserInfoByUid(userId)
	if userInfo.Uid == 0 {
		ctx.JSON(iris.StatusOK, entity.Result{Status:6003, Data:"", Msg:"未找到数据..."})
		return
	}
	ctx.JSON(iris.StatusOK, entity.Result{Status:0, Data:userInfo, Msg:"success"})
}
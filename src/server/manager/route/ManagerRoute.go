package route

import (
	"github.com/kataras/iris"
	"goshop/src/server/manager/service"
	"goshop/src/server/entity"
	managerEntity "goshop/src/server/manager/entity"
	"fmt"
)

var (
	managerService = service.ManagerService{}
)

type data struct {
	manager managerEntity.Manager `json:"manager"`
	menu managerEntity.Menu	`json:"menu"`
}

func SetUpManagerRoute() {
	iris.Get("/manager/login", managerLogin)
}

func managerLogin(ctx *iris.Context) {
	username := ctx.URLParam("username")
	password := ctx.URLParam("password")
	if username == "" || password == "" {
		ctx.JSON(iris.StatusOK,entity.Result{Status:1005, Data:"", Msg:"参数非法"})
		return
	}
	manager, menu := managerService.ManagerLogin(username, password)
	d := data{manager:manager,menu:menu}
	fmt.Println(d)
	ctx.JSON(iris.StatusOK, entity.Result{Status:0, Data:d, Msg:"success"})
}

package route

import (
	"goshop/src/server/service"
	"github.com/kataras/iris"
	"goshop/src/server/entity"
)

var (
	orderService = service.OrderService{}
)

func SetUpOrderRoute() {
	iris.Get("/order/add", addOrder)
	iris.Get("/order/uid/:uid", CheckLogin, getOrderByUid)
	iris.Get("/order/oid/:oid", getOrderByOid)
	iris.Post("/order/pay/:oid", payOrder)
	iris.Get("/order/payBack", callBackPayOrder)
}

func addOrder(ctx *iris.Context) {
	addOrderFlag := orderService.AddOrder(ctx)
	if addOrderFlag {
		ctx.JSON(iris.StatusOK, entity.Result{Status:0, Data:"", Msg:"success"})
		return
	}
	ctx.JSON(iris.StatusOK, entity.Result{Status:-1, Data:"", Msg:"系统异常..."})
}

func getOrderByUid(ctx *iris.Context) {
	uid, _ := ctx.ParamInt("uid")
	page, _ := ctx.URLParamInt("page")
	pageSize, _ := ctx.URLParamInt("pageSize")
	if uid == 0 {
		ctx.JSON(iris.StatusOK, entity.Result{1005, "", "参数非法"})
		return
	}
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}
	orders, total := orderService.GetOrderByUid(uid, page, pageSize)
	ctx.JSON(iris.StatusOK, entity.View{Status:0, Data:orders, Msg:"success", Total:total})
}

func getOrderByOid(ctx *iris.Context) {
	oid, _ := ctx.ParamInt("oid")
	if oid == 0 {
		ctx.JSON(iris.StatusOK, entity.Result{1005, "", "参数非法"})
		return
	}
	order := orderService.GetOrderByOid(oid)
	ctx.JSON(iris.StatusOK, entity.Result{Status:0, Data:order, Msg:"success"})
}

func payOrder(ctx *iris.Context) {
	oid := ctx.Param("oid")
	addr := ctx.FormValueString("addr")
	name := ctx.FormValueString("name")
	phone := ctx.FormValueString("phone")
	bank := ctx.FormValueString("pd_FrpId")
	//更新订单详情
	orderService.UpdateOrderInfo(oid, addr, name, phone)
	//支付订单
	yibaoUrl := orderService.PayOrder(oid, bank)
	ctx.JSON(iris.StatusOK, entity.Result{Status:0, Data:yibaoUrl, Msg:"success"})
}

func callBackPayOrder(ctx *iris.Context) {
	oid := ctx.URLParam("r6_Order")
	orderService.CallBackPayOrder(oid)
	ctx.Render("orderList.html",nil)
}
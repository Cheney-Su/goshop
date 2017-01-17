package route

import (
	"github.com/kataras/iris"
	"goshop/src/server/service"
	"goshop/src/server/entity"
)

var (
	productService = service.ProductService{}
)

func setUpProductRoute() {

	iris.Get("/product/cid/:cid", getProductByCid)
	iris.Get("/product/csid/:csid", getProductByCsid)
	iris.Get("/product/pid/:pid", getProductByPid)
	iris.Get("/product/hot/", getProductHotList)
	iris.Get("/product/new/", getProductNewList)
}

func getProductHotList(ctx *iris.Context) {
	hotProduct, total := productService.GetProductHotList()
	ctx.JSON(iris.StatusOK, entity.View{Status:0, Data:hotProduct, Msg:"success", Total:total})
}

func getProductNewList(ctx *iris.Context) {
	newProduct, total := productService.GetProductNewList()
	ctx.JSON(iris.StatusOK, entity.View{Status:0, Data:newProduct, Msg:"success", Total:total})
}

func getProductByCid(ctx *iris.Context) {
	cid, _ := ctx.ParamInt("cid")
	page, _ := ctx.URLParamInt("page")
	pageSize, _ := ctx.URLParamInt("pageSize")
	if cid == 0 {
		ctx.JSON(iris.StatusOK, entity.Result{1005, "", "参数非法"})
		return
	}
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}
	productByCid, total := productService.GetProductByCid(cid, page, pageSize)
	//ctx.Render("productList.html", productByCid)
	ctx.JSON(iris.StatusOK, entity.View{Status:0, Data:productByCid, Msg:"success", Total:total})
}

func getProductByCsid(ctx *iris.Context) {
	csid, _ := ctx.ParamInt("csid")
	page, _ := ctx.URLParamInt("page")
	pageSize, _ := ctx.URLParamInt("pageSize")
	if csid == 0 {
		ctx.JSON(iris.StatusOK, entity.Result{1005, "", "参数非法"})
		return
	}
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}
	productByCsid, total := productService.GetProductByCsid(csid, page, pageSize)
	//ctx.Render("productList.html", productByCid)
	ctx.JSON(iris.StatusOK, entity.View{Status:0, Data:productByCsid, Msg:"success", Total:total})
}

func getProductByPid(ctx *iris.Context) {
	pid, _ := ctx.ParamInt("pid")
	if pid == 0 {
		ctx.JSON(iris.StatusOK, entity.Result{1005, "", "参数非法"})
		return
	}
	productByPid := productService.GetProductByPid(pid)
	ctx.JSON(iris.StatusOK, entity.Result{Status:0, Data:productByPid, Msg:"success"})
}

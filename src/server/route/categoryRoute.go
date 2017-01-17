package route

import (
	"github.com/kataras/iris"
	"goshop/src/server/service"
	"goshop/src/server/entity"
)

var (
	categoryService = &service.CategoryService{}
)

func setUpCategoryRoute() {

	iris.Get("/category/", getCategoryList)
	iris.Get("/category/categorysecond/", getCategorySecondList)
}

func getCategoryList(ctx *iris.Context) {
	category, total := categoryService.GetCategoryList()
	ctx.JSON(iris.StatusOK, entity.View{Status:0, Data:category, Msg:"success", Total:total})
	//ctx.JSON(iris.StatusOK, "你好")
}

func getCategorySecondList(ctx *iris.Context) {
	secondcategory, total := categoryService.GetCategorySecondList()
	ctx.JSON(iris.StatusOK, entity.View{Status:0, Data:secondcategory, Msg:"success", Total:total})
	//ctx.JSON(iris.StatusOK, "你好")
}



package dao

import (
	"goshop/src/server/entity"
	"goshop/src/server/config"
	"goshop/src/server/utils"
)

type CategoryDao struct {

}

func (dao *CategoryDao) GetCategoryList() ([]entity.Category,int) {
	//函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。
	sql := "select cid,cname from category"
	rows, err := config.ShopDB.Query(sql)
	utils.DelError(err)
	//创建一个切片
	category := make([]entity.Category, 0)
	for rows.Next() {
		//创建一个实体
		entity := entity.Category{}
		rows.Scan(&entity.Cid, &entity.Cname)
		//增长切片的长度
		category = append(category, entity)
	}
	return category,len(category)
}

func (dao *CategoryDao) GetCategorySecondList() ([]entity.Category,int) {
	categoryEntity,total := dao.GetCategoryList()
	//函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。
	sql := "select csid,csname from categorysecond where cid = ?"
	for i := 0; i < len(categoryEntity); i++ {
		rows, err := config.ShopDB.Query(sql, categoryEntity[i].Cid)
		utils.DelError(err)
		//创建一个切片
		categorySeconds := make([]entity.CategorySecond, 0)
		for rows.Next() {
			//创建一个实体
			secondEntity := entity.CategorySecond{}
			rows.Scan(&secondEntity.Csid, &secondEntity.Csname)
			//增长切片的长度
			categorySeconds = append(categorySeconds, secondEntity)
		}
		categoryEntity[i].CategorySeconds = categorySeconds
	}
	return categoryEntity,total
}

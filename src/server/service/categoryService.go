package service

import (
	"goshop/src/server/dao"
	"goshop/src/server/entity"
)

var (
	categoryDao = &dao.CategoryDao{}
)

type CategoryService struct {

}

func (CategoryService) GetCategoryList() ([]entity.Category, int) {
	return categoryDao.GetCategoryList()
}

func (CategoryService) GetCategorySecondList() ([]entity.Category, int) {
	return categoryDao.GetCategorySecondList()
}




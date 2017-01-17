package service

import (
	"goshop/src/server/entity"
	"goshop/src/server/dao"
)

var (
	productService = &dao.ProductDao{}
)

type ProductService struct {

}

func (ProductService) GetProductHotList() ([]entity.Product, int) {
	return productService.GetProductHotList()
}

func (ProductService) GetProductNewList() ([]entity.Product, int) {
	return productService.GetProductNewList()
}

func (ProductService) GetProductByCid(id, page, pageSize int) ([]entity.Product, int) {
	start := (page - 1) * pageSize
	end := pageSize
	return productService.GetProductByCid(id, start, end), productService.GetProductTotalByCid(id)
}

func (ProductService) GetProductByCsid(id, page, pageSize int) ([]entity.Product, int) {
	start := (page - 1) * pageSize
	end := pageSize
	return productService.GetProductByCsid(id, start, end), productService.GetProductTotalByCsid(id)
}

func (ProductService) GetProductByPid(id int) entity.Product {
	return productService.GetProductByPid(id)
}

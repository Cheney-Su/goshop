package dao

import (
	"goshop/src/server/entity"
	"goshop/src/server/utils"
	"goshop/src/server/config"
)

type ProductDao struct {

}

func (ProductDao) GetProductHotList() ([]entity.Product, int) {
	sql := "select p.pid,p.pname,p.market_price,p.shop_price,p.image,p.pdesc,p.is_hot,pdate,cs.csid,cs.csname,c.cid,c.cname " +
		"from product p,categorysecond cs,category c " +
		"where p.csid = cs.csid and cs.cid = c.cid and is_hot = 1 order by pdate desc limit 10";
	rows, err := config.ShopDB.Query(sql)
	utils.DelError(err)
	//创建切片
	products := make([]entity.Product, 0)
	for rows.Next() {
		entity := entity.Product{}
		//categorySecond := entity.CategorySecond{}
		rows.Scan(&entity.Pid, &entity.Pname, &entity.MarketPrice, &entity.ShopPrice, &entity.Image, &entity.Pdesc,
			&entity.IsHot, &entity.Pdate, &entity.CategorySecond.Csid, &entity.CategorySecond.Csname,
			&entity.CategorySecond.Category.Cid, &entity.CategorySecond.Category.Cname)
		products = append(products, entity)
	}
	return products, len(products)
}

func (ProductDao) GetProductNewList() ([]entity.Product, int) {
	sql := "select p.pid,p.pname,p.market_price,p.shop_price,p.image,p.pdesc,p.is_hot,pdate,cs.csid,cs.csname,c.cid,c.cname " +
		"from product p,categorysecond cs,category c " +
		"where p.csid = cs.csid and cs.cid = c.cid order by pdate desc limit 10";
	rows, err := config.ShopDB.Query(sql)
	utils.DelError(err)
	//创建切片
	products := make([]entity.Product, 0)
	for rows.Next() {
		entity := entity.Product{}
		//categorySecond := entity.CategorySecond{}
		rows.Scan(&entity.Pid, &entity.Pname, &entity.MarketPrice, &entity.ShopPrice, &entity.Image, &entity.Pdesc,
			&entity.IsHot, &entity.Pdate, &entity.CategorySecond.Csid, &entity.CategorySecond.Csname,
			&entity.CategorySecond.Category.Cid, &entity.CategorySecond.Category.Cname)
		products = append(products, entity)
	}
	return products, len(products)
}

func (ProductDao) GetProductByCid(id, start, end int) []entity.Product {
	sql := "select p.pid,p.pname,p.market_price,p.shop_price,p.image,p.pdesc,p.is_hot,pdate,cs.csid,cs.csname,c.cid,c.cname " +
		"from product p,categorysecond cs,category c " +
		"where c.cid = ? and p.csid = cs.csid and cs.cid = c.cid order by pdate desc limit ?,?";
	rows, err := config.ShopDB.Query(sql, id, start, end)
	utils.DelError(err)
	//创建切片
	products := make([]entity.Product, 0)
	for rows.Next() {
		entity := entity.Product{}
		//categorySecond := entity.CategorySecond{}
		rows.Scan(&entity.Pid, &entity.Pname, &entity.MarketPrice, &entity.ShopPrice, &entity.Image, &entity.Pdesc,
			&entity.IsHot, &entity.Pdate, &entity.CategorySecond.Csid, &entity.CategorySecond.Csname,
			&entity.CategorySecond.Category.Cid, &entity.CategorySecond.Category.Cname)
		products = append(products, entity)
	}
	return products
}

func (ProductDao) GetProductTotalByCid(id int) int {
	sql := "select count(*) count from product p,categorysecond cs,category c " +
		"where c.cid = ? and p.csid = cs.csid and cs.cid = c.cid order by pdate desc";
	rows, err := config.ShopDB.Query(sql, id)
	utils.DelError(err)
	var total int
	for rows.Next() {
		rows.Scan(&total)
	}
	return total
}

func (ProductDao) GetProductByCsid(id, start, end int) []entity.Product {
	sql := "select p.pid,p.pname,p.market_price,p.shop_price,p.image,p.pdesc,p.is_hot,pdate,cs.csid,cs.csname,c.cid,c.cname " +
		"from product p,categorysecond cs,category c " +
		"where cs.csid = ? and p.csid = cs.csid and cs.cid = c.cid order by pdate desc limit ?,?";
	rows, err := config.ShopDB.Query(sql, id, start, end)
	utils.DelError(err)
	//创建切片
	products := make([]entity.Product, 0)
	for rows.Next() {
		entity := entity.Product{}
		//categorySecond := entity.CategorySecond{}
		rows.Scan(&entity.Pid, &entity.Pname, &entity.MarketPrice, &entity.ShopPrice, &entity.Image, &entity.Pdesc,
			&entity.IsHot, &entity.Pdate, &entity.CategorySecond.Csid, &entity.CategorySecond.Csname,
			&entity.CategorySecond.Category.Cid, &entity.CategorySecond.Category.Cname)
		products = append(products, entity)
	}
	return products
}

func (ProductDao) GetProductTotalByCsid(id int) int {
	sql := "select count(*) from product p,categorysecond cs,category c " +
		"where cs.csid = ? and p.csid = cs.csid and cs.cid = c.cid order by pdate desc";
	rows, err := config.ShopDB.Query(sql, id)
	utils.DelError(err)
	var total int
	for rows.Next() {
		rows.Scan(&total)
	}
	return total
}

func (ProductDao) GetProductByPid(id int) entity.Product {
	sql := "select p.pid,p.pname,p.market_price,p.shop_price,p.image,p.pdesc,p.is_hot,pdate,cs.csid,cs.csname,c.cid,c.cname " +
		"from product p,categorysecond cs,category c " +
		"where p.pid = ? and p.csid = cs.csid and cs.cid = c.cid order by pdate desc limit 1";
	rows, err := config.ShopDB.Query(sql, id)
	utils.DelError(err)
	//创建切片
	product := entity.Product{}
	for rows.Next() {
		rows.Scan(&product.Pid, &product.Pname, &product.MarketPrice, &product.ShopPrice, &product.Image, &product.Pdesc,
			&product.IsHot, &product.Pdate, &product.CategorySecond.Csid, &product.CategorySecond.Csname,
			&product.CategorySecond.Category.Cid, &product.CategorySecond.Category.Cname)
	}
	return product
}




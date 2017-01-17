package dao

import (
	"goshop/src/server/entity"
	"goshop/src/server/config"
	"goshop/src/server/utils"
	"database/sql"
)

var (
	productDao = ProductDao{}
)

type OrderDao struct {

}

func (OrderDao) AddOrder(order entity.Order) bool {
	sb := "insert into orders(ordertime,state,total,uid) values(?,?,?,?)"
	rs, err := config.ShopDB.Exec(sb, order.Ordertime, order.State, order.Total, order.User.Uid)
	utils.DelError(err)
	oid, err := rs.LastInsertId()
	rows, err := rs.LastInsertId()
	utils.DelError(err)
	order.Oid = int(oid)
	//插入订单项
	addOrderItemsFlag := addOrderItems(order)
	if rows > 0 && addOrderItemsFlag {
		return true
	}
	return false
}

func addOrderItems(order entity.Order) bool {
	sb := "insert into orderitem(count,total,oid,pid) values(?,?,?,?)"
	orderItems := order.OrderItems
	var rs sql.Result
	var err error
	for _, vale := range orderItems {
		rs, err = config.ShopDB.Exec(sb, vale.Count, vale.Total, order.Oid, vale.Product.Pid)
	}
	utils.DelError(err)
	rows, err := rs.RowsAffected()
	utils.DelError(err)
	if rows > 0 {
		return true
	}
	return false
}

func (OrderDao) GetOrderByUid(uid, start, end int) []entity.Order {
	orders := make([]entity.Order, 0)
	sql := "select oid,addr,name,ordertime,phone,state,total,uid from orders where uid = ? order by oid desc limit ?,?"
	rows, err := config.ShopDB.Query(sql, uid, start, end)
	utils.DelError(err)
	for rows.Next() {
		order := entity.Order{}
		rows.Scan(&order.Oid, &order.Addr, &order.Name, &order.Ordertime, &order.Phone,
			&order.State, &order.Total, &order.User.Uid)
		order.OrderItems = getOrderItemsByOid(order.Oid)
		orders = append(orders, order)
	}
	return orders
}

func (OrderDao) GetOrderByOid(oid int) entity.Order {
	order := entity.Order{}
	sql := "select oid,addr,name,ordertime,phone,state,total,uid from orders where oid = ? order by oid desc limit 1"
	rows, err := config.ShopDB.Query(sql, oid)
	utils.DelError(err)
	for rows.Next() {
		rows.Scan(&order.Oid, &order.Addr, &order.Name, &order.Ordertime, &order.Phone,
			&order.State, &order.Total, &order.User.Uid)
		order.OrderItems = getOrderItemsByOid(order.Oid)
	}
	return order
}

func (OrderDao) GetOrderTotalByUid(uid int) int {
	sql := "select count(*) from orders where uid = ?";
	rows, err := config.ShopDB.Query(sql, uid)
	utils.DelError(err)
	var total int
	for rows.Next() {
		rows.Scan(&total)
	}
	return total
}

func getOrderItemsByOid(oid int) []entity.OrderItem {
	orderItems := make([]entity.OrderItem, 0)
	sql := "select itemid,count,total,oid,pid from orderitem where oid = ? "
	rows, err := config.ShopDB.Query(sql, oid)
	utils.DelError(err)
	for rows.Next() {
		item := entity.OrderItem{}
		rows.Scan(&item.ItemId, &item.Count, &item.Total, &item.Order.Oid, &item.Product.Pid)
		item.Product = productDao.GetProductByPid(item.Product.Pid)
		orderItems = append(orderItems, item)
	}
	return orderItems
}

func (OrderDao) UpdateOrderInfo(order entity.Order) entity.Order {
	sql := "update orders set addr = ? ,name = ?,phone = ?,state = ? where oid = ?"
	rs, err := config.ShopDB.Exec(sql, order.Addr, order.Name, order.Phone, order.State, order.Oid)
	utils.DelError(err)
	rows, err := rs.RowsAffected()
	utils.DelError(err)
	if rows > 0 {
		return order
	}
	return entity.Order{}
}
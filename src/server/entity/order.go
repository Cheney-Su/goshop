package entity

type Order struct {
	Oid        int
	Total      float64
	Ordertime  string
	State      int         // 1:未付款   2:订单已经付款   3:已经发货   4:订单结束
	Name       string
	Phone      string
	Addr       string
	User       User        // 用户的外键:对象
	OrderItems []OrderItem // 配置订单项的集合
}

type OrderItem struct {
	ItemId  int
	Count   int
	Total   float64
	Product Product // 商品外键:对象
	Order   Order   // 订单外键:对象
}

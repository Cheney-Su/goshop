package entity

type Cart struct {
	CartItemMap map[int]CartItem `json:"cartItemMap"` // 购物项集合:Map的key就是商品pid,value:购物项
	Total       float64          `json:"total"`       // 购物总计:
}

type CartItem struct {
	Product Product `json:"product"` // 购物项中商品信息
	Count   int     `json:"count"`   // 购买某种商品数量
	Total   float64 `json:"total"`   // 购买某种商品小计
}

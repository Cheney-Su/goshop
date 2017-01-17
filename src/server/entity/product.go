package entity

type Product struct {
	Pid            int
	Pname          string
	MarketPrice    float64
	ShopPrice      float64
	Image          string
	Pdesc          string
	IsHot          int
	Pdate          string
	CategorySecond CategorySecond // 二级分类的外键:使用二级分类的对象.
}

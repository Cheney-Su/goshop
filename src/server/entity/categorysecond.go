package entity

type CategorySecond struct {
	Csid     int      //主键id
	Csname   string   //二级菜单名称
	Category Category // 所属一级分类.存的是一级分类的对象.
}

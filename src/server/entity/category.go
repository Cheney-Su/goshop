package entity

type Category struct {
	Cid             string          //主键id
	Cname           string          //一级菜单名称
	CategorySeconds []CategorySecond //包含多个二级菜单名称
}

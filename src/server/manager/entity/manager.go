package entity

type Manager struct {
	Uid        int
	Username  string
	Password  string
	SysRoleId int
}

type Menu struct {
	Mid         int
	ParentId   int
	ButtonName string
	SysRoleId  int
	LinkUrl    string
}

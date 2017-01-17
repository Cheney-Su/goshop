package dao

import (
	"goshop/src/server/manager/entity"
	"goshop/src/server/config"
	"goshop/src/server/utils"
)

type ManagerDao struct {
}

func (ManagerDao) ManagerLogin(username, password string) entity.Manager {
	sql := "select uid,username,password,sysRoleId from adminuser where username = ? and password = ?"
	rows, err := config.ShopDB.Query(sql, username, password)
	utils.DelError(err)
	manager := entity.Manager{}
	for rows.Next() {
		rows.Scan(&manager.Uid, &manager.Username, &manager.Password,&manager.SysRoleId)
	}
	return manager
}

func (ManagerDao) GetManagerMenu(sysRoleId int) entity.Menu {
	sql := "select mid,parentId,buttonName,sysRoleId,linkUrl from menu where sysRoleId = ?"
	rows, err := config.ShopDB.Query(sql, sysRoleId)
	utils.DelError(err)
	menu := entity.Menu{}
	for rows.Next() {
		rows.Scan(&menu.Mid, &menu.ParentId, &menu.ButtonName,&menu.SysRoleId,&menu.LinkUrl)
	}
	return menu
}
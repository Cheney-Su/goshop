package service

import (
	"goshop/src/server/manager/dao"
	"goshop/src/server/manager/entity"
)

var (
	managerDao = dao.ManagerDao{}
)

type ManagerService struct {

}

func (ManagerService) ManagerLogin(username,password string) (entity.Manager,entity.Menu) {
	manager := managerDao.ManagerLogin(username,password)
	return manager,managerDao.GetManagerMenu(manager.SysRoleId)
}

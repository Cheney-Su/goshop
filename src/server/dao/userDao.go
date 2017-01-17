package dao

import (
	"goshop/src/server/config"
	"goshop/src/server/utils"
	"goshop/src/server/entity"
)

type UserDao struct {

}

func (UserDao) UserLogin(username, password string) bool {
	sql := "select count(*) from user where username = ? and password = ? and state = 1"
	rows, err := config.ShopDB.Query(sql, username, password)
	utils.DelError(err)
	var rs int
	for rows.Next() {
		rows.Scan(&rs)
	}
	if rs > 0 {
		return true
	}
	return false
}

func (UserDao) GetUserByName(username string) entity.User {
	sql := "select uid,username,password,name,email,phone,addr,state,code from user where username = ? and state = 1"
	rows, err := config.ShopDB.Query(sql, username)
	utils.DelError(err)
	user := entity.User{}
	for rows.Next() {
		rows.Scan(&user.Uid, &user.Username, &user.Password, &user.Name, &user.Email,
			&user.Phone, &user.Addr, &user.State, &user.Code)
	}
	return user
}

func (UserDao) GetUserInfoByUid(userId string) entity.User {
	sql := "select uid,username,password,name,email,phone,addr,state,code from user where uid = ? and state = 1"
	rows, err := config.ShopDB.Query(sql, userId)
	utils.DelError(err)
	user := entity.User{}
	for rows.Next() {
		rows.Scan(&user.Uid, &user.Username, &user.Password, &user.Name, &user.Email,
			&user.Phone, &user.Addr, &user.State, &user.Code)
	}
	return user
}

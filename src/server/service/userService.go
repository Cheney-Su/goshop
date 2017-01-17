package service

import (
	"goshop/src/server/dao"
	"github.com/kataras/iris"
	"goshop/src/server/redis"
	"strconv"
	"goshop/src/server/entity"
)

var (
	userDao = &dao.UserDao{}
)

type UserService struct {

}

func (UserService) UserLogin(username, password string, ctx *iris.Context) (bool, string) {
	var toUrl string
	userLoginResult := userDao.UserLogin(username, password)
	if userLoginResult {
		//获取登录前一个页面的url，登录成功直接跳转
		toUrl = ctx.Session().Get("toUrl").(string)
		//将用户id保存到session中
		user := userDao.GetUserByName(username)
		ctx.Session().Set("userId", strconv.Itoa(user.Uid))
		//将购物车中的商品缓存到对应的用户id上，并删除之前sessionId的key
		sessionId := ctx.Session().ID()
		value := redis.Get(sessionId)
		redis.Delete(sessionId)
		redis.Set(strconv.Itoa(user.Uid), value)
	}
	return userLoginResult, toUrl
}

func (UserService) GetUserInfoByUid(userId string) entity.User {
	return userDao.GetUserInfoByUid(userId)
}
package service

import (
	"goshop/src/server/redis"
	"log"
	"goshop/src/server/entity"
	"strconv"
	"encoding/json"
)

type CartService struct {

}

func (CartService) ListCart(userId, sessionId string) entity.Cart {
	log.Println("start...")
	var cart entity.Cart
	if len(userId) != 0 {
		//用户登录还未处理
		log.Println("用户已登录...")
		value := redis.Get(userId)
		json.Unmarshal([]byte(value), &cart)
	} else {
		log.Println("用户未登录...")
		value := redis.Get(sessionId)
		json.Unmarshal([]byte(value), &cart)
	}
	return cart
}

func (CartService) AddCart(userId, sessionId, pidString, countString string) entity.Cart {
	log.Println("start...")
	var userCart entity.Cart
	if len(userId) != 0 {
		//用户登录还未处理
		log.Println("用户已登录...")
		userCart = addCart(userId, pidString, countString)
	} else {
		log.Println("用户未登录...")
		userCart = addCart(sessionId, pidString, countString)
		//缓存一天
		redis.Expire(sessionId, 24 * 60 * 60)
	}
	return userCart
}

func addCart(id, pidString, countString string) entity.Cart {
	var userCart entity.Cart
	pid, _ := strconv.Atoi(pidString)
	countInt, _ := strconv.Atoi(countString)
	count, _ := strconv.ParseFloat(countString, 64)
	value := redis.Get(id)
	product := productService.GetProductByPid(pid)
	if len(value) == 0 {
		cartItem := entity.CartItem{Product:product, Count:countInt, Total:(count * product.ShopPrice)}
		cartItemMap := make(map[int]entity.CartItem)
		cartItemMap[product.Pid] = cartItem
		cart := entity.Cart{CartItemMap:cartItemMap, Total:cartItem.Total}
		userCart = cart
		byteJson, _ := json.Marshal(cart)
		redis.Set(id, string(byteJson))
		log.Println("redis缓存中没有该session的购物车信息...")
	} else {
		var cart entity.Cart
		json.Unmarshal([]byte(value), &cart)
		cartItemMap := cart.CartItemMap
		cartTotal := cart.Total
		cartItemMapLen := len(cartItemMap)
		if cartItemMapLen == 0 {
			cartItem := entity.CartItem{Product:product, Count:countInt, Total:(count * product.ShopPrice)}
			cartItemMap[product.Pid] = cartItem
			cartTotal += count * product.ShopPrice
		}
		for i := 0; i < cartItemMapLen; i++ {
			cartItem := cartItemMap[product.Pid]
			if product.Pid == cartItem.Product.Pid {
				log.Println("redis购物车中已存在该商品的信息")
				cartItemCount := cartItem.Count + countInt
				cartItem.Count = cartItemCount
				cartItem.Total = float64(cartItemCount) * product.ShopPrice
				cartItemMap[product.Pid] = cartItem
				cartTotal += count * product.ShopPrice
				break
			}
			if i == (len(cartItemMap) - 1) {
				cartItem = entity.CartItem{Product:product, Count:countInt, Total:(count * product.ShopPrice)}
				cartItemMap[product.Pid] = cartItem
				cartTotal += count * product.ShopPrice
			}

		}
		cart.Total = cartTotal
		userCart = cart
		byteJson, _ := json.Marshal(cart)
		redis.Set(id, string(byteJson))
		log.Println("redis缓存中有该session的购物车信息...")
	}
	return userCart
}

func (CartService) DeleteCart(userId, sessionId, pidString string) bool {
	log.Println("start...")
	if len(userId) != 0 {
		//用户登录还未处理
		log.Println("用户已登录...")
		deleteCart(userId, pidString)
	} else {
		log.Println("用户未登录...")
		deleteCart(sessionId, pidString)
	}
	return true
}

func deleteCart(id, pidString string) bool {
	pid, _ := strconv.Atoi(pidString)
	value := redis.Get(id)
	product := productService.GetProductByPid(pid)
	var cart entity.Cart
	json.Unmarshal([]byte(value), &cart)
	cartItemMap := cart.CartItemMap
	cartTotal := cart.Total
	cartItemMapLen := len(cartItemMap)
	for i := 0; i < cartItemMapLen; i++ {
		cartItem := cartItemMap[product.Pid]
		if product.Pid == cartItem.Product.Pid {
			log.Println("redis购物车中已存在该商品的信息")
			cartTotal -= cartItem.Total
			delete(cartItemMap, product.Pid)
		}
	}
	cart.Total = cartTotal
	byteJson, _ := json.Marshal(cart)
	redis.Set(id, string(byteJson))
	log.Println("redis缓存中有该session的购物车信息...")
	return true
}

func (CartService) ClearCart(userId, sessionId string) bool {
	log.Println("start...")
	var flag bool
	if len(userId) != 0 {
		//用户登录还未处理
		log.Println("用户已登录...")
		flag = redis.Delete(userId)
	} else {
		log.Println("用户未登录...")
		flag = redis.Delete(sessionId)
	}
	return flag
}
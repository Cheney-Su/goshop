package service

import (
	"goshop/src/server/dao"
	"github.com/kataras/iris"
	"goshop/src/server/redis"
	"goshop/src/server/entity"
	"encoding/json"
	"time"
	"strconv"
	"goshop/src/server/utils"
)

var (
	orderDao = dao.OrderDao{}
)

type OrderService struct {

}

func (OrderService) AddOrder(ctx *iris.Context) bool {
	userId := ctx.Session().GetString("userId")
	value := redis.Get(userId)
	var cart entity.Cart
	json.Unmarshal([]byte(value), &cart)
	var order entity.Order
	order.Total = cart.Total
	order.State = 1        //未付款
	order.Ordertime = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	user := userDao.GetUserInfoByUid(userId)
	order.User = user
	cartItemMap := cart.CartItemMap
	orderItems := make([]entity.OrderItem, 0)
	for _, cartItem := range cartItemMap {
		orderItem := entity.OrderItem{}
		orderItem.Count = cartItem.Count
		orderItem.Total = cartItem.Total
		orderItem.Product = cartItem.Product
		orderItem.Order = order

		orderItems = append(orderItems, orderItem)
	}
	order.OrderItems = orderItems
	//json, _ := json.Marshal(order)
	//fmt.Println(string(json))
	orderDao.AddOrder(order)
	return true
}

func (OrderService) GetOrderByUid(uid, page, pageSize int) ([]entity.Order, int) {
	start := (page - 1) * pageSize
	end := pageSize
	return orderDao.GetOrderByUid(uid, start, end), orderDao.GetOrderTotalByUid(uid)
}
func (OrderService) GetOrderByOid(oid int) entity.Order {
	return orderDao.GetOrderByOid(oid)
}

func (OrderService) UpdateOrderInfo(oid, addr, name, phone string) entity.Order {
	oidInt, _ := strconv.Atoi(oid)
	order := entity.Order{Oid:oidInt, Addr:addr, Name:name, Phone:phone, State:1}
	return orderDao.UpdateOrderInfo(order)
}

func (OrderService) PayOrder(oid, bank string) string {
	// 付款需要的参数:
	p0_Cmd := utils.Cmd        // 业务类型:
	p1_MerId := utils.MerId        // 商户编号:
	p2_Order := "60017"        // 订单编号:
	p3_Amt := "0.01"        // 付款金额:
	p4_Cur := "CNY"        // 交易币种:
	p5_Pid := "test"        // 商品名称:
	p6_Pcat := "testType"        // 商品种类:
	p7_Pdesc := "thisistestpay"        // 商品描述:
	p8_Url := "http://localhost:8088/order/payBack"        // 商户接收支付成功数据的地址:
	p9_SAF := ""        // 送货地址:
	pa_MP := ""        // 商户扩展信息:
	pd_FrpId := bank        // 支付通道编码:
	pr_NeedResponse := "1"        // 应答机制:

	hmac := utils.Hmac(p0_Cmd, p1_MerId, p2_Order, p3_Amt, p4_Cur, p5_Pid, p6_Pcat, p7_Pdesc, p8_Url, p9_SAF, pa_MP, pd_FrpId, pr_NeedResponse)
	url := "https://www.yeepay.com/app-merchant-proxy/node?"
	params := "p0_Cmd=" + p0_Cmd + "&"
	params += "p1_MerId=" + p1_MerId + "&"
	params += "p2_Order=" + p2_Order + "&"
	params += "p3_Amt=" + p3_Amt + "&"
	params += "p4_Cur=" + p4_Cur + "&"
	params += "p5_Pid=" + p5_Pid + "&"
	params += "p6_Pcat=" + p6_Pcat + "&"
	params += "p7_Pdesc=" + p7_Pdesc + "&"
	params += "p8_Url=" + p8_Url + "&"
	params += "p9_SAF=" + p9_SAF + "&"
	params += "pa_MP=" + pa_MP + "&"
	params += "pd_FrpId=" + pd_FrpId + "&"
	params += "pr_NeedResponse=" + pr_NeedResponse + "&"
	params += "hmac=" + hmac

	return url + params
}

func (OrderService) CallBackPayOrder(oid string) {
	oidInt, _ := strconv.Atoi(oid)
	order := orderDao.GetOrderByOid(oidInt)
	order.State = 2
	orderDao.UpdateOrderInfo(order)
}
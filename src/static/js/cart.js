/**
 * Created by Administrator on 2016/12/29.
 */
//购物车总额全局变量
var cartTotal = 0
function delCart(pid) {
    $.ajax({
        url: '/cart/delete/' + pid,
        type: 'get',
        dataType: 'json',
        data: "",
        error: function () {
            alert('error');
        },
        success: function (data) {
            //请求成功
            if (data.status == 0) {
                var result = data.data
                var len = 0
                $.each($(".cart" + pid).parent().parent(), function (index, val) {
                    len = val.childNodes.length
                })
                if (len == 1) {
                    var stepHtml = "<span><h2>亲!您还没有购物!请先去购物!</h2></span>"
                    $("#step24").html(stepHtml);
                } else {
                    $.each($(".cart" + pid).prev(), function (index, val) {
                        cartTotal -= val.innerText
                    })
                    $(".cartTotal").html(cartTotal);
                    $(".cart" + pid).parent().remove()
                }

            }
        }
    });
}

function clearCart() {
    $.ajax({
        url: '/cart/clear',
        type: 'get',
        dataType: 'json',
        data: "",
        error: function () {
            alert('error');
        },
        success: function (data) {
            //请求成功
            if (data.status == 0) {
                var stepHtml = "<span><h2>亲!您还没有购物!请先去购物!</h2></span>"
                $("#step24").html(stepHtml);
            }
        }
    });
}

function addOrder() {
    $.ajax({
        url: '/order/add',
        type: 'get',
        dataType: 'json',
        data: "",
        error: function () {
            alert('error');
        },
        success: function (data) {
            //请求成功
            if (data.status == 0) {
                var result = data.Data
                window.location.href = "/views/orderList.html?uid=" + uid
            } else {
                alert("系统异常，请稍后重试...")
                return
            }
        }
    });
}

$(function () {
    if (getQueryString("pid") != null) {
        $.ajax({
            url: "/cart/add/" + getQueryString("pid") + "?count=" + getQueryString("count"),
            type: 'post',
            dataType: 'json',
            data: "",
            error: function () {
                alert('error');
            },
            success: function (data) {
                var cartTotalHtml = 0
                //请求成功
                if (data.status == 0) {
                    var result = data.data
                    var cartItemMap = eval(result.cartItemMap)
                    var cartListHtml = ""
                    $.each(cartItemMap, function (index, val) {
                        cartListHtml += "<tr><td width='60px'><img src='/image/" + val.product.Image + "'/></td>" +
                            "<td><a href='/views/product.html?cid=" + val.product.CategorySecond.Category.Cid +
                            "&csid=" + val.product.CategorySecond.Csid + "&pid=" + val.product.Pid + "'>" + val.product.Pname + "</a></td>" +
                            "<td>￥" + val.product.ShopPrice + "</td>" +
                            "<td class='quantity' width='60px'>" + val.count + "</td>" +
                            "<td width='140px'><span class='subtotal'>" + val.total + "</span></td>" +
                            "<td class='cart" + val.product.Pid + "'><a href='javascript:void(0)' onclick='delCart(" + val.product.Pid + ")' class='delete'>删除</a></td></tr>"
                    });
                    cartTotal = result.total
                }
                $("#cartItems").html(cartListHtml);
                $(".cartTotal").html(cartTotal);
            }
        });
    } else {
        $.ajax({
            url: '/cart/list/',
            type: 'get',
            dataType: 'json',
            data: "",
            error: function () {
                alert('error');
            },
            success: function (data) {
                var cartTotalHtml = 0
                //请求成功
                if (data.status == 0) {
                    var result = data.data
                    var cartItemMap = eval(result.cartItemMap)
                    var cartListHtml = ""
                    if (cartItemMap != null) {
                        $.each(cartItemMap, function (index, val) {
                            cartListHtml += "<tr><td width='60px'><img src='/image/" + val.product.Image + "'/></td>" +
                                "<td><a href='/views/product.html?cid=" + val.product.CategorySecond.Category.Cid +
                                "&csid=" + val.product.CategorySecond.Csid + "&pid=" + val.product.Pid + "'>" + val.product.Pname + "</a></td>" +
                                "<td>￥" + val.product.ShopPrice + "</td>" +
                                "<td class='quantity' width='60px'>" + val.count + "</td>" +
                                "<td width='140px'><span class='subtotal'>" + val.total + "</span></td>" +
                                "<td class='cart" + val.product.Pid + "'><a href='javascript:void(0)' onclick='delCart(" + val.product.Pid + ")' class='delete'>删除</a></td></tr>"
                        });
                        cartTotal = result.total
                    }
                }
                if (cartListHtml == "") {
                    var stepHtml = "<span><h2>亲!您还没有购物!请先去购物!</h2></span>"
                    $("#step24").html(stepHtml);
                } else {
                    $("#cartItems").html(cartListHtml);
                    $(".cartTotal").html(cartTotal);
                }
            }
        });
    }

})
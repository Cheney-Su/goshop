/**
 * Created by Administrator on 2016/12/21.
 */
var result = ""
var uid = ""
$(function () {
    $.ajax({
        url: '/category/',
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
                var categoryListHtml = ""
                for (i = 0; i < result.length; i++) {
                    //异步加载
                    categoryListHtml += "<li><a href='/views/productList.html?cid=" + result[i].Cid + "'>" + result[i].Cname + "</a> |</li>"
                }
            }
            $("#categoryList").html(categoryListHtml);
        }
    });

    $.ajax({
        url: '/user/userInfo',
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
                var userLoginHtml = ""
                uid = result.Uid
                //异步加载
                var userLoginHtml = "<li id='headerLogin' class='headerLogin' style='display: list-item;'>" + result.Name + "|</li>" +
                    "<li id='headerLogin' class='headerLogin' style='display: list-item;'>" +
                    "<a href='/views/orderList.html?uid=" + result.Uid + "'>我的订单</a>|</li>" +
                    "<li id='headerRegister' class='headerRegister' style='display: list-item;'>" +
                    "<a href='${ pageContext.request.contextPath }/user_quit.action'>退出</a>|</li>" +
                    "<li><a>会员中心</a> |</li> <li><a>购物指南</a> |</li> <li><a>关于我们</a></li>";
            }
            $("#existUser").html(userLoginHtml);
        }
    });

    $("#loginForm").submit(function () {
        $.ajax({
            url: '/user/login',
            type: 'post',
            dataType: 'json',
            data: $("#loginForm").serialize(),
            error: function () {
                alert('error');
            },
            success: function (data) {
                if (data.status == 0) {
                    result = data.Data
                    window.location.href = result
                } else {
                    alert("亲，用户名不存在或者密码错误...")
                }
            }
        })
        return false
    });
})
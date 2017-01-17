/**
 * Created by Administrator on 2017/1/4.
 */
var page = 1, pageSize = 20

var orderList = function (page, pageSize) {
    $.ajax({
        url: '/order/uid/' + getQueryString("uid") + "?page=" + page + "&pageSize=" + pageSize,
        type: 'get',
        dataType: 'json',
        data: "",
        error: function () {
            alert('error');
        },
        success: function (data) {
            var orderListHtml = ""
            //请求成功
            if (data.status == 0) {
                var result = data.data
                for (var i = 0; i < result.length; i++) {
                    orderListHtml += "<tr><th colspan='5'>订单编号:" + result[i].Oid +
                        "&nbsp;&nbsp;&nbsp;&nbsp;订单金额:<font style='color: black'>" + result[i].Total +
                        "元</font>&nbsp;&nbsp;&nbsp;&nbsp;<font style='color: red'>";
                    if (result[i].State == 1) {
                        orderListHtml += "<button onclick='window.location.href = " + "\"/views/order.html?oid=" + result[i].Oid + "\"'>付款</button>";
                    } else if (result[i].State == 2) {
                        orderListHtml += "已付款，等待卖家发货！！！";
                    } else if (result[i].State == 3) {
                        orderListHtml += "<button>确认收货</button>";
                    } else if (result[i].State == 4) {
                        orderListHtml += "交易成功";
                    }
                    orderListHtml += "</font></th></tr>" +
                        "<tr><th>图片</th><th>商品</th><th>价格</th><th>数量</th><th>小计</th></tr>";
                    var orderItem = result[i].OrderItems
                    for (var j = 0; j < orderItem.length; j++) {
                        orderListHtml += "<tr><td width='60px'><img src='/image/" + orderItem[j].Product.Image + "'/></td>" +
                            "<td><a href='/views/product.html?cid=" + orderItem[j].Product.CategorySecond.Category.Cid +
                            "&csid=" + orderItem[j].Product.CategorySecond.Csid + "&pid=" + orderItem[j].Product.Pid + "'>" + orderItem[j].Product.Pname + "</a></td>" +
                            "<td>￥" + orderItem[j].Product.ShopPrice + "</td>" +
                            "<td class='quantity' width='60px'>" + orderItem[j].Count + "</td>" +
                            "<td width='140px'><span class='subtotal'>" + orderItem[j].Total + "</span></td></tr>"
                    }
                }
                var paginationHtml = "";
                var totalPage = Math.ceil(data.total / pageSize);
                if (totalPage >= 1) {
                    paginationHtml += "<span>第" + page + "/" + totalPage + "页</span>"
                    if (page != 1) {
                        paginationHtml += "<a href='javascript:void(0)' class='firstPage' onclick='orderList(" + 1 + "," + pageSize + ")'>&nbsp;</a>";
                        paginationHtml += "<a href='javascript:void(0)' class='previousPage' onclick='orderList(" + (page - 1) + ", " + pageSize + ")'>&nbsp;</a>";
                    }
                    for (i = 1; i <= totalPage; i++) {
                        if (i == page) {
                            paginationHtml += "<span class='currentPage'>" + i + "</span>";
                        } else {
                            paginationHtml += "<a href='javascript:void(0)' onclick='orderList(" + i + ", " + pageSize + ")'>" + i + "</a>";
                        }
                    }
                    if (page != totalPage) {
                        paginationHtml += "<a href='javascript:void(0)' class='nextPage' onclick='orderList(" + (page + 1) + ", " + pageSize + ")'>&nbsp;</a>";
                        paginationHtml += "<a href='javascript:void(0)' class='lastPage' onclick='orderList(" + totalPage + ", " + pageSize + ")'>&nbsp;</a>";
                    }
                }

            } else if (data.status == 6003) {
                window.location.href = "/views/login.html"
            } else {
                alert("系统异常，请稍后重试...")
                return
            }
            $("#orderList").html(orderListHtml)
            $("#pagination").html(paginationHtml);
        }
    });
}

$(function () {
    orderList(page, pageSize)
})

/**
 * Created by Administrator on 2017/1/5.
 */
function payOrder() {
    // $("#orderForm").attr("action", "/order/pay/" + getQueryString("oid"));
    // $("#orderForm").submit()
    var data = {
        addr: $("#addr").val(),
        name: $("#name").val(),
        phone: $("#phone").val(),
        pd_FrpId: $("input[name='pd_FrpId']:checked").val()
    }
    $.ajax({
        url: '/order/pay/' + getQueryString("oid"),
        type: 'post',
        dataType: 'json',
        data: data,
        error: function () {
            alert('error');
        },
        success: function (data) {
            if (data.status == 0) {
                window.location.href = data.data;
            }
        }
    });
}

$(function () {
    $.ajax({
        url: '/order/oid/' + getQueryString("oid"),
        type: 'get',
        dataType: 'json',
        data: "",
        error: function () {
            alert('error');
        },
        success: function (data) {
            var orderHtml = ""
            //请求成功
            if (data.status == 0) {
                var result = data.data
                orderHtml += "<tr><th colspan='5'>订单编号:" + result.Oid + "&nbsp;&nbsp;&nbsp;&nbsp;</th></tr>" +
                    "<tr><th>图片</th><th>商品</th><th>价格</th><th>数量</th><th>小计</th></tr>";
                var orderItem = result.OrderItems
                for (var i = 0; i < orderItem.length; i++) {
                    orderHtml += "<tr><td width='60px'><img src='/image/" + orderItem[i].Product.Image + "'/></td>" +
                        "<td><a href='/views/product.html?cid=" + orderItem[i].Product.CategorySecond.Category.Cid +
                        "&csid=" + orderItem[i].Product.CategorySecond.Csid + "&pid=" + orderItem[i].Product.Pid + "'>" + orderItem[i].Product.Pname + "</a></td>" +
                        "<td>￥" + orderItem[i].Product.ShopPrice + "</td>" +
                        "<td class='quantity' width='60px'>" + orderItem[i].Count + "</td>" +
                        "<td width='140px'><span class='subtotal'>" + orderItem[i].Total + "</span></td></tr>"
                }
                $("#addr").val(result.Addr)
                $("#name").val(result.Name)
                $("#phone").val(result.Phone)
            }
            $("#order").html(orderHtml)
            $("#effectivePrice").html("￥" + result.Total + "元")
        }
    })
})
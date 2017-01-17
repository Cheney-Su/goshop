/**
 * Created by Administrator on 2016/12/21.
 */
function getQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]);
    return null;
}

$(function () {
    $.ajax({
        url: '/views/menu.html',
        type: 'get',
        dataType: 'html',
        data: "",
        error: function () {
            alert('error');
        },
        success: function (data) {
            $("#myDiv").html(data);
        }
    });

    $.ajax({
        url: '/product/hot/',
        type: 'get',
        dataType: 'json',
        data: '',
        error: function () {
            alert('error')
        },
        success: function (data) {
            if (data.status == 0) {
                var result = data.data
                var productHotListHtml = ""
                for (i = 0; i < result.length; i++) {
                    productHotListHtml += "<li><a href='/views/product.html?cid=" + result[i].CategorySecond.Category.Cid +
                        "&csid=" + result[i].CategorySecond.Csid + "&pid=" + result[i].Pid + "'><img src='/image/" + result[i].Image + "' style='display: block'/></a>"
                }
                $("#hotProductAd").html(productHotListHtml)
            }
        }
    })

    $.ajax({
        url: '/product/new/',
        type: 'get',
        dataType: 'json',
        data: '',
        error: function () {
            alert('error')
        },
        success: function (data) {
            if (data.status == 0) {
                var result = data.data
                var productNewListHtml = ""
                for (i = 0; i < result.length; i++) {
                    productNewListHtml += "<li><a href=''><img src='/image/" + result[i].Image + "' style='display: block'/></a>"
                }
                $("#newProductAd").html(productNewListHtml)
            }
        }
    })

    $.ajax({
        url: '/index/filter?toUrl=' + document.referrer,
        type: 'get',
        dataType: 'json',
        data: '',
        error: function () {
            alert('error')
        },
        success: function (data) {
        }
    })
})
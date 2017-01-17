/**
 * Created by Administrator on 2016/12/26.
 */
var page = 1, pageSize = 20
//分页获取商品列表，根据一级分类，二级分类
var product = function (cid, csid, page, pageSize) {
    if (cid == 0) {
        cid = getQueryString("cid")
    }
    var url
    if (csid == 0) {
        url = "/product/cid/" + cid + "?page=" + page + "&pageSize=" + pageSize
    } else {
        url = "/product/csid/" + csid + "?page=" + page + "&pageSize=" + pageSize
    }
    $.ajax({
        url: url,
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
                var productListHtml = ""
                for (i = 0; i < result.length; i++) {
                    //异步加载
                    productListHtml += "<li><a href='/views/product.html?cid=" + result[i].CategorySecond.Category.Cid +
                        "&csid=" + result[i].CategorySecond.Csid + "&pid=" + result[i].Pid + "'>" +
                        "<img src='/image/" + result[i].Image + "' width='170' height='170' style='display: inline-block'/>" +
                        "<span style='color:green'>" + result[i].Pname + "</span>" +
                        "<span class='price'>商城价：" + result[i].ShopPrice + "</span>" +
                        "</a>"
                }
                var paginationHtml = "";
                var totalPage = Math.ceil(data.total / pageSize);
                if (totalPage >= 1) {
                    paginationHtml += "<span>第" + page + "/" + totalPage + "页</span>"
                    if (page != 1) {
                        paginationHtml += "<a href='javascript:void(0)' class='firstPage' onclick='product(" + cid + "," + csid + "," + 1 + ", " + pageSize + ")'>&nbsp;</a>";
                        paginationHtml += "<a href='javascript:void(0)' class='previousPage' onclick='product(" + cid + "," + csid + "," + (page - 1) + ", " + pageSize + ")'>&nbsp;</a>";
                    }
                    for (i = 1; i <= totalPage; i++) {
                        if (i == page) {
                            paginationHtml += "<span class='currentPage'>" + i + "</span>";
                        } else {
                            paginationHtml += "<a href='javascript:void(0)' onclick='product(" + cid + "," + csid + "," + i + ", " + pageSize + ")'>" + i + "</a>";
                        }
                    }
                    if (page != totalPage) {
                        paginationHtml += "<a href='javascript:void(0)' class='nextPage' onclick='product(" + cid + "," + csid + "," + (page + 1) + ", " + pageSize + ")'>&nbsp;</a>";
                        paginationHtml += "<a href='javascript:void(0)' class='lastPage' onclick='product(" + cid + "," + csid + "," + totalPage + ", " + pageSize + ")'>&nbsp;</a>";
                    }
                }
            }
            $("#result").html(productListHtml);
            $("#pagination").html(paginationHtml);
        }
    });
}

//商品加入购物车操作，在url上拼接上商品id，数量
function saveCart() {
    cartForm = $("#cartForm").attr("action")
    $("#cartForm").attr("action", cartForm + "?pid=" + $("#formPid").val() + "&count=" + $("#count").val())
    $("#cartForm").submit();
}

$(function () {
    //默认获取用户点击导航栏一级分类的商品
    product(0, 0, page, pageSize)

    //获取侧边栏的一级二级分类菜单栏
    $.ajax({
        url: '/category/categorysecond/',
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
                var hotProductCategoryListHtml = ""
                for (i = 0; i < result.length; i++) {
                    var secondsResult = result[i].CategorySeconds
                    // 异步加载
                    hotProductCategoryListHtml += "<dl><dt>" +
                        "<a href='javascript:void(0)' onclick='product(" + result[i].Cid + ",0,page,pageSize)'>" + result[i].Cname + "</a></dt>";
                    for (j = 0; j < secondsResult.length; j++) {
                        hotProductCategoryListHtml += "<dd>" +
                            "<a href='javascript:void(0)' onclick='product(0," + secondsResult[j].Csid + ",page,pageSize)'>" + secondsResult[j].Csname + "</a></dd>";
                    }
                    hotProductCategoryListHtml += "</dl>";
                }
            }
            $("#hotProductCategory").html(hotProductCategoryListHtml);
        }
    });

    $.ajax({
        url: '/product/pid/' + getQueryString("pid"),
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
                $("#pname").html(result.Pname)
                $("#pid").html(result.Pid)
                $("#shopPrice").html(result.ShopPrice)
                $("#marketPrice").html(result.MarketPrice)
                $("#pdesc").html(result.Pdesc)
                $(".pimage").attr('src', "/image/" + result.Image)
                $("#formPid").val(result.Pid)
                // var hotProductCategoryListHtml = ""
                // for (i = 0; i < result.length; i++) {
                //     var secondsResult = result[i].CategorySeconds
                //     // 异步加载
                //     hotProductCategoryListHtml += "<dl><dt>" +
                //         "<a href='javascript:void(0)' onclick='product(" + result[i].Cid + ",0,page,pageSize)'>" + result[i].Cname + "</a></dt>";
                //     for (j = 0; j < secondsResult.length; j++) {
                //         hotProductCategoryListHtml += "<dd>" +
                //             "<a href='javascript:void(0)' onclick='product(0," + secondsResult[j].Csid + ",page,pageSize)'>" + secondsResult[j].Csname + "</a></dd>";
                //     }
                //     hotProductCategoryListHtml += "</dl>";
                // }
            }
            // $("#hotProductCategory").html(hotProductCategoryListHtml);
        }
    });

})
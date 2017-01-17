/**
 * Created by Administrator on 2017/1/3.
 */
var result = ""
$(function () {
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

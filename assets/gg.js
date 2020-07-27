// jquery ajax
function gg_request(type ,url, data, successCallback) {
    layui.use(['layer', 'jquery'], function () {
        let layer = layui.layer;
        let $ = layui.jquery;
        let loading = layer.load(0, {shade: false});
        $.ajax({
            type: type,
            url: url,
            dataType: 'json',
            contentType: "application/json; charset=utf-8",
            data: data,
            success: function (res) {
                successCallback(res);
                layer.close(loading);
            },
            error: function (e) {
                layer.close(loading);
                console.log(e);
            }
        })
    });
}



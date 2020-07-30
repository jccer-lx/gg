var cache = []
var layerIndex;

// jquery ajax
function gg_request(type, url, data, successCallback) {
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

function gg_table(gg_table_option) {
    layui.use(['table', 'layer', 'jquery'], function () {
        let table = layui.table;
        var layer = layui.layer;
        let $ = layui.jquery;
        let apiPath = '/' + gg_table_option.table_name + '/api';
        let viewPath = '/' + gg_table_option.table_name + '/view';
        // console.log(apiPath);
        // console.log('#tool-' + gg_table_option.table_name + '-table');
        table.render({
            toolbar: '#tool-' + gg_table_option.table_name + '-table',
            elem: '#' + gg_table_option.table_name + '-table',
            url: apiPath + '/list', //数据接口
            page: true, //开启分页
            cols: gg_table_option.cols
        });

        function tableReload() {
            table.reload(gg_table_option.table_name + '-table', {
                url: apiPath + '/list'
            })
        }

        //监听事件
        table.on('toolbar(' + gg_table_option.table_name + '-table)', function (obj) {
            switch (obj.event) {
                case 'add':
                    $.get(viewPath + '/add', function (content) {
                        layerIndex = layer.open({
                            type: 1,
                            anim: 5,
                            title: "新增",
                            shadeClose: true, //开启遮罩关闭
                            content: content,
                            maxmin: true,
                            area: ['60vw', '60vh'],
                        });
                        // layer.full(addIndex);
                    })
                    break;
                case 'del':
                    layer.confirm('确定删除？', {
                        btn: ['yes', 'no'] //按钮
                    }, function () {
                        let checkStatus = table.checkStatus(gg_table_option.table_name + '-table');
                        let ids = [];
                        $.each(checkStatus.data, function (i, item) {
                            ids.push(item.id)
                        })
                        gg_request("delete", apiPath + '/delete', JSON.stringify({ids: ids}), function (res) {
                            if (res.code === 0) {
                                tableReload();
                                layer.msg('成功');
                            } else {
                                layer.msg(res.msg);
                            }
                        });

                    }, function () {

                    });
            }
        });

        table.on('edit(' + gg_table_option.table_name + '-table)', function (obj) { //注：edit是固定事件名，test是table原始容器的属性 lay-filter="对应的值"
            gg_request('put', apiPath + '/edit/' + obj.data.id, JSON.stringify(obj.data), function (res) {
                if (res.code !== 0) {
                    tableReload()
                    layer.msg(res.msg)
                }
            });

        });

        table.on('tool(' + gg_table_option.table_name + '-table)', function (obj) {
            let layEvent = obj.event; //获得 lay-event 对应的值（也可以是表头的 event 参数对应的值）
            $.each(gg_table_option.cols[0], function (i, col) {
                if (col.event === layEvent) {
                    col.event_func(obj);
                }
            });
        });
    });
}


function gg_upload_pic(elem, multiple, done) {
    layui.use(['layer', 'jquery', 'upload'], function () {
        let upload = layui.upload;
        let $ = layui.jquery;
        let layer = layui.layer;
        upload.render({
            elem: elem, //绑定元素
            url: '/upload/api/pic', //上传接口
            accept: 'images',
            acceptMime: 'image/*',
            multiple: multiple,
            field: 'pic_file',
            size: 10240,
            done: function (res) {
                if (res.code === 0) {
                    done(res)
                } else {
                    layer.msg(res.msg)
                }
            },
            error: function () {
                layer.msg("上传失败")
            }
        });
    });
}

function gg_add_form(form_name, beforeSubmit) {
    layui.use(['layer', 'jquery', 'form', 'table', 'upload'], function () {
        let layer = layui.layer;
        let table = layui.table;
        let $ = layui.jquery;
        let form = layui.form;
        let apiPath = '/' + form_name + '/api';
        let viewPath = '/' + form_name + '/view';
        $('#' + form_name + '-add >button.layui-btn').on("click", function () {
            form.render();
            if(beforeSubmit !== undefined){
                beforeSubmit();
            }
            gg_request('post', apiPath + '/add',
                JSON.stringify(form.val(form_name + '-add')),
                function (res) {
                    if (res.code === 0) {
                        tableReload()
                        layer.close(layerIndex);
                    } else {
                        layer.msg(res.msg)
                    }
                });
        });

        function tableReload() {
            table.reload(form_name + '-table', {
                url: apiPath + '/list'
            })
        }
    });
}

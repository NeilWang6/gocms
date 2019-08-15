var sexMap = ["未知","男","女"];
var gradeMap = {1:"幼小班",2:"幼中班",3:"幼大班",10:"一年级",20:"二年级",30:"三年级",40:"四年级",50:"五年级",60:"六年级",70:"初一",80:"初二",90:"初三",100:"高一",110:"高二",120:"高三"};
var studentStatusMap = {"-1":"删除","0":"未报名","1":"正常","2":"休眠"};
var teacherStatusMap = {"-1":"删除","0":"正常","1":"休眠"};
var relateMap = ["爸爸","妈妈","爷爷","奶奶","外公","外婆","其他"];
var lengthMap = ["0.5","1","1.5","2"];
var contractStatusMap = {"0":"正常","1":"禁用","2":"退费"};
var paymentMap = {"cash":"现金","pos":"pos机"};
var contractTypeMap = {"0":"一对一","1":"小班","2":"托班","-1":"合计"};
var contractTypeMap1 = {"0":"一对一","1":"小班","2":"托班"};
var classTimeRangeMap = ["6-8","8-10","10-12","13-15","15-17","17-19", "19-21","21-23"];
var balanceRate = ["0-10","10-20","20-50","50-100","100-100000"];
var cityMap = ["上海","山东"];
function sexName(sexId) {
    return sexMap[sexId];
}
function gradeName(gradeId) {
    return gradeMap[gradeId];
}
function studentStatusAlias(statusId) {
    return studentStatusMap[statusId];
}
function teacherStatusAlias(statusId) {
    return teacherStatusMap[statusId];
}
function contractStatusAlias(statusId) {
    return contractStatusMap[statusId];
}

function totalDetail(data, id) {
    var len = data.length;
    var html = [];
    for (var i=0;i < len; i++) {
        var tet = data[i]["color"];
        html.push('<li><i class="fa fa-circle-o " style="color: '+data[i]['color']+'"></i> '+data[i]['label']+': '+data[i]['data']+'</li>');
    }
    var $select = $("#"+id);
    $select.html(html.join(''));
}
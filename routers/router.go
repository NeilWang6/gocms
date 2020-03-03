package routers

import (
	"github.com/astaxie/beego"
	"github.com/cuua/gocms/controllers"
	"github.com/cuua/gocms/controllers/class"
	"github.com/cuua/gocms/controllers/fronted"
	"github.com/cuua/gocms/controllers/fronted_manage"
	"github.com/cuua/gocms/controllers/student"
	"github.com/cuua/gocms/controllers/teacher"
)

func init() {
	/*student start*/
	beego.Router("/student/student/index", &student.StudentController{}, "*:Index")
	beego.Router("/student/student/datagrid", &student.StudentController{}, "Get,Post:DataGrid")
	beego.Router("/student/student/edit/?:id", &student.StudentController{}, "Get,Post:Edit")
	beego.Router("/student/student/delete", &student.StudentController{}, "Get,Post:Delete")
	beego.Router("/student/student/datalist", &student.StudentController{}, "Get,Post:DataList")
	beego.Router("/student/student/search", &student.StudentController{}, "Get:Search")
	beego.Router("/student/school/datalist", &student.SchoolController{}, "Get,Post:DataList")
	beego.Router("/student/school/index", &student.SchoolController{}, "Get,Post:Index")
	beego.Router("/student/school/datagrid", &student.SchoolController{}, "Get,Post:DataGrid")
	beego.Router("/student/school/edit/?:id", &student.SchoolController{}, "Get,Post:Edit")
	beego.Router("/student/school/delete", &student.SchoolController{}, "Get,Post:Delete")
	beego.Router("/student/studentschool/datalist", &student.StudentSchoolController{}, "Get,Post:DataList")
	beego.Router("/student/studentschool/index", &student.StudentSchoolController{}, "Get,Post:Index")
	beego.Router("/student/studentschool/datagrid", &student.StudentSchoolController{}, "Get,Post:DataGrid")
	beego.Router("/student/studentschool/edit/?:id", &student.StudentSchoolController{}, "*:Edit")
	beego.Router("/student/studentschool/delete", &student.StudentSchoolController{}, "Post:Delete")
	beego.Router("/student/student/team", &student.StudentController{}, "*:Team")
	beego.Router("/student/student/teamdatagrid", &student.StudentController{}, "Get,Post:TeamDataGrid")
	beego.Router("/student/student/teamedit/?:id", &student.StudentController{}, "Get,Post:TeamEdit")
	beego.Router("/student/studentarea/datalist", &student.StudentAreaController{}, "Get,Post:DataList")
	beego.Router("/student/studentarea/index", &student.StudentAreaController{}, "*:Index")
	beego.Router("/student/studentarea/datagrid", &student.StudentAreaController{}, "Get,Post:DataGrid")
	beego.Router("/student/studentarea/edit/?:id", &student.StudentAreaController{}, "*:Edit")
	beego.Router("/student/studentarea/delete", &student.StudentAreaController{}, "Post:Delete")
	beego.Router("/student/contract/index", &student.ContractController{}, "*:Index")
	beego.Router("/student/contract/datagrid", &student.ContractController{}, "Get,Post:DataGrid")
	beego.Router("/student/contract/edit/?:id", &student.ContractController{}, "Get,Post:Edit")
	beego.Router("/student/contract/delete", &student.ContractController{}, "Get,Post:Delete")
	beego.Router("/student/contract/datalist", &student.ContractController{}, "Get,Post:DataList")
	beego.Router("/student/contract/print/?:id", &student.ContractController{}, "Get,Post:Print")
	beego.Router("/student/contractprice/index", &student.ContractPriceController{}, "*:Index")
	beego.Router("/student/contractprice/edit", &student.ContractPriceController{}, "*:Edit")
	beego.Router("/student/contractprice/search", &student.ContractPriceController{}, "Get,Post:Search")
	beego.Router("/student/student/total", &student.StudentController{}, "Get:Total")
	beego.Router("/student/student/totaldata", &student.StudentController{}, "Post:TotalData")
	beego.Router("/student/student/new", &student.StudentController{}, "Get:New")
	beego.Router("/student/student/newdata", &student.StudentController{}, "Post:NewData")
	beego.Router("/student/contract/total", &student.ContractController{}, "Get:Total")
	beego.Router("/student/contract/totaldata", &student.ContractController{}, "Post:TotalData")
	beego.Router("/student/contract/new", &student.ContractController{}, "Get:New")
	beego.Router("/student/contract/newdata", &student.ContractController{}, "Post:NewData")
	beego.Router("/student/contract/balance", &student.ContractController{}, "Get:Balance")
	beego.Router("/student/contract/balancegrid", &student.ContractController{}, "Post:BalanceGrid")
	/*student end*/
	/*teacher start*/
	beego.Router("/teacher/teacher/index", &teacher.TeacherController{}, "*:Index")
	beego.Router("/teacher/teacher/datagrid", &teacher.TeacherController{}, "Get,Post:DataGrid")
	beego.Router("/teacher/teacher/edit/?:id", &teacher.TeacherController{}, "Get,Post:Edit")
	beego.Router("/teacher/teacher/delete", &teacher.TeacherController{}, "Get,Post:Delete")
	beego.Router("/teacher/teacher/datalist", &teacher.TeacherController{}, "Get,Post:DataList")
	beego.Router("/teacher/subject/datalist", &teacher.SubjectController{}, "Get,Post:DataList")
	beego.Router("/teacher/subject/index", &teacher.SubjectController{}, "Get,Post:Index")
	beego.Router("/teacher/subject/datagrid", &teacher.SubjectController{}, "Get,Post:DataGrid")
	beego.Router("/teacher/subject/edit/?:id", &teacher.SubjectController{}, "Get,Post:Edit")
	beego.Router("/teacher/subject/delete", &teacher.SubjectController{}, "Get,Post:Delete")
	beego.Router("/teacher/teachersubject/edit/?:id", &teacher.TeacherSubjectController{}, "Get,Post:Edit")
	beego.Router("/teacher/teacher/salary", &teacher.TeacherController{}, "Get:Salary")
	beego.Router("/teacher/teacher/salarygrid", &teacher.TeacherController{}, "Post:SalaryGrid")
	beego.Router("/teacher/teacher/salarydetail/?:id", &teacher.TeacherController{}, "Get:SalaryDetail")
	beego.Router("/teacher/teacher/salarydetailgrid", &teacher.TeacherController{}, "Get,Post:SalaryDetailGrid")
	beego.Router("/teacher/teacher/useup", &teacher.TeacherController{}, "Get,Post:Useup")
	beego.Router("/teacher/teacher/useupgrid", &teacher.TeacherController{}, "Get,Post:UseupGrid")
	beego.Router("/teacher/teacher/userate", &teacher.TeacherController{}, "Get:Userate")
	beego.Router("/teacher/teacher/userategrid", &teacher.TeacherController{}, "Get,Post:UserateGrid")
	beego.Router("/teacher/expend/index", &teacher.ExpendController{}, "Get:Index")
	beego.Router("/teacher/expend/datagrid", &teacher.ExpendController{}, "Get,Post:DataGrid")
	beego.Router("/teacher/expend/edit/?:id", &teacher.ExpendController{}, "Get,Post:Edit")
	beego.Router("/teacher/expend/delete", &teacher.ExpendController{}, "Post:Delete")
	beego.Router("/teacher/expend/balance", &teacher.ExpendController{}, "Get:Balance")
	beego.Router("/teacher/expend/balancegrid", &teacher.ExpendController{}, "Get,Post:BalanceGrid")
	beego.Router("/teacher/expend/profit", &teacher.ExpendController{}, "Get:Profit")
	beego.Router("/teacher/expend/profitgrid", &teacher.ExpendController{}, "Get,Post:ProfitGrid")
	/*teacher end*/
	/*class start*/
	beego.Router("/class/schedule/index", &class.ScheduleController{}, "*:Index")
	beego.Router("/class/schedule/datagrid", &class.ScheduleController{}, "Get,Post:DataGrid")
	beego.Router("/class/schedule/edit", &class.ScheduleController{}, "Post:Edit")
	beego.Router("/class/schedule/clear", &class.ScheduleController{}, "Post:Clear")
	beego.Router("/class/classrecord/datalist", &class.ClassRecordController{}, "Get,Post:DataList")
	beego.Router("/class/classrecord/index", &class.ClassRecordController{}, "Get,Post:Index")
	beego.Router("/class/classrecord/datagrid", &class.ClassRecordController{}, "Get,Post:DataGrid")
	beego.Router("/class/classrecord/edit/?:id", &class.ClassRecordController{}, "Get,Post:Edit")
	beego.Router("/class/classrecord/delete", &class.ClassRecordController{}, "Get,Post:Delete")
	beego.Router("/class/classrecord/handle", &class.ClassRecordController{}, "Get,Post:Handle")
	beego.Router("/class/classrecord/confirm", &class.ClassRecordController{}, "Post:Confirm")
	beego.Router("/class/classrecord/cancle", &class.ClassRecordController{}, "Post:Cancle")
	beego.Router("/class/classrecord/single", &class.ClassRecordController{}, "Get,Post:Single")
	beego.Router("/class/classrecord/singleadd", &class.ClassRecordController{}, "Get,Post:SingleAdd")
	beego.Router("/class/classrecord/class", &class.ClassRecordController{}, "Get,Post:Class")
	beego.Router("/class/classrecord/classdata", &class.ClassRecordController{}, "Get,Post:ClassData")
	beego.Router("/class/classrecord/classgrade", &class.ClassRecordController{}, "Get:ClassGrade")
	beego.Router("/class/classrecord/classweek", &class.ClassRecordController{}, "Get:ClassWeek")
	beego.Router("/class/classrecord/classweekdata", &class.ClassRecordController{}, "Post:ClassWeekData")
	/*class end*/
	/*前端页面*/
	beego.Router("/frontedmanage/banner/index", &fronted_manage.BannerController{}, "Get:Index")
	beego.Router("/frontedmanage/banner/datagrid", &fronted_manage.BannerController{}, "Get,Post:DataGrid")
	beego.Router("/frontedmanage/banner/edit/?:id", &fronted_manage.BannerController{}, "Get,Post:Edit")
	beego.Router("/frontedmanage/banner/delete", &fronted_manage.BannerController{}, "Post:Delete")
	beego.Router("/frontedmanage/news/index", &fronted_manage.NewsController{}, "Get:Index")
	beego.Router("/frontedmanage/news/datagrid", &fronted_manage.NewsController{}, "Get,Post:DataGrid")
	beego.Router("/frontedmanage/news/edit/?:id", &fronted_manage.NewsController{}, "Get,Post:Edit")
	beego.Router("/frontedmanage/news/delete", &fronted_manage.NewsController{}, "Post:Delete")
	beego.Router("/frontedmanage/project/index", &fronted_manage.ProjectController{}, "Get:Index")
	beego.Router("/frontedmanage/project/datagrid", &fronted_manage.ProjectController{}, "Get,Post:DataGrid")
	beego.Router("/frontedmanage/project/edit/?:id", &fronted_manage.ProjectController{}, "Get,Post:Edit")
	beego.Router("/frontedmanage/project/delete", &fronted_manage.ProjectController{}, "Post:Delete")
	beego.Router("/frontedmanage/us/index", &fronted_manage.UsController{}, "Get:Index")
	beego.Router("/frontedmanage/us/datagrid", &fronted_manage.UsController{}, "Get,Post:DataGrid")
	beego.Router("/frontedmanage/us/edit/?:id", &fronted_manage.UsController{}, "Get,Post:Edit")
	beego.Router("/frontedmanage/us/delete", &fronted_manage.UsController{}, "Post:Delete")
	beego.Router("/frontedmanage/staff/index", &fronted_manage.StaffController{}, "Get:Index")
	beego.Router("/frontedmanage/staff/datagrid", &fronted_manage.StaffController{}, "Get,Post:DataGrid")
	beego.Router("/frontedmanage/staff/edit/?:id", &fronted_manage.StaffController{}, "Get,Post:Edit")
	beego.Router("/frontedmanage/staff/delete", &fronted_manage.StaffController{}, "Post:Delete")
	beego.Router("/frontedmanage/recruit/index", &fronted_manage.RecruitController{}, "Get:Index")
	beego.Router("/frontedmanage/recruit/datagrid", &fronted_manage.RecruitController{}, "Get,Post:DataGrid")
	beego.Router("/frontedmanage/recruit/edit/?:id", &fronted_manage.RecruitController{}, "Get,Post:Edit")
	beego.Router("/frontedmanage/recruit/delete", &fronted_manage.RecruitController{}, "Post:Delete")
	beego.Router("/frontedmanage/history/index", &fronted_manage.HistoryController{}, "Get:Index")
	beego.Router("/frontedmanage/history/datagrid", &fronted_manage.HistoryController{}, "Get,Post:DataGrid")
	beego.Router("/frontedmanage/history/edit/?:id", &fronted_manage.HistoryController{}, "Get,Post:Edit")
	beego.Router("/frontedmanage/history/delete", &fronted_manage.HistoryController{}, "Post:Delete")
	/*前端页面*/
	/*script start*/
	beego.Router("/script/genclass", &controllers.ScriptController{}, "*:GenClass")
	beego.Router("/script/studentstatus", &controllers.ScriptController{}, "*:StudentStatus")
	beego.Router("/script/teamstatus", &controllers.ScriptController{}, "*:TeamStatus")
	/*script end*/
	/*admin router start*/
	beego.Router("/test/index", &controllers.TestController{}, "*:Index")
	//用户角色路由
	beego.Router("/role/index", &controllers.RoleController{}, "*:Index")
	beego.Router("/role/datagrid", &controllers.RoleController{}, "Get,Post:DataGrid")
	beego.Router("/role/edit/?:id", &controllers.RoleController{}, "Get,Post:Edit")
	beego.Router("/role/delete", &controllers.RoleController{}, "Post:Delete")
	beego.Router("/role/datalist", &controllers.RoleController{}, "Post:DataList")
	beego.Router("/role/allocate", &controllers.RoleController{}, "Post:Allocate")
	beego.Router("/role/updateseq", &controllers.RoleController{}, "Post:UpdateSeq")

	//资源路由
	beego.Router("/resource/index", &controllers.ResourceController{}, "*:Index")
	beego.Router("/resource/treegrid", &controllers.ResourceController{}, "POST:TreeGrid")
	beego.Router("/resource/edit/?:id", &controllers.ResourceController{}, "Get,Post:Edit")
	beego.Router("/resource/parent", &controllers.ResourceController{}, "Post:ParentTreeGrid")
	beego.Router("/resource/delete", &controllers.ResourceController{}, "Post:Delete")
	//快速修改顺序
	beego.Router("/resource/updateseq", &controllers.ResourceController{}, "Post:UpdateSeq")

	//通用选择面板
	beego.Router("/resource/select", &controllers.ResourceController{}, "Get:Select")
	//用户有权管理的菜单列表（包括区域）
	beego.Router("/resource/usermenutree", &controllers.ResourceController{}, "POST:UserMenuTree")
	beego.Router("/resource/checkurlfor", &controllers.ResourceController{}, "POST:CheckUrlFor")

	//后台用户路由
	beego.Router("/backenduser/index", &controllers.BackendUserController{}, "*:Index")
	beego.Router("/backenduser/datagrid", &controllers.BackendUserController{}, "POST:DataGrid")
	beego.Router("/backenduser/edit/?:id", &controllers.BackendUserController{}, "Get,Post:Edit")
	beego.Router("/backenduser/delete", &controllers.BackendUserController{}, "Post:Delete")
	//后台用户中心
	beego.Router("/usercenter/profile", &controllers.UserCenterController{}, "Get:Profile")
	beego.Router("/usercenter/basicinfosave", &controllers.UserCenterController{}, "Post:BasicInfoSave")
	beego.Router("/usercenter/uploadimage", &controllers.UserCenterController{}, "Post:UploadImage")
	beego.Router("/usercenter/passwordsave", &controllers.UserCenterController{}, "Post:PasswordSave")

	beego.Router("/home/index", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/login", &controllers.HomeController{}, "*:Login")
	beego.Router("/home/dologin", &controllers.HomeController{}, "Post:DoLogin")
	beego.Router("/home/logout", &controllers.HomeController{}, "*:Logout")
	beego.Router("/home/ckeditorupload", &controllers.HomeController{}, "Post:CkeditorUpload")

	beego.Router("/home/404", &controllers.HomeController{}, "*:Page404")
	beego.Router("/home/error/?:error", &controllers.HomeController{}, "*:Error")

	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	/*admin router end*/

	beego.Router("/index", &fronted.MainController{}, "*:Index")
	beego.Router("/", &fronted.MainController{}, "*:Index")

	beego.Router("/news", &fronted.NewsController{}, "Get:List")
	beego.Router("/news/?:id", &fronted.NewsController{}, "Get:Detail")
	beego.Router("/us", &fronted.UsController{}, "Get:List")
	beego.Router("/project", &fronted.ProjectController{}, "Get:List")
	beego.Router("/staff", &fronted.StaffController{}, "Get:List")
	beego.Router("/recruit", &fronted.RecruitController{}, "Get:List")

	/*移动端*/
	beego.Router("/m", &fronted.MainController{}, "*:Mindex")
	beego.Router("/m/index", &fronted.MainController{}, "*:Mindex")
	beego.Router("/m/news", &fronted.NewsController{}, "*:Mlist")
	beego.Router("/m/news/?:id", &fronted.NewsController{}, "*:Mdetail")
	beego.Router("/m/project", &fronted.ProjectController{}, "*:Mlist")
	beego.Router("/m/recruit", &fronted.RecruitController{}, "*:Mlist")
	beego.Router("/m/us", &fronted.UsController{}, "*:Mlist")
	beego.Router("/m/staff", &fronted.StaffController{}, "*:Mlist")

	/*脚本*/
	beego.Router("/ext/genclass", &fronted.ExtController{}, "Get:GenClass")
	beego.Router("/ext/studentstatus", &fronted.ExtController{}, "Get:StudentStatus")
	beego.Router("/ext/teamstatus", &fronted.ExtController{}, "Get:TeamStatus")
}

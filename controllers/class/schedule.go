package class

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"gocms/controllers"
	"gocms/enums"
	"gocms/models"
	"gocms/models/class"
	"gocms/models/student"
	"gocms/models/teacher"
	"gocms/utils"
	"strconv"
	"strings"
	"time"
)

type ScheduleController struct {
	controllers.BaseController
}

func (c *ScheduleController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.CheckAuthor()
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
}

//DataList 课表列表
func (c *ScheduleController) DataList() {
	var params = class.ScheduleQueryParam{}
	//获取数据列表和总数
	data := class.ScheduleDataList(&params)
	//定义返回的数据结构
	c.JsonResult(enums.JRCodeSucc, "", data)
}

func (c *ScheduleController) Index() {
	utils.JDebug(c.CurUser)
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "schedule/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "schedule/index_footerjs.html"

	obj := make([]*class.Schedule, 0)
	student := make([]*student.Student, 0)
	teacher := make([]*teacher.Teacher, 0)
	subject := make([]*teacher.Subject, 0)
	queryobj := orm.NewOrm().QueryTable(models.ScheduleTBName()).Filter("status", class.ScheduleStatusValid)
	queryostu := orm.NewOrm().QueryTable(models.StudentTBName()).Filter("status", student.StudentStatusValid)
	queryotch := orm.NewOrm().QueryTable(models.TeacherTBName()).Filter("status", teacher.TeacherStatusValid)
	if c.CurUser.IsSuper != true {
		queryobj = queryobj.Filter("school_id", c.CurUser.SchoolId)
		queryostu = queryostu.Filter("school_id", c.CurUser.SchoolId)
		queryotch = queryotch.Filter("school_id", c.CurUser.SchoolId)
		if c.CurUser.RealName == "教师组" { // 如果是教师组，需查明是哪个教师
			singleTch := teacher.Teacher{}
			orm.NewOrm().QueryTable(models.TeacherTBName()).Filter("cardno", c.CurUser.UserName).One(&singleTch, "id")
			queryobj = queryobj.Filter("teacher_id", singleTch.Id)
			queryotch = queryotch.Filter("id", singleTch.Id)
		}
	}
	queryobj.All(&obj)
	queryostu.All(&student, "id", "name")
	queryotch.All(&teacher, "id", "name")
	orm.NewOrm().QueryTable(models.SubjectTBName()).All(&subject, "id", "name")
	studentMap1 := make(map[int]string)
	teacherMap := make(map[int]string)
	subjectMap := make(map[int]string)
	for _, val := range student {
		studentMap1[val.Id] = val.Name
	}
	for _, val := range teacher {
		teacherMap[val.Id] = val.Name
	}
	for _, val := range subject {
		subjectMap[val.Id] = val.Name
	}
	var final []class.ScheduleWithName
	for _, val := range obj {
		tmp := class.ScheduleWithName{Schedule: *val, Class10Name: studentMap1[val.Class10], Class20Name: studentMap1[val.Class20], Class30Name: studentMap1[val.Class30], Class40Name: studentMap1[val.Class40], Class50Name: studentMap1[val.Class50], Class60Name: studentMap1[val.Class60], Class70Name: studentMap1[val.Class70], Class80Name: studentMap1[val.Class80], Class90Name: studentMap1[val.Class90], Class100Name: studentMap1[val.Class100], Class110Name: studentMap1[val.Class110], Class120Name: studentMap1[val.Class120], Class130Name: studentMap1[val.Class130], Class140Name: studentMap1[val.Class140], Class150Name: studentMap1[val.Class150], Class160Name: studentMap1[val.Class160], Class170Name: studentMap1[val.Class170], Class180Name: studentMap1[val.Class180], Class190Name: studentMap1[val.Class190], Class200Name: studentMap1[val.Class200], Class210Name: studentMap1[val.Class210], Class220Name: studentMap1[val.Class220], Class230Name: studentMap1[val.Class230], Class240Name: studentMap1[val.Class240], Class250Name: studentMap1[val.Class250], Class260Name: studentMap1[val.Class260], Class270Name: studentMap1[val.Class270], Class280Name: studentMap1[val.Class280], Class290Name: studentMap1[val.Class290], Class300Name: studentMap1[val.Class300], Class310Name: studentMap1[val.Class310], Class320Name: studentMap1[val.Class320], Class330Name: studentMap1[val.Class330], Class340Name: studentMap1[val.Class340], Class350Name: studentMap1[val.Class350], Class360Name: studentMap1[val.Class360], Class370Name: studentMap1[val.Class370], Class380Name: studentMap1[val.Class380], Class390Name: studentMap1[val.Class390], Class400Name: studentMap1[val.Class400], Class410Name: studentMap1[val.Class410], Class420Name: studentMap1[val.Class420], Class430Name: studentMap1[val.Class430], Class440Name: studentMap1[val.Class440], Class450Name: studentMap1[val.Class450], Class460Name: studentMap1[val.Class460], Class470Name: studentMap1[val.Class470], Class480Name: studentMap1[val.Class480], Class490Name: studentMap1[val.Class490], Class500Name: studentMap1[val.Class500], Class510Name: studentMap1[val.Class510], Class520Name: studentMap1[val.Class520], Class530Name: studentMap1[val.Class530], Class540Name: studentMap1[val.Class540], Class550Name: studentMap1[val.Class550], Class560Name: studentMap1[val.Class560], TeacherName: teacherMap[val.TeacherId], SubjectName: subjectMap[val.SubjectId]}
		final = append(final, tmp)
	}
	c.Data["final"] = final
	//页面里按钮权限控制
	c.Data["canEdit"] = c.CheckActionAuthor("ScheduleController", "Edit")
	c.Data["canDelete"] = c.CheckActionAuthor("ScheduleController", "Delete")
	//c.Data["canAllocate"] = c.CheckActionAuthor("StudentAreaController", "Allocate")
}

// DataGrid 表格获取数据
func (c *ScheduleController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params class.ScheduleQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := class.SchedulePageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

// Edit 添加 编辑 页面
func (c *ScheduleController) Edit() {
	//如果是Post请求，则由Save处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := class.Schedule{}
	if Id > 0 {
		m.Id = Id
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			c.PageError("数据无效，请刷新后重试")
		}
	}
	c.Data["m"] = m
	c.SetTpl("schedule/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "schedule/edit_footerjs.html"
}

//Save 添加、编辑页面 保存
func (c *ScheduleController) Save() {
	var err error
	m := class.Schedule{}
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	m.UpdatedAt = time.Now().Format(utils.FormatDateTime)
	m.CreatedAt = time.Now().Format(utils.FormatDateTime)
	m.Status = class.ScheduleStatusValid
	o := orm.NewOrm()
	if m.Id == 0 {
		if _, err = o.Insert(&m); err == nil {
			c.JsonResult(enums.JRCodeSucc, "添加成功", m.Id)
		} else {
			c.JsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		}
	} else {
		if _, err = o.Update(&m); err == nil {
			c.JsonResult(enums.JRCodeSucc, "编辑成功", m.Id)
		} else {
			c.JsonResult(enums.JRCodeFailed, err.Error(), m.Id)
		}
	}
}

//Delete 批量删除
func (c *ScheduleController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	o := orm.NewOrm()
	num, err := o.QueryTable(models.ScheduleTBName()).Filter("id__in", ids).Delete()
	if err == nil {
		c.JsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.JsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

func (c *ScheduleController) Clear() {
	id, _ := c.GetInt("id")
	m := class.Schedule{Id: id}
	o := orm.NewOrm()
	o.Read(&m)

	m.Class10 = 0
	m.Class20 = 0
	m.Class30 = 0
	m.Class40 = 0
	m.Class50 = 0
	m.Class60 = 0
	m.Class70 = 0
	m.Class80 = 0
	m.Class90 = 0
	m.Class100 = 0
	m.Class110 = 0
	m.Class120 = 0
	m.Class130 = 0
	m.Class140 = 0
	m.Class150 = 0
	m.Class160 = 0
	m.Class170 = 0
	m.Class180 = 0
	m.Class190 = 0
	m.Class200 = 0
	m.Class210 = 0
	m.Class220 = 0
	m.Class230 = 0
	m.Class240 = 0
	m.Class250 = 0
	m.Class260 = 0
	m.Class270 = 0
	m.Class280 = 0
	m.Class290 = 0
	m.Class300 = 0
	m.Class310 = 0
	m.Class320 = 0
	m.Class330 = 0
	m.Class340 = 0
	m.Class350 = 0
	m.Class360 = 0
	m.Class370 = 0
	m.Class380 = 0
	m.Class390 = 0
	m.Class400 = 0
	m.Class410 = 0
	m.Class420 = 0
	m.Class430 = 0
	m.Class440 = 0
	m.Class450 = 0
	m.Class460 = 0
	m.Class470 = 0
	m.Class480 = 0
	m.Class490 = 0
	m.Class500 = 0
	m.Class510 = 0
	m.Class520 = 0
	m.Class530 = 0
	m.Class540 = 0
	m.Class550 = 0
	m.Class560 = 0
	if _, err := o.Update(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, err.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "清除成功", "")

}

func getStudentIds(schedule *class.Schedule) []int {
	var studentIdArr []int
	if schedule.Class10 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class10)
	}
	if schedule.Class20 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class20)
	}
	if schedule.Class30 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class30)
	}
	if schedule.Class40 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class40)
	}
	if schedule.Class50 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class50)
	}
	if schedule.Class60 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class60)
	}
	if schedule.Class70 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class70)
	}
	if schedule.Class80 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class80)
	}
	if schedule.Class90 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class90)
	}
	if schedule.Class100 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class100)
	}
	if schedule.Class110 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class110)
	}
	if schedule.Class120 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class120)
	}
	if schedule.Class130 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class130)
	}
	if schedule.Class140 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class140)
	}
	if schedule.Class150 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class150)
	}
	if schedule.Class160 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class160)
	}
	if schedule.Class170 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class170)
	}
	if schedule.Class180 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class180)
	}
	if schedule.Class190 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class190)
	}
	if schedule.Class200 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class200)
	}
	if schedule.Class210 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class210)
	}
	if schedule.Class220 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class220)
	}
	if schedule.Class230 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class230)
	}
	if schedule.Class240 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class240)
	}
	if schedule.Class250 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class250)
	}
	if schedule.Class260 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class260)
	}
	if schedule.Class270 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class270)
	}
	if schedule.Class280 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class280)
	}
	if schedule.Class290 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class290)
	}
	if schedule.Class300 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class300)
	}
	if schedule.Class310 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class310)
	}
	if schedule.Class320 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class320)
	}
	if schedule.Class330 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class330)
	}
	if schedule.Class340 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class340)
	}
	if schedule.Class350 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class350)
	}
	if schedule.Class360 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class360)
	}
	if schedule.Class370 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class370)
	}
	if schedule.Class380 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class380)
	}
	if schedule.Class390 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class390)
	}
	if schedule.Class400 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class400)
	}
	if schedule.Class410 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class410)
	}
	if schedule.Class420 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class420)
	}
	if schedule.Class430 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class430)
	}
	if schedule.Class440 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class440)
	}
	if schedule.Class450 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class450)
	}
	if schedule.Class460 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class460)
	}
	if schedule.Class470 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class470)
	}
	if schedule.Class480 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class480)
	}
	if schedule.Class490 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class490)
	}
	if schedule.Class500 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class500)
	}
	if schedule.Class510 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class510)
	}
	if schedule.Class520 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class520)
	}
	if schedule.Class530 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class530)
	}
	if schedule.Class540 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class540)
	}
	if schedule.Class550 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class550)
	}
	if schedule.Class560 > 0 {
		studentIdArr = append(studentIdArr, schedule.Class560)
	}

	return studentIdArr
}

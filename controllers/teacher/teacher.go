package teacher

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
)

type TeacherController struct {
	controllers.BaseController
}

func (c *TeacherController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.CheckAuthor()
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
}

func (c *TeacherController) Index() {
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "teacher/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "teacher/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.CheckActionAuthor("TeacherController", "Edit")
	c.Data["canDelete"] = c.CheckActionAuthor("TeacherController", "Delete")
	c.Data["canAllocate"] = c.CheckActionAuthor("TeacherController", "Allocate")
	c.Data["canSubject"] = c.CheckActionAuthor("TeacherSubjectController", "Edit")
}

// DataGrid 表格获取数据
func (c *TeacherController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params teacher.TeacherQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	params.CurUser = c.CurUser
	data, total := teacher.TeacherPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

//DataList 列表
func (c *TeacherController) DataList() {
	var params = teacher.TeacherQueryParam{}
	//获取数据列表和总数
	params.CurUser = c.CurUser
	data := teacher.TeacherDataList(&params)
	//定义返回的数据结构
	c.JsonResult(enums.JRCodeSucc, "", data)
}

func (c *TeacherController) Edit() {
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := teacher.Teacher{Id: Id}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.QueryTable(models.TeacherTBName()).Filter("id", Id).RelatedSel().One(&m)
		if err != nil {
			c.PageError("数据无效，请刷新后重试")
		}
	}
	//校区信息
	if m.School == nil {
		m.School = &student.School{}
	}
	c.Data["m"] = m
	c.SetTpl("teacher/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "teacher/edit_footerjs.html"
}

//Delete 批量删除
func (c *TeacherController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	if num, err := teacher.TeacherDelete(ids); err == nil {
		c.JsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.JsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

//Save 添加、编辑页面 保存
func (c *TeacherController) Save() {
	var err error
	m := teacher.Teacher{}
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	schoolId, _ := c.GetInt("SchoolId")
	school := student.School{Id: schoolId}
	m.School = &school
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
			c.JsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}
}

// 教师工资
func (c *TeacherController) Salary() {
	startDay, endDay := utils.MonthRange()
	c.Data["showMoreQuery"] = false
	c.Data["DateStart"] = startDay
	c.Data["DateEnd"] = endDay
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "teacher/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "teacher/salary_footerjs.html"
}

func (c *TeacherController) SalaryGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params class.ClassRecordQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	params.CurUser = c.CurUser
	data, _ := class.SalaryPageList(&params)
	handle := splitSlice(data)
	m := []*class.ClassRecord{}
	for _, val := range handle {
		var length, amount float64
		var tmp *class.ClassRecord
		for _, v := range val {
			tmp = v
			length += v.Length
			amount += v.Amount
			tmp.Length = length
			tmp.Amount = amount
		}
		m = append(m, tmp)
	}
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = len(m)
	result["rows"] = m
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *TeacherController) SalaryDetail() {
	id, _ := c.GetInt(":id")
	startDay, endDay := utils.MonthRange()
	c.Data["showMoreQuery"] = false
	c.Data["DateStart"] = startDay
	c.Data["DateEnd"] = endDay
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.Data["id"] = id
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "teacher/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "teacher/salarydetail_footerjs.html"
}

func (c *TeacherController) SalaryDetailGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params class.ClassRecordQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	params.CurUser = c.CurUser
	data, total := class.SalaryPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *TeacherController) Useup() {
	startDay, endDay := utils.MonthRange()
	c.Data["showMoreQuery"] = false
	c.Data["DateStart"] = startDay
	c.Data["DateEnd"] = endDay
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "teacher/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "teacher/useup_footerjs.html"
}

func (c *TeacherController) UseupGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params class.ClassRecordQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	params.CurUser = c.CurUser
	if c.CurUser.IsSuper != true {
		tch := teacher.Teacher{}
		err := orm.NewOrm().QueryTable(models.TeacherTBName()).Filter("cardno", c.CurUser.UserName).One(&tch)
		if err != nil {
			result := make(map[string]interface{})
			result["total"] = 0
			result["rows"] = []int{}
			c.Data["json"] = result
			c.ServeJSON()
		}
		params.TeacherId = tch.Id
	}
	data, _ := class.SalaryPageList(&params)
	handle := splitSlice(data)
	m := []*class.ClassRecord{}
	for _, val := range handle {
		var length, amount float64
		var tmp *class.ClassRecord
		for _, v := range val {
			tmp = v
			length += v.Length
			amount += v.Amount
			tmp.Length = length
			tmp.Amount = amount
		}
		m = append(m, tmp)
	}
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = len(m)
	result["rows"] = m
	c.Data["json"] = result
	c.ServeJSON()
}

//教师预排课
func (c *TeacherController) Userate() {
	startDay, endDay := utils.WeekRange()
	c.Data["showMoreQuery"] = false
	c.Data["DateStart"] = startDay
	c.Data["DateEnd"] = endDay
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "teacher/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "teacher/userate_footerjs.html"
}

func (c *TeacherController) UserateGrid() {
	var params class.ClassRecordQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	params.CurUser = c.CurUser
	data, _ := class.UserateDataList(&params)
	handle := splitSlice(data)
	type Userate struct {
		Class  class.ClassRecord
		Used   float64
		Before float64
	}
	m := []Userate{}
	for _, val := range handle {
		var used, before float64
		var tmp Userate
		for _, v := range val {
			tmp.Class = *v
			if v.Status == class.ClassRecordStatusConfirm {
				used += v.Length
			} else {
				before += v.Length
			}
			tmp.Used = used
			tmp.Before = before
		}
		m = append(m, tmp)
	}

	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = len(m)
	result["rows"] = m
	c.Data["json"] = result
	c.ServeJSON()
}

//按某个字段排序
//type sortByAge []*class.ClassRecord
//
//func (s sortByAge) Len() int           { return len(s) }
//func (s sortByAge) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
//func (s sortByAge) Less(i, j int) bool { return s[i].teacher.Id < s[j].teacher.Id}

//切片分组
func splitSlice(list []*class.ClassRecord) [][]*class.ClassRecord {
	//sort.Sort(sortByAge(list))
	returnData := make([][]*class.ClassRecord, 0)
	i := 0
	var j int
	for {
		if i >= len(list) {
			break
		}
		for j = i + 1; j < len(list) && list[i].Teacher.Id == list[j].Teacher.Id; j++ {
		}
		returnData = append(returnData, list[i:j])
		i = j
	}
	return returnData
}

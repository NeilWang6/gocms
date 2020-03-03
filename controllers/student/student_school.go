package student

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/cuua/gocms/controllers"
	"github.com/cuua/gocms/enums"
	"github.com/cuua/gocms/models"
	"github.com/cuua/gocms/models/student"
	"strconv"
	"strings"
)

type StudentSchoolController struct {
	controllers.BaseController
}

func (c *StudentSchoolController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.CheckAuthor()
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
}

//DataList 角色列表
func (c *StudentSchoolController) DataList() {
	var params = student.StudentSchoolQueryParam{}
	//获取数据列表和总数
	data := student.StudentSchoolDataList(&params)
	//定义返回的数据结构
	c.JsonResult(enums.JRCodeSucc, "", data)
}

func (c *StudentSchoolController) Index() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "studentschool/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "studentschool/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.CheckActionAuthor("StudentSchoolController", "Edit")
	c.Data["canDelete"] = c.CheckActionAuthor("StudentSchoolController", "Delete")
	//c.Data["canAllocate"] = c.CheckActionAuthor("StudentAreaController", "Allocate")
}

// DataGrid 表格获取数据
func (c *StudentSchoolController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params student.StudentSchoolQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := student.StudentSchoolPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

// Edit 添加 编辑 页面
func (c *StudentSchoolController) Edit() {
	//如果是Post请求，则由Save处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := student.StudentSchool{}
	if Id > 0 {
		m.Id = Id
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			c.PageError("数据无效，请刷新后重试")
		}
	}
	c.Data["m"] = m
	c.SetTpl("studentschool/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "studentschool/edit_footerjs.html"
}

//Save 添加、编辑页面 保存
func (c *StudentSchoolController) Save() {
	var err error
	m := student.StudentSchool{}
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
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

//Delete 批量删除
func (c *StudentSchoolController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	o := orm.NewOrm()
	num, err := o.QueryTable(models.StudentSchoolTBName()).Filter("id__in", ids).Delete()
	if err == nil {
		c.JsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.JsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

package teacher

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"gocms/controllers"
	"gocms/enums"
	"gocms/models"
	"gocms/models/teacher"
)

type TeacherSubjectController struct {
	controllers.BaseController
}

func (c *TeacherSubjectController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.CheckAuthor()
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
}

// Edit 添加 编辑 页面
func (c *TeacherSubjectController) Edit() {
	//如果是Post请求，则由Save处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := []teacher.TeacherSubject{}
	orm.NewOrm().QueryTable(models.TeacherSubjectTBName()).Filter("teacher_id", Id).All(&m)
	var subjectIds []int
	for _, val := range m {
		subjectIds = append(subjectIds, val.SubjectId)
	}
	jsonId, _ := json.Marshal(subjectIds)
	c.Data["subjectIds"] = string(jsonId)
	c.Data["TeacherId"] = Id
	c.SetTpl("teachersubject/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "teachersubject/edit_footerjs.html"
}

//Save 添加、编辑页面 保存
func (c *TeacherSubjectController) Save() {
	m := teacher.TeacherSubject{}
	var SubjectId []int
	c.Ctx.Input.Bind(&SubjectId, "SubjectId")
	teacherId, _ := c.GetInt("TeacherId")
	if teacherId <= 0 {
		c.JsonResult(enums.JRCodeFailed, "添加失败", teacherId)
	}
	o := orm.NewOrm()
	subjectMap := make(map[int]bool)
	for _, val := range SubjectId { // 添加
		m.Id = 0
		m.SubjectId = val
		m.TeacherId = teacherId
		o.ReadOrCreate(&m, "teacher_id", "subject_id")
		subjectMap[val] = true
	}
	obj := []teacher.TeacherSubject{}
	o.QueryTable(models.TeacherSubjectTBName()).Filter("teacher_id", teacherId).All(&obj)
	for _, val := range obj { // 删除
		if _, ok := subjectMap[val.SubjectId]; !ok { // 该数据是不需要的
			o.Delete(&val)
		}
	}
	c.JsonResult(enums.JRCodeSucc, "修改成功", teacherId)
}

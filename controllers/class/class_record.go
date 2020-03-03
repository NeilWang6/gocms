package class

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/cuua/gocms/controllers"
	Student2 "github.com/cuua/gocms/controllers/student"
	"github.com/cuua/gocms/enums"
	"github.com/cuua/gocms/models"
	"github.com/cuua/gocms/models/class"
	"github.com/cuua/gocms/models/student"
	"github.com/cuua/gocms/models/teacher"
	"github.com/cuua/gocms/utils"
	"strconv"
	"strings"
	"time"
)

type ClassRecordController struct {
	controllers.BaseController
}

func (c *ClassRecordController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.CheckAuthor()
}

//DataList 学校列表
func (c *ClassRecordController) DataList() {
	var params = class.ClassRecordQueryParam{}
	//获取数据列表和总数
	params.CurUser = c.CurUser
	data := class.ClassRecordDataList(&params)
	//定义返回的数据结构
	c.JsonResult(enums.JRCodeSucc, "", data)
}

func (c *ClassRecordController) Index() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "classrecord/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "classrecord/index_footerjs.html"
	//页面里按钮权限控制
	//c.Data["canEdit"] = c.CheckActionAuthor("ClassRecordController", "Edit")
	//c.Data["canDelete"] = c.CheckActionAuthor("ClassRecordController", "Delete")
	//c.Data["canAllocate"] = c.CheckActionAuthor("StudentAreaController", "Allocate")
}

// DataGrid 表格获取数据
func (c *ClassRecordController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params class.ClassRecordQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	params.CurUser = c.CurUser
	data, total := class.ClassRecordPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

// Edit 添加 编辑 页面
func (c *ClassRecordController) Edit() {
	//如果是Post请求，则由Save处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := class.ClassRecord{}
	if Id > 0 {
		m.Id = Id
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			c.PageError("数据无效，请刷新后重试")
		}
	}
	c.Data["m"] = m
	c.SetTpl("classrecord/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "classrecord/edit_footerjs.html"
}

//Save 添加、编辑页面 保存
func (c *ClassRecordController) Save() {
	var err error
	m := class.ClassRecord{}
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
func (c *ClassRecordController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	o := orm.NewOrm()
	num, err := o.QueryTable(models.ClassRecordTBName()).Filter("id__in", ids).Delete()
	if err == nil {
		c.JsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.JsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

//
func (c *ClassRecordController) Handle() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "classrecord/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "classrecord/handle_footerjs.html"
	//页面里按钮权限控制
	c.Data["canConfirm"] = c.CheckActionAuthor("ClassRecordController", "Confirm")
	c.Data["canCancle"] = c.CheckActionAuthor("ClassRecordController", "Cancle")
	//c.Data["canDelete"] = c.CheckActionAuthor("ClassRecordController", "Delete")
}

// 确认上课
func (c *ClassRecordController) Confirm() {
	id, _ := c.GetInt("id")
	m := class.ClassRecord{Id: id}
	orm.NewOrm().Read(&m)
	choose, err := student.ChooseContract(m.Student.Id, m.Type, m.Length)
	if err != nil {
		c.JsonResult(enums.JRCodeFailed, err.Error(), "")
	}
	studentm := student.Student{Id: m.Student.Id}
	err = orm.NewOrm().Read(&studentm)
	if err != nil {
		c.JsonResult(enums.JRCodeFailed, err.Error(), "")
	}
	teacherPrice, err := teacher.TeacherPrice(m.Teacher.Id, studentm.Grade, m.Type)
	if err != nil {
		c.JsonResult(enums.JRCodeFailed, err.Error(), "")
	}
	// 事物处理
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		c.JsonResult(enums.JRCodeFailed, "请重试", "")
	}
	var leftLength float64 = 0
	var stuAmount float64 = 0
	var firstErr, secondErr, stuErr, classErr error
	var stuPrice, contractIdArr []int
	for _, val := range choose { // 先减合同余额
		contractIdArr = append(contractIdArr, val.Id) // 合同记录
		stuPrice = append(stuPrice, val.Price)
		if val.SurplusQuantity < m.Length {
			stuAmount += val.SurplusQuantity * float64(val.Price)
			leftLength = m.Length - val.SurplusQuantity
			val.SurplusQuantity = 0
			val.Surplus = 0
			_, firstErr = o.Update(&val, "surplus_quantity", "surplus")
		} else {
			val.SurplusQuantity = val.SurplusQuantity - (m.Length - leftLength)
			val.Surplus = val.Surplus - (float64(val.Price) * (m.Length - leftLength))
			_, secondErr = o.Update(&val, "surplus_quantity", "surplus")
			stuAmount += (m.Length - leftLength) * float64(val.Price)
		}
	}
	contractIdStr, _ := json.Marshal(contractIdArr)
	stuPriceStr, _ := json.Marshal(stuPrice)
	m.Status = class.ClassRecordStatusConfirm
	m.Price = teacherPrice
	m.Grade = studentm.Grade
	m.Amount = float64(teacherPrice) * m.Length
	m.StudentAmount = stuAmount
	m.StudentPrice = string(stuPriceStr)
	m.ContractId = string(contractIdStr)
	switch m.Type { // 扣学生余额, 余额及状态更改
	case student.ContractTypeYi:
		m.StudentSurplus = studentm.Balance1 - stuAmount
		m.StudentQuantitySurplus = studentm.Balance1Length - m.Length
		studentm.Balance1 = studentm.Balance1 - stuAmount
		studentm.Balance1Length = studentm.Balance1Length - m.Length
		_, classErr = o.Update(&m)
		_, stuErr = o.Update(&studentm, "balance1_length", "balance1")
	case student.ContractTypeXiao:
		m.StudentSurplus = studentm.Balance2 - stuAmount
		m.StudentQuantitySurplus = studentm.Balance2Length - m.Length
		studentm.Balance2 = studentm.Balance2 - stuAmount
		studentm.Balance2Length = studentm.Balance2Length - m.Length
		_, stuErr = o.Update(&studentm, "balance2_length", "balance2")
		_, classErr = o.Update(&m)
	case student.ContractTypeTuo:
		m.StudentSurplus = studentm.Balance3 - stuAmount
		m.StudentQuantitySurplus = studentm.Balance3Length - m.Length
		studentm.Balance3 = studentm.Balance3 - stuAmount
		studentm.Balance3Length = studentm.Balance3Length - m.Length
		_, stuErr = o.Update(&studentm, "balance3_length", "balance3")
		_, classErr = o.Update(&m)
	default:
		stuErr = errors.New("找不到该学生")
	}
	if firstErr != nil || secondErr != nil || stuErr != nil || classErr != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	if err != nil {
		c.JsonResult(enums.JRCodeFailed, "数据有误，请联系管理员", "")
	}
	c.JsonResult(0, "操作成功", "")
}

// 取消记录
func (c *ClassRecordController) Cancle() {
	id, _ := c.GetInt("id")
	_, err := orm.NewOrm().QueryTable(models.ClassRecordTBName()).Filter("id", id).Update(orm.Params{
		"status": class.ClassRecordStatusDelete,
	})
	if err != nil {
		c.JsonResult(enums.JRCodeFailed, err.Error(), "")
	}
	c.JsonResult(0, "成功", "")
}

// 排课管理
func (c *ClassRecordController) Single() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "classrecord/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "classrecord/single_footerjs.html"
}

func (c *ClassRecordController) SingleAdd() {
	m := class.ClassRecord{}
	if err := c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", "")
	}
	teacherId, _ := c.GetInt("TeacherId")
	studentId, _ := c.GetInt("StudentId")
	subjectId, _ := c.GetInt("SubjectId")
	teacherm := teacher.Teacher{Id: teacherId}
	student := student.Student{Id: studentId}
	subject := teacher.Subject{Id: subjectId}
	m.Teacher = &teacherm
	m.Student = &student
	m.Subject = &subject
	valid := validation.Validation{}
	valid.Required(m.Student.Id, "学生")
	valid.Required(m.Subject.Id, "科目")
	valid.Required(m.Teacher.Id, "教师")
	valid.Required(m.Time, "时段")
	valid.Required(m.Length, "课时长")
	valid.Required(m.Date, "日期")
	if valid.HasErrors() {
		c.JsonResult(enums.JRCodeFailed, valid.Errors[0].Key+valid.Errors[0].Message, "")
	}
	/*排课只允许本周内*/
	monday, sunday := utils.WeekRange()
	if m.Date < monday || m.Date > sunday {
		c.JsonResult(enums.JRCodeFailed, "只可以排本周的课程", "")
	}
	m.CreatedAt = time.Now().Format(utils.FormatDateTime)
	_, _, err := orm.NewOrm().ReadOrCreate(&m, "teacher_id", "student_id", "subject_id", "date", "time", "length")
	if err != nil {
		c.JsonResult(enums.JRCodeFailed, err.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "添加成功", "")
}

func (c *ClassRecordController) Class() {
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "classrecord/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "classrecord/class_footerjs.html"
}

func (c *ClassRecordController) ClassData() {
	dateStart := c.GetString("DateStart")
	dateEnd := c.GetString("DateEnd")
	schoolId := c.GetString("SchoolId")
	group := c.GetString("group")
	count := 0.00
	length := 0.00
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("sum(student_amount) amount", "sum(length) length", "type", "grade").From(models.ClassRecordTBName()).Where("status = ?")
	if dateStart != "" {
		qb = qb.And("DATE_FORMAT(date, '%Y-%m-%d') >= '" + dateStart + "'")
	}
	if dateEnd != "" {
		qb = qb.And("DATE_FORMAT(date, '%Y-%m-%d') <= '" + dateEnd + "'")
	}
	if schoolId != "" {
		qb = qb.And("school_id = " + schoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql := qb.GroupBy(group).String()
	type Foo struct {
		Amount float64
		Type   int
		Length float64
		Grade  int
	}
	m := make([]*Foo, 0)
	orm.NewOrm().Raw(sql, class.ClassRecordStatusConfirm).QueryRows(&m)
	obj := make([]Student2.Flot, 0)
	obj1 := make([]Student2.Flot, 0)
	tmp := Student2.Flot{}
	tmp1 := Student2.Flot{}
	for key, val := range m {
		count += val.Amount
		length += val.Length
		switch group {
		case "type":
			tmp.Label = student.ContractTypeMap[val.Type]
			tmp1.Label = student.ContractTypeMap[val.Type]
		default:
			tmp.Label = student.GetGradeName(val.Grade)
			tmp1.Label = student.GetGradeName(val.Grade)
		}
		tmp.Data = int(val.Amount)
		tmp.Color = utils.GetColor(key)
		tmp1.Data = int(val.Length)
		tmp1.Color = utils.GetColor(key)
		obj = append(obj, tmp)
		obj1 = append(obj1, tmp1)
	}
	ret := map[string]interface{}{"count": count, "length": length, "obj": obj, "obj1": obj1}
	c.JsonResult(enums.JRCodeSucc, "succ", ret)
}

func (c *ClassRecordController) ClassGrade() {
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "classrecord/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "classrecord/classgrade_footerjs.html"
}

func (c *ClassRecordController) ClassWeek() {
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "classrecord/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "classrecord/classweek_footerjs.html"
}

func (c *ClassRecordController) ClassWeekData() {
	dateStart := c.GetString("DateStart")
	dateEnd := c.GetString("DateEnd")
	schoolId := c.GetString("SchoolId")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("DATE_FORMAT(created_at,'%Y%u') weeks", "sum(amount) amount", "sum(length) length").From(models.ClassRecordTBName()).Where("status = ?")
	if dateStart != "" {
		qb = qb.And("DATE_FORMAT(created_at, '%Y-%m-%d') >= '" + dateStart + "'")
	}
	if dateEnd != "" {
		qb = qb.And("DATE_FORMAT(created_at, '%Y-%m-%d') <= '" + dateEnd + "'")
	}
	if schoolId != "" {
		qb = qb.And("school_id = " + schoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql := qb.GroupBy("weeks").String()
	type NewTotal struct {
		Weeks  string
		Amount float64
		Length float64
	}
	m := make([]NewTotal, 0)
	amount, length := 0.00, 0.00
	amountObj := [][]interface{}{}
	lengthObj := [][]interface{}{}
	orm.NewOrm().Raw(sql, class.ClassRecordStatusConfirm).QueryRows(&m)
	for _, val := range m {
		amount += val.Amount
		length += val.Length
		tmp := []interface{}{}
		tmp1 := []interface{}{}
		tmp = append(tmp, val.Weeks)
		tmp = append(tmp, val.Amount)
		tmp1 = append(tmp1, val.Weeks)
		tmp1 = append(tmp1, val.Length)
		amountObj = append(amountObj, tmp)
		lengthObj = append(lengthObj, tmp)
	}
	c.JsonResult(enums.JRCodeSucc, "succ", map[string]interface{}{"lengthObj": lengthObj, "amountObj": amountObj, "amount": amount, "length": length})
}

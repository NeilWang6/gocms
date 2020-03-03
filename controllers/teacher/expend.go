package teacher

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/cuua/gocms/controllers"
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

type ExpendController struct {
	controllers.BaseController
}

func (c *ExpendController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.CheckAuthor()
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
}

func (c *ExpendController) Index() {
	firstDate := utils.GetFirstDateOfMonth(time.Now())
	c.Data["firstDate"] = firstDate
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "expend/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "expend/index_footerjs.html"
	c.Data["canEdit"] = c.CheckActionAuthor("ExpendController", "Edit")
	c.Data["canDelete"] = c.CheckActionAuthor("ExpendController", "Delete")
}

func (c *ExpendController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params teacher.ExpendQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := teacher.ExpendPageList(&params)
	totalM := &teacher.Expend{Name: "合计"}
	for _, val := range data {
		totalM.Amount += val.Amount
	}
	total++
	data = append(data, totalM)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

// Edit 添加 编辑 页面
func (c *ExpendController) Edit() {
	//如果是Post请求，则由Save处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := &teacher.Expend{}
	if Id > 0 {
		m.Id = Id
		o := orm.NewOrm()
		err := o.Read(m)
		if err != nil {
			c.PageError("数据无效，请刷新后重试")
		}
	}
	c.Data["m"] = m
	c.SetTpl("expend/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "expend/edit_footerjs.html"
}

//Save 添加、编辑页面 保存
func (c *ExpendController) Save() {
	var err error
	m := &teacher.Expend{}
	//获取form里的值
	if err = c.ParseForm(m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	startDate := utils.GetFirstDateOfMonth(time.Now())
	m.Month = startDate.Format("2006-01-02")
	o := orm.NewOrm()
	if m.Id == 0 {
		if _, err = o.Insert(m); err == nil {
			c.JsonResult(enums.JRCodeSucc, "添加成功", m.Id)
		} else {
			c.JsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		}

	} else {
		if _, err = o.Update(m); err == nil {
			c.JsonResult(enums.JRCodeSucc, "编辑成功", m.Id)
		} else {
			c.JsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}
}

//Delete 批量删除
func (c *ExpendController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	o := orm.NewOrm()
	num, err := o.QueryTable(models.ExpendTBName()).Filter("id__in", ids).Delete()
	if err == nil {
		c.JsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.JsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

func (c *ExpendController) Balance() {
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "expend/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "expend/balance_footerjs.html"
}

func (c *ExpendController) BalanceGrid() {
	var params teacher.ExpendQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	date := params.DateStart
	var startDate, endDate time.Time
	if date == "" {
		startDate = utils.GetFirstDateOfMonth(time.Now())
		endDate = utils.GetLastDateOfMonth(time.Now())
	} else {
		startDate, _ = time.Parse(utils.FormatDateTime, date+" 00:00:00")
		startDate = utils.GetFirstDateOfMonth(startDate)
		endDate = utils.GetLastDateOfMonth(startDate)
	}
	monthFirst := startDate.Format("2006-01-02")
	monthEnd := endDate.Format("2006-01-02")
	// 支出
	type Ex struct {
		Amount float64
	}
	expend := &Ex{}
	qb, _ := orm.NewQueryBuilder("mysql")
	qb = qb.Select("sum(amount) amount").From(models.ExpendTBName()).Where("month = ?")
	if params.SchoolId != "" {
		qb = qb.And("school_id = " + params.SchoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql := qb.String()
	orm.NewOrm().Raw(sql, monthFirst).QueryRow(expend)
	// 课耗
	student := &Ex{}
	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("sum(student_amount) amount").From(models.ClassRecordTBName()).Where("date >= ?").And("date <= ?").And("status = ?").String()
	orm.NewOrm().Raw(sql, monthFirst, monthEnd, class.ClassRecordStatusConfirm).QueryRow(student)
	// 教师工资
	teacher := &Ex{}
	qb, _ = orm.NewQueryBuilder("mysql")
	qb = qb.Select("sum(student_amount) amount").From(models.ClassRecordTBName()).Where("date >= ?").And("date <= ?").And("status = ?")
	if params.SchoolId != "" {
		qb = qb.And("school_id = " + params.SchoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql = qb.GroupBy("teacher_id", "subject_id", "date", "time").String()
	orm.NewOrm().Raw(sql, monthFirst, monthEnd, class.ClassRecordStatusConfirm).QueryRow(teacher)
	balanceTotal := student.Amount - expend.Amount - teacher.Amount
	data := map[string]interface{}{"expend": expend.Amount, "student": student.Amount, "teacher": teacher.Amount, "balanceTotal": balanceTotal}
	row := []interface{}{data}
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = 1
	result["rows"] = row
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *ExpendController) Profit() {
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "expend/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "expend/profit_footerjs.html"
}

func (c *ExpendController) ProfitGrid() {
	var params teacher.ExpendQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	date := params.DateStart
	var startDate, endDate time.Time
	if date == "" {
		startDate = utils.GetFirstDateOfMonth(time.Now())
		endDate = utils.GetLastDateOfMonth(time.Now())
	} else {
		startDate, _ = time.Parse(utils.FormatDateTime, date+" 00:00:00")
		startDate = utils.GetFirstDateOfMonth(startDate)
		endDate = utils.GetLastDateOfMonth(startDate)
	}
	monthFirst := startDate.Format("2006-01-02")
	monthEnd := endDate.Format("2006-01-02") + " 23:59:59" //
	// 支出
	type Ex struct {
		Amount float64
	}
	expend := &Ex{}
	qb, _ := orm.NewQueryBuilder("mysql")
	qb = qb.Select("sum(amount) amount").From(models.ExpendTBName()).Where("month = ?")
	if params.SchoolId != "" {
		qb = qb.And("school_id = " + params.SchoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql := qb.String()
	orm.NewOrm().Raw(sql, monthFirst).QueryRow(expend)
	// 合同金额
	contract := &Ex{}
	qb, _ = orm.NewQueryBuilder("mysql")
	qb = qb.Select("sum(amount) amount").From(models.ContractTBName() + " AS contract").LeftJoin(models.StudentTBName() + " AS student").On("student.id = contract.student_id").Where("contract.created_at >= ?").And("contract.created_at <= ?").And("contract.status in (?,?)")
	if params.SchoolId != "" {
		qb = qb.And("contract.school_id = " + params.SchoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("contract.school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql = qb.String()
	orm.NewOrm().Raw(sql, monthFirst, monthEnd, student.ContractStatusValid, student.ContractStatusBack).QueryRow(contract)
	// 退费金额
	refund := &Ex{}
	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("sum(surplus) amount").From(models.ContractTBName()).Where("updated_at >= ?").And("updated_at <= ?").And("status = ?").String()
	orm.NewOrm().Raw(sql, monthFirst, monthEnd, student.ContractStatusBack).QueryRow(refund)
	// post机金额
	pos := &Ex{}
	qb, _ = orm.NewQueryBuilder("mysql")
	qb = qb.Select("sum(amount) amount").From(models.ContractTBName() + " AS contract").LeftJoin(models.StudentTBName() + " AS student").On("student.id = contract.student_id").Where("contract.created_at >= ?").And("contract.created_at <= ?").And("contract.status in (?,?)").And("contract.payment = ?")
	if params.SchoolId != "" {
		qb = qb.And("contract.school_id = " + params.SchoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("contract.school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql = qb.String()
	orm.NewOrm().Raw(sql, monthFirst, monthEnd, student.ContractStatusValid, student.ContractStatusBack, "pos").QueryRow(pos)
	profitTotal := contract.Amount - refund.Amount - expend.Amount - pos.Amount
	data := map[string]interface{}{"expend": expend.Amount, "contract": contract.Amount, "refund": refund.Amount, "pos": pos.Amount, "profitTotal": profitTotal}
	row := []interface{}{data}
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = 1
	result["rows"] = row
	c.Data["json"] = result
	c.ServeJSON()
}

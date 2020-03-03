package student

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/cuua/gocms/controllers"
	"github.com/cuua/gocms/enums"
	"github.com/cuua/gocms/models"
	"github.com/cuua/gocms/models/student"
	"github.com/cuua/gocms/utils"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type ContractController struct {
	controllers.BaseController
}

func (c *ContractController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.CheckAuthor()
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
}

//DataList 学校列表
func (c *ContractController) DataList() {
	var params = student.ContractQueryParam{}
	//获取数据列表和总数
	params.CurUser = c.CurUser
	data := student.ContractDataList(&params)
	//定义返回的数据结构
	c.JsonResult(enums.JRCodeSucc, "", data)
}

func (c *ContractController) Index() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "contract/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "contract/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.CheckActionAuthor("ContractController", "Edit")
	c.Data["canDelete"] = c.CheckActionAuthor("ContractController", "Delete")
	//c.Data["canAllocate"] = c.CheckActionAuthor("StudentAreaController", "Allocate")
}

// DataGrid 表格获取数据
func (c *ContractController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params student.ContractQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	params.CurUser = c.CurUser
	data, total := student.ContractPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

// Edit 添加 编辑 页面
func (c *ContractController) Edit() {
	//如果是Post请求，则由Save处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := student.Contract{}
	if Id > 0 {
		m.Id = Id
		err := orm.NewOrm().QueryTable(models.ContractTBName()).RelatedSel("student").Filter("id", Id).One(&m)
		if err != nil {
			c.PageError("数据无效，请刷新后重试")
		}
	}
	if m.Student == nil {
		studentm := student.Student{}
		schoolm := student.School{}
		studentm.School = &schoolm
		m.Student = &studentm
	}
	c.Data["m"] = m
	c.Data["school_id"] = c.CurUser.SchoolId
	c.SetTpl("contract/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "contract/edit_footerjs.html"
}

//Save 添加、编辑页面 保存
func (c *ContractController) Save() {
	var err error
	m := student.Contract{}
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	studentId, _ := c.GetInt("StudentId", 0)
	studentm := student.Student{Id: studentId}
	orm.NewOrm().Read(&studentm)
	m.Student = &studentm

	valid := validation.Validation{}
	valid.Required(m.Student.Id, "学生")
	valid.Range(m.Type, 0, 2, "类型")
	valid.Required(m.Price, "单价")
	valid.Required(m.Quantity, "课时长")
	valid.Required(m.Payment, "支付方式")
	valid.Range(m.Status, 0, 2, "状态")
	if valid.HasErrors() {
		c.JsonResult(enums.JRCodeFailed, valid.Errors[0].Key+valid.Errors[0].Message, "")
	}
	var price student.ContractPrice
	sprice, err := price.Search(m.Type, studentm.Grade, m.Quantity)
	if err == nil {
		priceCol := "Grade" + strconv.Itoa(studentm.Grade)
		immutable := reflect.ValueOf(sprice)
		val := immutable.FieldByName(priceCol).Int()
		m.Price = int(val)
	}
	m.Amount = float64(m.Price) * m.Quantity
	m.Surplus = m.Amount
	m.SurplusQuantity = m.Quantity
	m.CreatedAt = time.Now()
	o := orm.NewOrm()
	// 这里有几种状态 正常禁用退费
	var conErr, stuErr error
	if m.Id == 0 {
		o.Begin()
		switch m.Status {
		case 0: // 正常情况，需增加学生余额
			_, conErr = o.Insert(&m)
			switch m.Type {
			case student.ContractTypeYi:
				studentm.Balance1 = studentm.Balance1 + m.Amount
				studentm.Balance1Length = studentm.Balance1Length + m.Quantity
			case student.ContractTypeXiao:
				studentm.Balance2 = studentm.Balance1 - m.Amount
				studentm.Balance2Length = studentm.Balance1Length + m.Quantity
			case student.ContractTypeTuo:
				studentm.Balance3 = studentm.Balance1 - m.Amount
				studentm.Balance3Length = studentm.Balance1Length + m.Quantity
			}
			_, stuErr = o.Update(&studentm)
		case 1:
			_, conErr = o.Insert(&m)
		case 2:
			_, conErr = o.Insert(&m)
		}
		if conErr != nil || stuErr != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
		}
		if err == nil {
			c.JsonResult(enums.JRCodeSucc, "添加成功", m.Id)
		} else {
			c.JsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		}
	} else {
		chkm := student.Contract{Id: m.Id}
		orm.NewOrm().Read(&chkm)
		if chkm.Status != 0 {
			c.JsonResult(enums.JRCodeFailed, "不允许修改", "")
		}
		o.Begin()
		if m.Status == 2 { // 表示退费，需扣除学生余额
			_, conErr = o.Update(&m, "status")
			switch m.Type {
			case student.ContractTypeYi:
				studentm.Balance1 = studentm.Balance1 - m.Amount
				studentm.Balance1Length = studentm.Balance1Length - m.Quantity
			case student.ContractTypeXiao:
				studentm.Balance2 = studentm.Balance2 - m.Amount
				studentm.Balance2Length = studentm.Balance2Length - m.Quantity
			case student.ContractTypeTuo:
				studentm.Balance3 = studentm.Balance3 - m.Amount
				studentm.Balance3Length = studentm.Balance3Length - m.Quantity
			}
			_, stuErr = o.Update(&studentm)
		} else {
			_, conErr = o.Update(&m, "status")
		}
		if conErr != nil || stuErr != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
		}
		if err == nil { // 合同只允许修改状态
			c.JsonResult(enums.JRCodeSucc, "编辑成功", m.Id)
		} else {
			fmt.Println(err)
			c.JsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}
}

//Delete 批量删除
func (c *ContractController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	o := orm.NewOrm()
	num, err := o.QueryTable(models.ContractTBName()).Filter("id__in", ids).Delete()
	if err == nil {
		c.JsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.JsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

func (c *ContractController) Total() {
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "contract/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "contract/total_footerjs.html"
}

func (c *ContractController) TotalData() {
	dateStart := c.GetString("DateStart")
	dateEnd := c.GetString("DateEnd")
	schoolId := c.GetString("SchoolId")
	count := 0.0
	backCount := 0.0
	// 正常
	qb, _ := orm.NewQueryBuilder("mysql")
	qb = qb.Select("sum(amount) as amount", "contract.type type").From(models.ContractTBName() + " AS contract").LeftJoin(models.StudentTBName() + " AS student").On("student.id = contract.student_id").Where("contract.status in (?,?)")
	if dateStart != "" {
		qb = qb.And("DATE_FORMAT(contract.created_at, '%Y-%m-%d') >= '" + dateStart + "'")
	}
	if dateEnd != "" {
		qb = qb.And("DATE_FORMAT(contract.created_at, '%Y-%m-%d') <= '" + dateEnd + "'")
	}
	if schoolId != "" {
		qb = qb.And("student.school_id = " + schoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("student.school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql := qb.GroupBy("contract.type").OrderBy("amount DESC").String()
	type Foo struct {
		Amount float64
		Type   int
	}
	m := make([]*Foo, 0)
	orm.NewOrm().Raw(sql, student.ContractStatusValid, student.ContractStatusBack).QueryRows(&m)
	tmp := Flot{}
	validObj := make([]Flot, 0)
	for key, val := range m {
		count += val.Amount
		tmp.Label = student.ContractTypeMap[val.Type]
		tmp.Data = int(val.Amount)
		tmp.Color = utils.GetColor(key)
		if len(validObj) <= 5 {
			validObj = append(validObj, tmp)
		}
	}
	// 退费
	qb, _ = orm.NewQueryBuilder("mysql")
	qb = qb.Select("sum(amount) as amount", "contract.type type").From(models.ContractTBName() + " AS contract").LeftJoin(models.StudentTBName() + " AS student").On("student.id = contract.student_id").Where("contract.status = ?")
	if dateStart != "" {
		qb = qb.And("DATE_FORMAT(contract.created_at, '%Y-%m-%d') >= '" + dateStart + "'")
	}
	if dateEnd != "" {
		qb = qb.And("DATE_FORMAT(contract.created_at, '%Y-%m-%d') <= '" + dateEnd + "'")
	}
	if schoolId != "" {
		qb = qb.And("student.school_id = " + schoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("student.school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql = qb.GroupBy("contract.type").OrderBy("amount DESC").String()
	bm := make([]*Foo, 0)
	orm.NewOrm().Raw(sql, student.ContractStatusBack).QueryRows(&bm)
	backObj := make([]Flot, 0)
	for key, val := range bm {
		backCount += val.Amount
		tmp.Label = student.ContractTypeMap[val.Type]
		tmp.Data = int(val.Amount)
		tmp.Color = utils.GetColor(key)
		if len(backObj) <= 5 {
			backObj = append(backObj, tmp)
		}
	}
	ret := map[string]interface{}{"Count": count, "Valid": validObj, "BackCount": backCount, "Back": backObj}
	c.JsonResult(enums.JRCodeSucc, "succ", ret)
}
func (c *ContractController) New() {
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "contract/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "contract/new_footerjs.html"
}

// 合同统计数据
func (c *ContractController) NewData() {
	dateStart := c.GetString("DateStart")
	dateEnd := c.GetString("DateEnd")
	schoolId := c.GetString("SchoolId")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("DATE_FORMAT(contract.created_at,'%Y%u') weeks", "count(*) count").From(models.ContractTBName() + " AS contract").LeftJoin(models.StudentTBName() + " AS student").On("student.id = contract.student_id").Where("contract.status = ?")
	if dateStart != "" {
		qb = qb.And("DATE_FORMAT(contract.created_at, '%Y-%m-%d') >= '" + dateStart + "'")
	}
	if dateEnd != "" {
		qb = qb.And("DATE_FORMAT(contract.created_at, '%Y-%m-%d') <= '" + dateEnd + "'")
	}
	if schoolId != "" {
		qb = qb.And("student.school_id = " + schoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("student.school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql := qb.GroupBy("weeks").String()
	type NewTotal struct {
		Weeks string
		Count int
	}
	m := make([]NewTotal, 0)
	ret := [][]interface{}{}
	orm.NewOrm().Raw(sql, student.ContractStatusValid).QueryRows(&m)
	for _, val := range m {
		tmp := []interface{}{}
		tmp = append(tmp, val.Weeks)
		tmp = append(tmp, val.Count)
		ret = append(ret, tmp)
	}
	c.JsonResult(enums.JRCodeSucc, "succ", ret)
}

func (c *ContractController) Balance() {
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "contract/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "contract/balance_footerjs.html"
}

func (c *ContractController) BalanceGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params student.ContractQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("sum(surplus) surplus", "count(surplus_quantity) surplus_quantity", "contract.type").From(models.ContractTBName() + " AS contract").LeftJoin(models.StudentTBName() + " AS student").On("student.id = contract.student_id").Where("contract.status = ?")
	if params.DateStart != "" {
		qb = qb.And("DATE_FORMAT(contract.created_at, '%Y-%m-%d') >= '" + params.DateStart + "'")
	}
	if params.DateEnd != "" {
		qb = qb.And("DATE_FORMAT(contract.created_at, '%Y-%m-%d') <= '" + params.DateEnd + "'")
	}
	if params.SchoolId != "" {
		qb = qb.And("student.school_id = " + params.SchoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("student.school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql := qb.GroupBy("contract.type").String()
	m := make([]student.Contract, 0)
	orm.NewOrm().Raw(sql, student.ContractStatusValid).QueryRows(&m)
	total := student.Contract{Type: -1}
	for _, val := range m {
		total.Surplus += val.Surplus
		total.SurplusQuantity += val.SurplusQuantity
	}
	m = append(m, total)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = len(m)
	result["rows"] = m
	c.Data["json"] = result
	c.ServeJSON()
}

//打印
func (c *ContractController) Print() {
	id, _ := c.GetInt(":id")
	m := student.Contract{Id: id}
	orm.NewOrm().Read(&m)
	studentm := student.Student{Id: m.Student.Id}
	orm.NewOrm().Read(&studentm)
	c.Data["m"] = m
	c.Data["student"] = studentm
	c.Data["chin"] = utils.Num2Chinese(int(m.Amount))
	c.Data["school_id"] = c.CurUser.SchoolId
	if m.Type == student.ContractTypeYi {
		c.SetTpl("contract/print_yi.html", "shared/layout_pullbox.html")
	} else if m.Type == student.ContractTypeXiao {
		c.SetTpl("contract/print_xiao.html", "shared/layout_pullbox.html")
	} else if m.Type == student.ContractTypeTuo {
		c.SetTpl("contract/print_tuo.html", "shared/layout_pullbox.html")
	}
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "contract/print_footerjs.html"
}

package student

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"gocms/controllers"
	"gocms/enums"
	"gocms/models"
	"gocms/models/student"
	"reflect"
	"strconv"
	"strings"
)

type ContractPriceController struct {
	controllers.BaseController
}

func (c *ContractPriceController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.CheckAuthor()
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
}

func (c *ContractPriceController) Index() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "contractprice/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "contractprice/index_footerjs.html"
	type obj struct {
		student.School
		Myi   []student.ContractPrice
		Mxiao []student.ContractPrice
	}
	schools := make([]obj, 0)
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From(models.SchoolTBName()).String()
	orm.NewOrm().Raw(sql).QueryRows(&schools)
	//orm.NewOrm().QueryTable(models.SchoolTBName()).All(&schools)
	for key, val := range schools {
		yi := make([]student.ContractPrice, 0)
		xiao := make([]student.ContractPrice, 0)
		orm.NewOrm().QueryTable(models.ContractPriceTBName()).Filter("type", student.ContractTypeYi).Filter("school_id", val.Id).All(&yi)
		orm.NewOrm().QueryTable(models.ContractPriceTBName()).Filter("type", student.ContractTypeXiao).Filter("school_id", val.Id).All(&xiao)
		schools[key].Myi = yi
		schools[key].Mxiao = xiao

	}
	c.Data["schools"] = schools
}

func (c *ContractPriceController) Edit() {
	m := student.ContractPrice{}
	if err := c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	length := c.GetString("Length")
	arr := strings.Split(length, "-")
	if len(arr) < 2 {
		c.JsonResult(enums.JRCodeFailed, "时长范围格式有误", "")
	}
	m.Lower, _ = strconv.Atoi(arr[0])
	m.Upper, _ = strconv.Atoi(arr[1])
	if m.Id > 0 {
		_, err := orm.NewOrm().Update(&m)
		if err != nil {
			c.JsonResult(enums.JRCodeFailed, "失败", "")
		}
		c.JsonResult(enums.JRCodeSucc, "编辑成功", "")
	}
	//c.JsonResult(enums.JRCodeFailed,"不允许添加","")
	_, err := orm.NewOrm().Insert(&m)
	if err != nil {
		c.JsonResult(enums.JRCodeFailed, "失败", "")
	}
	c.JsonResult(enums.JRCodeSucc, "编辑成功", "")
}

// 查找
func (c *ContractPriceController) Search() {
	m := student.ContractPrice{}
	if err := c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	length, _ := c.GetFloat("Quantity")
	grade, _ := c.GetInt("Grade")
	valid := validation.Validation{}
	valid.Range(m.Type, 0, 2, "类型")
	valid.Range(grade, 1, 300, "年级")
	valid.Required(m.SchoolId, "校区ID")
	valid.Required(length, "时长")
	if valid.HasErrors() {
		c.JsonResult(enums.JRCode501, valid.Errors[0].Key+valid.Errors[0].Message, "")
	}
	m, err := m.Search(m.Type, m.Grade, length)
	if err != nil {
		c.JsonResult(enums.JRCodeFailed, "没有符合条件的结果", "")
	}
	priceCol := "Grade" + strconv.Itoa(grade)
	statusCol := "Status" + strconv.Itoa(grade)
	immutable := reflect.ValueOf(m)
	val := immutable.FieldByName(priceCol).Int()
	status := immutable.FieldByName(statusCol).Int()
	c.JsonResult(enums.JRCodeSucc, "succ", map[string]interface{}{"val": val, "is_lock": status})
}

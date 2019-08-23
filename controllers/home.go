package controllers

import (
	"github.com/astaxie/beego/orm"
	"gocms/models/student"
	"gocms/models/teacher"
	"strings"
	"time"

	"gocms/enums"
	"gocms/models"
	"gocms/utils"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	//判断是否登录
	c.checkLogin()
	/*克隆上个月的支出*/
	first := utils.GetFirstDateOfMonth(time.Now())
	firstDate := first.Format("2006-01-02")
	m := &teacher.Expend{}
	err := orm.NewOrm().QueryTable(models.ExpendTBName()).Filter("month", firstDate).One(m)
	if err != nil { // 没有结果
		pre := []*teacher.Expend{}
		preDate := utils.GetPreFirstDateOfMonth(time.Now())
		preFirstDate := preDate.Format("2006-01-02")
		orm.NewOrm().QueryTable(models.ExpendTBName()).Filter("month", preFirstDate).All(&pre)
		for _, val := range pre {
			val.Month = firstDate
			val.Id = 0
			orm.NewOrm().Insert(val)
		}
	}
	/*克隆上个月的支出*/
	/*合同单价默认添加*/
	schools := make([]student.School, 0)
	orm.NewOrm().QueryTable(models.SchoolTBName()).All(&schools)
	for _, val := range schools {
		num, _ := orm.NewOrm().QueryTable(models.ContractPriceTBName()).Filter("school_id", val.Id).Count()
		if num == 6 {
			continue
		} else if num < 6 && num > 0 {
			orm.NewOrm().QueryTable(models.ContractPriceTBName()).Filter("school_id", val.Id).Delete()
		}
		m1 := student.ContractPrice{Type: student.ContractTypeYi, SchoolId: val.Id}
		m2 := student.ContractPrice{Type: student.ContractTypeXiao, SchoolId: val.Id}
		m1mul := make([]student.ContractPrice, 0)
		m1mul = append(m1mul, m1)
		m1mul = append(m1mul, m1)
		m1mul = append(m1mul, m1)
		m2mul := make([]student.ContractPrice, 0)
		m2mul = append(m2mul, m2)
		m2mul = append(m2mul, m2)
		m2mul = append(m2mul, m2)
		orm.NewOrm().InsertMulti(3, m1mul)
		orm.NewOrm().InsertMulti(3, m2mul)
	}
	/*合同单价默认添加*/
	c.SetTpl()
}
func (c *HomeController) Page404() {
	c.SetTpl()
}
func (c *HomeController) Error() {
	c.Data["error"] = c.GetString(":error")
	c.SetTpl("home/error.html", "shared/layout_pullbox.html")
}
func (c *HomeController) Login() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "home/login_headcssjs.html"
	c.LayoutSections["footerjs"] = "home/login_footerjs.html"
	c.SetTpl("home/login.html", "shared/layout_base.html")
}
func (c *HomeController) DoLogin() {
	username := strings.TrimSpace(c.GetString("UserName"))
	userpwd := strings.TrimSpace(c.GetString("UserPwd"))
	if len(username) == 0 || len(userpwd) == 0 {
		c.JsonResult(enums.JRCodeFailed, "用户名和密码不正确", "")
	}
	userpwd = utils.String2md5(userpwd)
	user, err := models.BackendUserOneByUserName(username, userpwd)
	//fmt.Println(username,userpwd,user,err)
	if user != nil && err == nil {
		if user.Status == enums.Disabled {
			c.JsonResult(enums.JRCodeFailed, "用户被禁用，请联系管理员", "")
		}
		//保存用户信息到session
		c.setBackendUser2Session(user.Id)
		//获取用户信息
		c.JsonResult(enums.JRCodeSucc, "登录成功", "")
	} else {
		c.JsonResult(enums.JRCodeFailed, "用户名或者密码错误", "")
	}
}
func (c *HomeController) Logout() {
	user := models.BackendUser{}
	c.SetSession("backenduser", user)
	c.PageLogin()
}

func (c *HomeController) CkeditorUpload() {
	type ck struct {
		Uploaded int               `json:"uploaded"`
		FileName string            `json:"file_name"`
		Url      string            `json:"url"`
		Error    map[string]string `json:"error"`
	}
	var resp ck
	f, h, err := c.GetFile("upload")
	if err != nil {
		resp.Error = map[string]string{"message": err.Error()}
		c.Data["json"] = resp
		c.ServeJSON()
		c.StopRun()
	}
	defer f.Close()
	// 保存位置在 static/upload, 没有文件夹要先创建
	file, err := utils.UploadImage(h, "news")
	if err != nil {
		resp.Error = map[string]string{"message": err.Error()}
	} else {
		resp.Uploaded = 1
		resp.Url = "/static/upload/" + file
		resp.FileName = file
	}
	c.Data["json"] = resp
	c.ServeJSON()
	c.StopRun()

}

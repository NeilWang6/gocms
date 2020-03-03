package fronted

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/cuua/gocms/models"
	"github.com/cuua/gocms/models/fronted"
)

type StaffController struct {
	BaseController
}

func (c *StaffController) Prepare() {
	c.BaseController.Prepare()
}

func (c *StaffController) List() {
	staffs := make([]fronted.Staff, 0)
	orm.NewOrm().QueryTable(models.StaffTBName()).OrderBy("school", "subject", "id").All(&staffs)
	type NavId struct {
		Id          int
		SubjectName string
	}
	nav := make(map[string][]NavId)
	for _, val := range staffs {
		m := NavId{Id: val.Id, SubjectName: val.SubjectName}
		nav[val.SchoolName] = append(nav[val.SchoolName], m)
	}
	c.Data["nav"] = nav
	c.Data["staffs"] = staffs
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "main/headcss_page.html"
}

func (c *StaffController) Mlist() {
	staffs := make([]fronted.Staff, 0)
	orm.NewOrm().QueryTable(models.StaffTBName()).OrderBy("school", "subject", "id").All(&staffs)
	type NavId struct {
		Id          int
		SubjectName string
	}
	nav := make(map[string][]NavId)
	for _, val := range staffs {
		m := NavId{Id: val.Id, SubjectName: val.SubjectName}
		nav[val.SchoolName] = append(nav[val.SchoolName], m)
	}
	fmt.Printf("%+v", staffs)
	c.Data["nav"] = nav
	c.Data["staffs"] = staffs
	c.SetMTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "main/m/headcss_page.html"
}

package fronted

import (
	"github.com/astaxie/beego/orm"
	"github.com/cuua/gocms/models"
	"github.com/cuua/gocms/models/fronted"
)

type RecruitController struct {
	BaseController
}

func (c *RecruitController) Prepare() {
	c.BaseController.Prepare()
}

func (c *RecruitController) List() {
	recruit := make([]fronted.Recruit, 0)
	orm.NewOrm().QueryTable(models.RecruitTBName()).OrderBy("-id").All(&recruit)
	c.Data["recruit"] = recruit
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "main/headcss_page.html"
}

func (c *RecruitController) Mlist() {
	recruit := make([]fronted.Recruit, 0)
	orm.NewOrm().QueryTable(models.RecruitTBName()).OrderBy("-id").All(&recruit)
	c.Data["recruit"] = recruit
	c.SetMTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "main/m/headcss_page.html"
}

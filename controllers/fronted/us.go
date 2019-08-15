package fronted

import (
	"github.com/astaxie/beego/orm"
	"gocms/models"
	"gocms/models/fronted"
)

type UsController struct {
	BaseController
}

func (c *UsController) Prepare() {
	c.BaseController.Prepare()
}

func (c *UsController) List() {
	us := make([]fronted.Us, 0)
	orm.NewOrm().QueryTable(models.UsTBName()).OrderBy("-id").Limit(1).All(&us)
	history := make([]fronted.History, 0)
	orm.NewOrm().QueryTable(models.HistoryTBName()).OrderBy("-date").All(&history)
	c.Data["us"] = us
	c.Data["history"] = history
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	//c.LayoutSections["headcssjs"] = "main/headcss_page.html"
	c.LayoutSections["headcssjs"] = "us/headcss_page.html"
}

func (c *UsController) Mlist() {
	us := fronted.Us{}
	orm.NewOrm().QueryTable(models.UsTBName()).OrderBy("-id").Limit(1).One(&us)
	history := make([]fronted.History, 0)
	orm.NewOrm().QueryTable(models.HistoryTBName()).OrderBy("-date").All(&history)
	c.Data["us"] = us
	c.Data["history"] = history
	c.SetMTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "us/m/headcss_page.html"
}

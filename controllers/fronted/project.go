package fronted

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/cuua/gocms/models"
	"github.com/cuua/gocms/models/fronted"
)

type ProjectController struct {
	BaseController
}

func (c *ProjectController) Prepare() {
	c.BaseController.Prepare()
}

type NavId struct {
	Id         int
	SchoolName string
	Sort       int
}

func (c *ProjectController) List() {
	projects := make([]fronted.Project, 0)
	orm.NewOrm().QueryTable(models.ProjectTBName()).OrderBy("-sort").All(&projects)
	nav := make(map[int][]NavId)
	for _, val := range projects {
		m := NavId{Id: val.Id, SchoolName: val.SchoolName, Sort: val.Sort}
		nav[val.Sort] = append(nav[val.Sort], m)
	}
	c.Data["nav"] = nav
	c.Data["projects"] = projects
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "main/headcss_page.html"
}

func (c *ProjectController) Mlist() {
	projects := make([]fronted.Project, 0)
	orm.NewOrm().QueryTable(models.ProjectTBName()).OrderBy("city", "-sort", "id").All(&projects)
	type NavId struct {
		Id         int
		SchoolName string
	}
	nav := make(map[string][]NavId)
	for _, val := range projects {
		m := NavId{Id: val.Id, SchoolName: val.SchoolName}
		nav[val.City] = append(nav[val.City], m)
	}
	fmt.Printf("%+v", nav)
	c.Data["nav"] = nav
	c.Data["projects"] = projects
	c.SetMTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "main/m/headcss_page.html"
}

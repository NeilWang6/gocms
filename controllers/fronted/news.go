package fronted

import (
	"github.com/astaxie/beego/orm"
	"github.com/cuua/gocms/models"
	"github.com/cuua/gocms/models/fronted"
	"github.com/cuua/gocms/utils"
)

type NewsController struct {
	BaseController
}

func (c *NewsController) Prepare() {
	c.BaseController.Prepare()
}

// 新闻列表
func (c *NewsController) List() {
	news := make([]fronted.News, 0)
	orm.NewOrm().QueryTable(models.NewsTBName()).OrderBy("-sort", "-id").Limit(6).All(&news)
	for key, val := range news {
		body := utils.TrimHtml(val.Content)
		bodyRune := []rune(body)
		if len(bodyRune) > 120 {
			news[key].Content = string(bodyRune[:120])
		} else {
			news[key].Content = utils.TrimHtml(val.Content)
		}
		news[key].Content += "..."
	}
	c.Data["news"] = news
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "main/headcss_page.html"
}

func (c *NewsController) Detail() {
	id, _ := c.GetInt(":id")
	m := fronted.News{Id: id}
	orm.NewOrm().Read(&m)
	c.Data["news"] = m
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "main/headcss_page.html"
}

func (c *NewsController) Mlist() {
	news := make([]fronted.News, 0)
	orm.NewOrm().QueryTable(models.NewsTBName()).OrderBy("-sort", "-id").Limit(6).All(&news)
	for key, val := range news {
		body := utils.TrimHtml(val.Content)
		bodyRune := []rune(body)
		if len(bodyRune) > 120 {
			news[key].Content = string(bodyRune[:120])
		} else {
			news[key].Content = utils.TrimHtml(val.Content)
		}
		news[key].Content += "..."
	}
	c.Data["news"] = news
	c.SetMTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "main/m/headcss_page.html"
}

func (c *NewsController) Mdetail() {
	id, _ := c.GetInt(":id")
	m := fronted.News{Id: id}
	orm.NewOrm().Read(&m)
	c.Data["news"] = m
	c.SetMTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "main/m/headcss_page.html"
}

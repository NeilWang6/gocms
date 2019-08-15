package fronted

import (
	"github.com/astaxie/beego/orm"
	"gocms/models"
	"gocms/models/fronted"
	"gocms/utils"
)

type MainController struct {
	BaseController
}

func (c *MainController) Prepare() {
	c.BaseController.Prepare()
}

func (c *MainController) Index() {
	obj := make([]fronted.Banner, 0)
	news := make([]fronted.News, 0)
	orm.NewOrm().QueryTable(models.BannerTBName()).OrderBy("-sort").All(&obj)
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
	c.Data["obj"] = obj
	c.Data["news"] = news
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "main/headcss_page.html"
	//c.LayoutSections["footerjs"] = "main/index_footerjs.html"
}

// 移动端首页
func (c *MainController) Mindex() {
	obj := make([]fronted.Banner, 0)
	news := make([]fronted.News, 0)
	orm.NewOrm().QueryTable(models.BannerTBName()).OrderBy("-sort").All(&obj)
	orm.NewOrm().QueryTable(models.NewsTBName()).OrderBy("-sort", "-id").Limit(6).All(&news)
	for key, val := range news {
		body := utils.TrimHtml(val.Content)
		bodyRune := []rune(body)
		if len(bodyRune) > 50 {
			news[key].Content = string(bodyRune[:50])
		} else {
			news[key].Content = utils.TrimHtml(val.Content)
		}
		news[key].Content += "..."
	}
	c.Data["obj"] = obj
	c.Data["news"] = news
	c.SetMTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "main/m/headcss_page.html"
}

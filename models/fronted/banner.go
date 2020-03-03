package fronted

import (
	"github.com/astaxie/beego/orm"
	"github.com/cuua/gocms/models"
)

func init() {
	orm.RegisterModel(new(Banner))
}

type Banner struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Url      string `json:"url"`     // 跳转地址
	ImageUrl string `json:"img_url"` // 图片地址
	Sort     int    `json:"sort"`
	State    int    `json:"state"`
}

func (*Banner) TableName() string {
	return models.BannerTBName()
}

// QueryParam 用于搜索的类
type BannerQueryParam struct {
	models.BaseQueryParam
	NameLike string
}

// RoleDataList 获取学校列表
func BannerDataList(params *BannerQueryParam) []*Banner {
	params.Limit = -1
	params.Sort = "Seq"
	params.Order = "asc"
	data, _ := BannerPageList(params)
	return data
}

// RolePageList 获取分页数据
func BannerPageList(params *BannerQueryParam) ([]*Banner, int64) {
	query := orm.NewOrm().QueryTable(models.BannerTBName())
	data := make([]*Banner, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("name__istartswith", params.NameLike)
	if params.CurUser.IsSuper != true {
		query = query.Filter("id", params.CurUser.SchoolId)
	}
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

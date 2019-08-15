package student

import (
	"github.com/astaxie/beego/orm"
	"gocms/models"
)

func init() {
	orm.RegisterModel(new(School))
}

type School struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// RoleQueryParam 用于搜索的类
type SchoolQueryParam struct {
	models.BaseQueryParam
	NameLike string
}

func (m *School) TableName() string {
	return models.SchoolTBName()
}

// RoleDataList 获取学校列表
func SchoolDataList(params *SchoolQueryParam) []*School {
	params.Limit = -1
	params.Sort = "Seq"
	params.Order = "asc"
	data, _ := SchoolPageList(params)
	return data
}

// RolePageList 获取分页数据
func SchoolPageList(params *SchoolQueryParam) ([]*School, int64) {
	query := orm.NewOrm().QueryTable(models.SchoolTBName())
	data := make([]*School, 0)
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

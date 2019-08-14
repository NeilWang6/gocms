package student

import (
	"github.com/astaxie/beego/orm"
	"sdrms/models"
)

func init() {
	orm.RegisterModel(new(StudentArea))
}

type StudentArea struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (m *StudentArea) TableName() string {
	return models.StudentAreaTBName()
}

// RoleQueryParam 用于搜索的类
type StudentAreaQueryParam struct {
	models.BaseQueryParam
	NameLike string
}

// RoleDataList 获取角色列表
func StudentAreaDataList(params *StudentAreaQueryParam) []*StudentArea {
	params.Limit = -1
	params.Sort = "Seq"
	params.Order = "asc"
	data, _ := StudentAreaPageList(params)
	return data
}

// RolePageList 获取分页数据
func StudentAreaPageList(params *StudentAreaQueryParam) ([]*StudentArea, int64) {
	query := orm.NewOrm().QueryTable(models.StudentAreaTBName())
	data := make([]*StudentArea, 0)
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
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

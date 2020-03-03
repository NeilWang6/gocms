package teacher

import (
	"github.com/astaxie/beego/orm"
	"github.com/cuua/gocms/models"
)

func init() {
	orm.RegisterModel(new(Subject))
}

type Subject struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (m *Subject) TableName() string {
	return models.SubjectTBName()
}

// RoleQueryParam 用于搜索的类
type SubjectQueryParam struct {
	models.BaseQueryParam
	NameLike string
}

// RoleDataList 获取学校列表
func SubjectDataList(params *SubjectQueryParam) []*Subject {
	params.Limit = -1
	params.Sort = "Seq"
	params.Order = "asc"
	data, _ := SubjectPageList(params)
	return data
}

// RolePageList 获取分页数据
func SubjectPageList(params *SubjectQueryParam) ([]*Subject, int64) {
	query := orm.NewOrm().QueryTable(models.SubjectTBName())
	data := make([]*Subject, 0)
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

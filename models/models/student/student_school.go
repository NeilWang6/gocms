package student

import (
	"github.com/astaxie/beego/orm"
	"sdrms/models"
)

func init() {
	orm.RegisterModel(new(StudentSchool))
}

type StudentSchool struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (m *StudentSchool) TableName() string {
	return models.StudentSchoolTBName()
}

// RoleQueryParam 用于搜索的类
type StudentSchoolQueryParam struct {
	models.BaseQueryParam
	NameLike string
}

// RoleDataList 获取角色列表
func StudentSchoolDataList(params *StudentSchoolQueryParam) []*StudentSchool {
	params.Limit = -1
	params.Sort = "Seq"
	params.Order = "asc"
	data, _ := StudentSchoolPageList(params)
	return data
}

// RolePageList 获取分页数据
func StudentSchoolPageList(params *StudentSchoolQueryParam) ([]*StudentSchool, int64) {
	query := orm.NewOrm().QueryTable(models.StudentSchoolTBName())
	data := make([]*StudentSchool, 0)
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

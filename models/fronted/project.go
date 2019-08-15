package fronted

import (
	"github.com/astaxie/beego/orm"
	"gocms/models"
	"time"
)

func init() {
	orm.RegisterModel(new(Project))
}

type Project struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	City       string    `json:"city"`
	SchoolId   int       `json:"school_id"`
	SchoolName string    `json:"school_name"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	Sort       int       `json:"sort"`
	ImageUrl   string    `json:"image_url"`
}

func (m *Project) TableName() string {
	return models.ProjectTBName()
}

// QueryParam 用于搜索的类
type ProjectQueryParam struct {
	models.BaseQueryParam
	NameLike string
}

// RoleDataList 获取学校列表
func ProjectDataList(params *ProjectQueryParam) []*Project {
	params.Limit = -1
	params.Sort = "Seq"
	params.Order = "asc"
	data, _ := ProjectPageList(params)
	return data
}

// RolePageList 获取分页数据
func ProjectPageList(params *ProjectQueryParam) ([]*Project, int64) {
	query := orm.NewOrm().QueryTable(models.ProjectTBName())
	data := make([]*Project, 0)
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

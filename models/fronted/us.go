package fronted

import (
	"github.com/astaxie/beego/orm"
	"gocms/models"
	"time"
)

func init() {
	orm.RegisterModel(new(Us))
}

type Us struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Sort      int       `json:"sort"`
}

func (*Us) TableName() string {
	return models.UsTBName()
}

// QueryParam 用于搜索的类
type UsQueryParam struct {
	models.BaseQueryParam
	NameLike string
}

// RoleDataList 获取学校列表
func UsDataList(params *UsQueryParam) []*Us {
	params.Limit = -1
	params.Sort = "Seq"
	params.Order = "asc"
	data, _ := UsPageList(params)
	return data
}

// RolePageList 获取分页数据
func UsPageList(params *UsQueryParam) ([]*Us, int64) {
	query := orm.NewOrm().QueryTable(models.UsTBName())
	data := make([]*Us, 0)
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

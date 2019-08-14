package fronted

import (
	"github.com/astaxie/beego/orm"
	"sdrms/models"
	"time"
)

func init() {
	orm.RegisterModel(new(Recruit))
}

type Recruit struct {
	Id        int       `json:"id"`
	ImageUrl  string    `json:"image_url"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Sort      int       `json:"sort"`
}

func (*Recruit) TableName() string {
	return models.RecruitTBName()
}

// QueryParam 用于搜索的类
type RecruitQueryParam struct {
	models.BaseQueryParam
	NameLike string
}

// RoleDataList 获取学校列表
func RecruitDataList(params *RecruitQueryParam) []*Recruit {
	params.Limit = -1
	params.Sort = "Seq"
	params.Order = "asc"
	data, _ := RecruitPageList(params)
	return data
}

// RolePageList 获取分页数据
func RecruitPageList(params *RecruitQueryParam) ([]*Recruit, int64) {
	query := orm.NewOrm().QueryTable(models.RecruitTBName())
	data := make([]*Recruit, 0)
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

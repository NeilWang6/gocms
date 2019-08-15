package fronted

import (
	"github.com/astaxie/beego/orm"
	"gocms/models"
	"time"
)

func init() {
	orm.RegisterModel(new(History))
}

type History struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Sort      int       `json:"sort"`
	Date      time.Time `json:"date"`
}

func (*History) TableName() string {
	return models.HistoryTBName()
}

// QueryParam 用于搜索的类
type HistoryQueryParam struct {
	models.BaseQueryParam
	NameLike string
}

// RoleDataList 获取学校列表
func HistoryDataList(params *HistoryQueryParam) []*History {
	params.Limit = -1
	params.Sort = "Seq"
	params.Order = "asc"
	data, _ := HistoryPageList(params)
	return data
}

// RolePageList 获取分页数据
func HistoryPageList(params *HistoryQueryParam) ([]*History, int64) {
	query := orm.NewOrm().QueryTable(models.HistoryTBName())
	data := make([]*History, 0)
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

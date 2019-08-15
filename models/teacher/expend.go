package teacher

import (
	"github.com/astaxie/beego/orm"
	"gocms/models"
	"gocms/utils"
	"time"
)

func init() {
	orm.RegisterModel(new(Expend))
}

type Expend struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	Month    string `json:"month"`
	SchoolId int    `json:"school_id"`
}

func (m *Expend) TableName() string {
	return models.ExpendTBName()
}

// RoleQueryParam 用于搜索的类
type ExpendQueryParam struct {
	models.BaseQueryParam
	SchoolId  string
	NameLike  string
	DateStart string
	DateEnd   string
}

// RoleDataList 获取学校列表
func ExpendDataList(params *ExpendQueryParam) []*Expend {
	params.Limit = -1
	params.Sort = "Seq"
	params.Order = "asc"
	data, _ := ExpendPageList(params)
	return data
}

// RolePageList 获取分页数据
func ExpendPageList(params *ExpendQueryParam) ([]*Expend, int64) {
	query := orm.NewOrm().QueryTable(models.ExpendTBName())
	data := make([]*Expend, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	var startDate time.Time
	if params.DateStart == "" {
		startDate = utils.GetFirstDateOfMonth(time.Now())
	} else {
		startDate, _ = time.Parse(utils.FormatDateTime, params.DateStart+" 00:00:00")
		startDate = utils.GetFirstDateOfMonth(startDate)
	}
	if params.SchoolId != "" {
		query = query.Filter("school_id", params.SchoolId)
	}
	date := startDate.Format("2006-01-02")
	query = query.Filter("month", date)
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

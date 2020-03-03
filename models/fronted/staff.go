package fronted

import (
	"github.com/astaxie/beego/orm"
	"github.com/cuua/gocms/models"
	"time"
)

func init() {
	orm.RegisterModel(new(Staff))
}

type Staff struct {
	Id          int       `json:"id"`
	ImageUrl    string    `json:"image_url"`
	Name        string    `json:"name"`
	Subject     int       `json:"subject"`
	SubjectName string    `json:"subject_name"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	Sort        int       `json:"sort"`
	School      string    `json:"school"`
	SchoolName  string    `json:"school_name""`
}

func (*Staff) TableName() string {
	return models.StaffTBName()
}

// QueryParam 用于搜索的类
type StaffQueryParam struct {
	models.BaseQueryParam
	NameLike string
}

// RoleDataList 获取学校列表
func StaffDataList(params *StaffQueryParam) []*Staff {
	params.Limit = -1
	params.Sort = "Seq"
	params.Order = "asc"
	data, _ := StaffPageList(params)
	return data
}

// RolePageList 获取分页数据
func StaffPageList(params *StaffQueryParam) ([]*Staff, int64) {
	query := orm.NewOrm().QueryTable(models.StaffTBName())
	data := make([]*Staff, 0)
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

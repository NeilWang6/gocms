package teacher

import (
	"github.com/astaxie/beego/orm"
	"gocms/models"
	"gocms/models/student"
	"reflect"
	"strconv"
	"time"
)

const TeacherStatusValid = 0
const TeacherStatusDelete = -1

func init() {
	orm.RegisterModel(new(Teacher))
}

type Teacher struct {
	Id                     int             `json:"id"`
	Name                   string          `json:"name"`
	Idcard                 string          `json:"idcard"`
	Cardno                 string          `json:"cardno"`
	Phone                  string          `json:"phone"`
	School                 *student.School `orm:"rel(one)"`
	Department             string          `json:"department"`
	Contacter              string          `json:"contacter"`
	ContacterPhone         string          `json:"contacter_phone"`
	Sex                    string          `json:"sex"`
	Picture                string          `json:"picture"`
	Birth                  string          `json:"birth"`
	EntryTime              string          `json:"entry_time"`
	ContractExpirationTime string          `json:"contract_expiration_time"`
	Reward                 string          `json:"reward"`
	Address                string          `json:"address"`
	CreatedAt              time.Time       `json:"created_at"`
	Note                   time.Time       `json:"note"`
	UpdatedAt              string          `json:"updated_at"`
	Status                 int             `json:"status"`
	Price1                 int             `json:"price1"`
	Price2                 int             `json:"price2"`
	Price3                 int             `json:"price3"`
	Price10                int             `json:"price10"`
	Price20                int             `json:"price20"`
	Price30                int             `json:"price30"`
	Price40                int             `json:"price40"`
	Price50                int             `json:"price50"`
	Price60                int             `json:"price60"`
	Price70                int             `json:"price70"`
	Price80                int             `json:"price80"`
	Price90                int             `json:"price90"`
	Price100               int             `json:"price100"`
	Price110               int             `json:"price110"`
	Price120               int             `json:"price120"`
	Xprice1                int             `json:"xprice1"`
	Xprice2                int             `json:"xprice2"`
	Xprice3                int             `json:"xprice3"`
	Xprice10               int             `json:"xprice10"`
	Xprice20               int             `json:"xprice20"`
	Xprice30               int             `json:"xprice30"`
	Xprice40               int             `json:"xprice40"`
	Xprice50               int             `json:"xprice50"`
	Xprice60               int             `json:"xprice60"`
	Xprice70               int             `json:"xprice70"`
	Xprice80               int             `json:"xprice80"`
	Xprice90               int             `json:"xprice90"`
	Xprice100              int             `json:"xprice100"`
	Xprice110              int             `json:"xprice110"`
	Xprice120              int             `json:"xprice120"`
	SocialSecurity         string          `json:"social_security"`
	Funds                  string          `json:"funds"`
	SalaryCard             string          `json:"salary_card"`
	Balance                float64         `json:"balance"`
}

func (m *Teacher) TableName() string {
	return models.TeacherTBName()
}

type TeacherQueryParam struct {
	models.BaseQueryParam
	NameLike  string
	SchoolId  string
	Status    string
	DateStart string
	DateEnd   string
}

// PageList 获取分页数据
func TeacherPageList(params *TeacherQueryParam) ([]*Teacher, int64) {
	query := orm.NewOrm().QueryTable(models.TeacherTBName()).RelatedSel()
	data := make([]*Teacher, 0)
	//默认排序
	sortorder := "Id"
	if params.Sort != "" {
		sortorder = params.Sort
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	if params.NameLike != "" {
		query = query.Filter("name__contains", params.NameLike)
	}
	if params.SchoolId != "" {
		query = query.Filter("school_id", params.SchoolId)
	}
	if params.Status != "" {
		query = query.Filter("status", params.Status)
	}
	if params.DateStart != "" {
		query = query.Filter("created_at__gte", params.DateStart)
	}
	if params.DateEnd != "" {
		query = query.Filter("created_at__lte", params.DateEnd)
	}
	if params.CurUser.IsSuper != true {
		query = query.Filter("school_id", params.CurUser.SchoolId)
	}
	if params.CurUser.RealName == "教师组" {
		teacher := Teacher{}
		orm.NewOrm().QueryTable(models.TeacherTBName()).Filter("cardno", params.CurUser.UserName).One(&teacher, "id")
		query = query.Filter("id", teacher.Id)
	}
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// RoleDataList 获取学生列表
func TeacherDataList(params *TeacherQueryParam) []*Teacher {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := TeacherPageList(params)
	return data
}

// TeacherDelete 批量删除
func TeacherDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(models.TeacherTBName())
	num, err := query.Filter("id__in", ids).Update(orm.Params{
		"status": TeacherStatusDelete,
	})
	return num, err
}

// 查询教师单价
func TeacherPrice(tid, grade, ctype int) (price int, err error) {
	m := Teacher{Id: tid}
	err = orm.NewOrm().Read(&m)
	if err != nil {
		return 0, err
	}
	var priceCol string
	if ctype == student.StudentType {
		priceCol = "Price" + strconv.Itoa(grade)
	} else {
		priceCol = "Xprice" + strconv.Itoa(grade)
	}
	immutable := reflect.ValueOf(m)
	val := immutable.FieldByName(priceCol).Int()
	return int(val), nil
}

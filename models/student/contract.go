package student

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"gocms/models"
	"strconv"
	"time"
)

const ContractTypeYi = 0
const ContractTypeXiao = 1
const ContractTypeTuo = 2
const ContractStatusValid = 0
const ContractStatusDelete = -1

const ContractStatusBack = 2

var ContractTypeMap = map[int]string{0: "一对一", 1: "小班", 2: "托班"}

func init() {
	orm.RegisterModel(new(Contract))
}

type Contract struct {
	Id              int       `json:"id"`
	Student         *Student  `orm:"rel(one)"`
	Name            string    `json:"name"`
	Type            int       `json:"type"`
	Price           int       `json:"price"`
	Quantity        float64   `json:"quantity"`
	SurplusQuantity float64   `json:"surplus_quantity"`
	Amount          float64   `json:"amount"`
	Surplus         float64   `json:"surplus"`
	Payment         string    `json:"payment"`
	CreatedAt       time.Time `json:"created_at"`
	UpdateAt        string    `json:"update_at"`
	Status          int       `json:"status"`
	EndAt           string
}

func (m *Contract) TableName() string {
	return models.ContractTBName()
}

// RoleQueryParam 用于搜索的类
type ContractQueryParam struct {
	models.BaseQueryParam
	SchoolId        string
	NameLike        string
	DateStart       string
	DateEnd         string
	Status          string
	DateStartRefund string
	DateEndRefund   string
}

// RoleDataList 获取学校列表
func ContractDataList(params *ContractQueryParam) []*Contract {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := ContractPageList(params)
	return data
}

// RolePageList 获取分页数据
func ContractPageList(params *ContractQueryParam) ([]*Contract, int64) {
	query := orm.NewOrm().QueryTable(models.ContractTBName()).RelatedSel("student")
	data := make([]*Contract, 0)
	//默认排序
	sortorder := "Id"
	if params.Sort != "" {
		sortorder = params.Sort
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	if params.NameLike != "" {
		query = query.Filter("Student__name__contains", params.NameLike)
	}
	if params.DateStart != "" {
		query = query.Filter("created_at__gte", params.DateStart)
	}
	if params.DateEnd != "" {
		query = query.Filter("created_at__lte", params.DateEnd)
	}
	if params.DateStartRefund != "" {
		query = query.Filter("updated_at__gte", params.DateStartRefund)
	}
	if params.DateEndRefund != "" {
		query = query.Filter("updated_at__lte", params.DateEndRefund)
	}
	if params.Status != "" {
		query = query.Filter("status", params.Status)
	}
	if params.CurUser.IsSuper != true {
		query = query.Filter("Student__school_id", params.CurUser.SchoolId)
	}
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

func ChooseContract(stuId, ctype int, length float64) (foo []Contract, err error) {
	contracts := make([]Contract, 0)
	_, err = orm.NewOrm().QueryTable(models.ContractTBName()).Filter("student_id", stuId).Filter("type", strconv.Itoa(ctype)).Filter("status", ContractStatusValid).OrderBy("id").All(&contracts)
	if err == orm.ErrNoRows {
		return nil, errors.New("没有符合条件的合同。")
	}
	var totalLen float64
	for _, val := range contracts {
		if val.SurplusQuantity > 0 {
			if val.SurplusQuantity >= length && len(foo) < 1 {
				foo = append(foo, val)
				break
			}
			foo = append(foo, val)
			totalLen += val.SurplusQuantity
			if totalLen >= length {
				break
			}
		}
	}
	if len(foo) == 0 {
		err = errors.New("没有符合条件的合同")
	}
	return
}

func GetContractTypeName(ctype int) string {
	if _, ok := ContractTypeMap[ctype]; !ok {
		return ""
	}
	return ContractTypeMap[ctype]
}

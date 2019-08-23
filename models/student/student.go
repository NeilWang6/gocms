package student

import (
	"github.com/astaxie/beego/orm"
	"gocms/models"
	"strconv"
	"strings"
	"time"
)

func init() {
	orm.RegisterModel(new(Student))
}

const StudentType = 0
const TeamType = 1
const StudentStatusInit = 0
const StudentStatusValid = 1
const StudentStatusDelete = -1
const StudentStatusStop = 2

var StudentTypeMap = []string{"未报名", "正常", "禁用"}
var StudentGradeMap = map[int]string{1: "幼小", 2: "幼中", 3: "幼大", 10: "一年级", 20: "二年级", 30: "三年级", 40: "四年级", 50: "五年级", 60: "六年级", 70: "初一", 80: "初二", 90: "初三", 100: "高一", 110: "高二", 120: "高三"}

type Student struct {
	Id             int            `json:"id"`
	Name           string         `json:"name"`
	Code           string         `json:"code"`
	Type           int            `json:"type"`
	Sex            string         `json:"sex"`
	Guarder        string         `json:"guarder"`
	Relate         string         `json:"relate"`
	Grade          int            `json:"grade"`
	Contact1       string         `json:"contact1"`
	Contact2       string         `json:"contact2"`
	School         *School        `orm:"rel(one)"`
	StudentSchool  *StudentSchool `orm:"rel(one)"`
	Adress         string         `json:"adress"`
	Note           string         `json:"note"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	Status         int            `json:"status"`
	GroupId        string         `json:"group_id"`
	GroupName      string         `json:"group_name"`
	Balance1       float64        `json:"balance1"`
	Balance1Length float64        `json:"balance1_length"`
	Balance2       float64        `json:"balance2"`
	Balance2Length float64        `json:"balance2_length"`
	Balance3       float64        `json:"balance3"`
	Balance3Length float64        `json:"balance3_length"`
}

func (m *Student) TableName() string {
	return models.StudentTBName()
}

// RoleQueryParam 用于搜索的类
type StudentQueryParam struct {
	models.BaseQueryParam
	Type            int
	NameLike        string
	SchoolId        string
	StudentSchoolId string
	Adress          string
	Grade           string
	Status          string
	DateStart       string
	DateEnd         string
	Balance2        string
	BalanceRate     string
	ContractType    string
}

// StudentPageList 获取分页数据
func StudentPageList(params *StudentQueryParam) ([]*Student, int64) {
	query := orm.NewOrm().QueryTable(models.StudentTBName())
	if params.Type == StudentType { // 是学生的话需关联查询
		query = query.RelatedSel()
	}
	data := make([]*Student, 0)
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
	if params.StudentSchoolId != "" {
		query = query.Filter("student_school_id", params.StudentSchoolId)
	}
	if params.Adress != "" {
		query = query.Filter("adress", params.Adress)
	}
	if params.Grade != "" {
		query = query.Filter("grade", params.Grade)
	}
	if params.Type > 0 && params.Type != 100 {
		query = query.Filter("type", params.Type)
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
	if params.Balance2 != "" {
		query = query.Filter("balance2__gt", params.Balance2)
	}
	if params.BalanceRate != "" && params.ContractType != "" {
		lengthInt, _ := strconv.Atoi(params.ContractType)
		lengthInt++
		lengthStr := strconv.Itoa(lengthInt)
		rate := strings.Split(params.BalanceRate, "-")
		query = query.Filter("balance"+lengthStr+"_length__gte", rate[0])
		query = query.Filter("balance"+lengthStr+"_length__lte", rate[1])
	}
	if params.CurUser.IsSuper != true {
		query = query.Filter("school_id", params.CurUser.SchoolId)
	}
	// 权限判断
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// Delete 批量删除
func StudentDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(models.StudentTBName())
	num, err := query.Filter("id__in", ids).Update(orm.Params{
		"status": StudentStatusDelete,
	})
	return num, err
}

// RoleDataList 获取学生列表
func StudentDataList(params *StudentQueryParam) []*Student {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := StudentPageList(params)
	return data
}

func GetGradeName(grade int) string {
	if _, ok := StudentGradeMap[grade]; !ok {
		return ""
	}
	return StudentGradeMap[grade]
}

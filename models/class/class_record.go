package class

import (
	"github.com/astaxie/beego/orm"
	"github.com/cuua/gocms/models"
	"github.com/cuua/gocms/models/student"
	"github.com/cuua/gocms/models/teacher"
	"github.com/cuua/gocms/utils"
)

func init() {
	orm.RegisterModel(new(ClassRecord))
}

const ClassRecordStatusInit = 0
const ClassRecordStatusConfirm = 1
const ClassRecordStatusDelete = -1

var ClassRate = []string{
	"6-8", "8-10", "10-12", "13-15", "15-17", "17-19", "19-21", "21-23",
}

type ClassRecord struct {
	Id                     int              `json:"id"`
	Teacher                *teacher.Teacher `orm:"rel(one)"`
	Student                *student.Student `orm:"rel(one)"`
	Subject                *teacher.Subject `orm:"rel(one)"`
	SchoolId               int              `json:"school_id"`
	ContractId             string           `json:"contract_id"`
	EndTime                string           `json:"end_time"`
	Length                 float64          `json:"length"`
	Status                 int              `json:"status"`
	Date                   string           `json:"date"`
	Time                   string           `json:"time"`
	Grade                  int              `json:"grade"`
	Price                  int              `json:"price"`
	Amount                 float64          `json:"amount"`
	Type                   int              `json:"type"`
	StudentPrice           string           `json:"student_price"`
	StudentAmount          float64          `json:"student_amount"`
	CreatedAt              string           `json:"created_at"`
	MsgState               int              `json:"msg_state"`
	StudentSurplus         float64          `json:"student_surplus"`
	StudentQuantitySurplus float64          `json:"student_quantity_surplus"`
	SchoolName             string           `json:"school_name"`
}

func (m *ClassRecord) TableName() string {
	return models.ClassRecordTBName()
}

// RoleQueryParam 用于搜索的类
type ClassRecordQueryParam struct {
	models.BaseQueryParam
	SchoolId        string
	TeacherId       int
	StudentNameLike string
	TeacherNameLike string
	DateStart       string
	DateEnd         string
	Status          string
}

// RoleDataList 获取列表
func ClassRecordDataList(params *ClassRecordQueryParam) []*ClassRecord {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := ClassRecordPageList(params)
	return data
}

// RolePageList 获取分页数据
func ClassRecordPageList(params *ClassRecordQueryParam) ([]*ClassRecord, int64) {
	query := orm.NewOrm().QueryTable(models.ClassRecordTBName()).RelatedSel("student", "teacher", "subject")
	data := make([]*ClassRecord, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	if params.StudentNameLike != "" {
		query = query.Filter("student__name__contains", params.StudentNameLike)
	}
	if params.TeacherNameLike != "" {
		query = query.Filter("teacher__name__contains", params.TeacherNameLike)
	}
	if params.DateStart != "" {
		query = query.Filter("date__gte", params.DateStart)
	}
	if params.DateEnd != "" {
		query = query.Filter("date__lte", params.DateEnd)
	}
	if params.Status != "" {
		query = query.Filter("status", params.Status)
	}
	if params.CurUser.IsSuper != true {
		query = query.Filter("student__school__id", params.CurUser.SchoolId)
		if params.CurUser.RealName == "教师组" {
			tea := teacher.Teacher{}
			orm.NewOrm().QueryTable(models.TeacherTBName()).Filter("cardno", params.CurUser.UserName).One(&tea, "id")
			query = query.Filter("teacher_id", tea.Id)
		}
	}
	total, _ := query.Count()
	query = query.OrderBy(sortorder).Limit(params.Limit, params.Offset)
	query.All(&data)
	return data, total
}

func SalaryPageList(params *ClassRecordQueryParam) ([]*ClassRecord, int64) {
	query := orm.NewOrm().QueryTable(models.ClassRecordTBName())
	if params.TeacherId > 0 { // 工资明细需查看科目
		query = query.RelatedSel("teacher", "student", "subject")
	} else {
		query = query.RelatedSel("teacher", "student")
	}
	data := make([]*ClassRecord, 0)
	//默认排序
	sortorder := "teacher__id"
	if params.TeacherId > 0 {
		query = query.Filter("teacher__id", params.TeacherId)
	}
	if params.DateStart != "" {
		query = query.Filter("date__gte", params.DateStart)
	}
	if params.DateEnd != "" {
		query = query.Filter("date__lte", params.DateEnd)
	}
	if params.SchoolId != "" {
		query = query.Filter("school_id", params.SchoolId)
	}
	if params.CurUser.IsSuper != true {
		query = query.Filter("student__school__id", params.CurUser.SchoolId)
	}
	query = query.Filter("status", ClassRecordStatusConfirm)
	query = query.GroupBy("teacher_id", "student_id", "subject_id", "date", "time")
	total, _ := query.Count()
	query = query.OrderBy(sortorder)
	if params.TeacherId > 0 {
		query = query.Limit(params.Limit, params.Offset)
	} else {
		query = query.Limit(total)
	}
	query.All(&data)
	return data, total
}

func UserateDataList(params *ClassRecordQueryParam) ([]*ClassRecord, int64) {
	query := orm.NewOrm().QueryTable(models.ClassRecordTBName()).RelatedSel("teacher")
	data := make([]*ClassRecord, 0)
	startDay, endDay := utils.MonthRange()
	//默认排序
	sortorder := "teacher__id"
	if params.TeacherId > 0 {
		query = query.Filter("teacher__id", params.TeacherId)
	}
	if params.DateStart != "" {
		query = query.Filter("date__gte", params.DateStart)
	} else {
		query = query.Filter("date__gte", startDay)
	}
	if params.DateEnd != "" {
		query = query.Filter("date__lte", params.DateEnd)
	} else {
		query = query.Filter("date__lte", endDay)
	}
	if params.SchoolId != "" {
		query = query.Filter("school_id", params.SchoolId)
	}
	if params.CurUser.IsSuper != true {
		query = query.Filter("teacher__school__id", params.CurUser.SchoolId)
	}
	query = query.GroupBy("teacher_id", "student_id", "subject_id", "date", "time")
	total, _ := query.Count()
	query = query.OrderBy(sortorder).Limit(-1)
	query.All(&data)
	return data, total
}

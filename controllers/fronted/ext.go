package fronted

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"gocms/enums"
	"gocms/models"
	"gocms/models/class"
	"gocms/models/student"
	"gocms/models/teacher"
	"gocms/utils"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type ExtController struct {
	BaseController
}

// 自动生成课表
func (c *ExtController) GenClass() {
	/*课表固定在星期天晚10点生成*/
	if time.Now().Weekday() != 0 || time.Now().Hour() < 22 {
		c.JsonResult(enums.JRCodeSucc, "time not match", "")
	}
	obj := make([]class.Schedule, 0)
	orm.NewOrm().QueryTable(models.ScheduleTBName()).Filter("status", class.ScheduleStatusValid).All(&obj)
	timeNow := time.Now()
	mStu := &student.Student{}
	mTea := &teacher.Teacher{}
	mSub := &teacher.Subject{}
	for _, val := range obj {
		immutable := reflect.ValueOf(val)
		for i := 1; i <= 56; i++ {
			m := class.ClassRecord{}
			classCol := "class" + strconv.Itoa(i) + "0"
			stuId := immutable.FieldByName(classCol).Int()
			length := immutable.FieldByName(classCol + "Length").Float()
			if length == 0 { // 默认为2
				length = 2
			}
			if stuId <= 0 {
				continue
			}
			var addDate, timeIndex int
			if i%7 != 0 {
				addDate = i % 7
				timeIndex = i / 7
			} else {
				addDate = 7
				timeIndex = i/7 - 1
			}
			date := timeNow.AddDate(0, 0, addDate)
			m.Teacher = mTea
			m.Student = mStu
			m.Subject = mSub
			m.SchoolId = val.SchoolId
			m.Teacher.Id = val.TeacherId
			m.Student.Id = int(stuId)
			m.Subject.Id = val.SubjectId
			m.Date = date.Format(utils.FormatDate)
			m.Time = class.ClassRate[timeIndex]
			m.Length = length
			m.CreatedAt = time.Now().Format(utils.FormatDateTime)
			/*查看教师是否休眠*/
			tea := teacher.Teacher{}
			orm.NewOrm().QueryTable(models.TeacherTBName()).Filter("id", val.TeacherId).One(&tea, "id", "name", "status")
			if tea.Name == "" || tea.Status == teacher.TeacherStatusDelete { // 教师不存在或已被删除
				continue
			}
			/*查看是否托班*/
			stu := student.Student{}
			if strings.Contains(tea.Name, "T-") {
				m.Type = 2
			} else {
				orm.NewOrm().QueryTable(models.StudentTBName()).Filter("id", stuId).One(&stu, "id", "type", "group_id", "status")
				if stu.Id == 0 || stu.Status != student.StudentStatusValid {
					continue
				}
				m.Type = stu.Type
			}
			if m.Type == student.TeamType { //如果是小班，需特殊处理
				stuIdArr := []int{}
				json.Unmarshal([]byte(stu.GroupId), stuIdArr)
				for _, val := range stuIdArr {
					m.Student.Id = val
					chk := orm.NewOrm().QueryTable(models.ClassRecordTBName()).Filter("student_id", m.Student.Id).Filter("teacher_id", m.Teacher.Id).Filter("subject_id", m.Subject.Id).Filter("date", m.Date).Filter("time", m.Time).Exist()
					if chk {
						continue
					}
					orm.NewOrm().Insert(&m)
				}
			} else {
				chk := orm.NewOrm().QueryTable(models.ClassRecordTBName()).Filter("student_id", m.Student.Id).Filter("teacher_id", m.Teacher.Id).Filter("subject_id", m.Subject.Id).Filter("date", m.Date).Filter("time", m.Time).Exist()
				if chk {
					continue
				}
				orm.NewOrm().Insert(&m)
			}
		}
	}
	c.JsonResult(enums.JRCodeSucc, "OK", "")
}

// 更改学生状态
func (c *ExtController) StudentStatus() {
	obj := make([]student.Student, 0)
	orm.NewOrm().QueryTable(models.StudentTBName()).
		Filter("type", student.StudentType).
		Filter("balance1__lte", 0).
		Filter("balance2__lte", 0).
		Filter("balance3__lte", 0).
		Filter("status", student.StudentStatusValid).
		All(&obj)
	for _, val := range obj {
		val.Status = student.StudentStatusInit
		con := orm.NewOrm().QueryTable(models.ContractTBName()).Filter("student_id", val.Id).Exist()
		if con {
			val.Status = student.StudentStatusStop
		}
		orm.NewOrm().Update(&val)
	}
	c.JsonResult(enums.JRCodeSucc, "ok", "")
}

// 班组状态
func (c *ExtController) TeamStatus() {
	obj := make([]student.Student, 0)
	orm.NewOrm().QueryTable(models.StudentTBName()).
		Filter("type", student.TeamType).
		All(&obj)
	for _, val := range obj {
		var idArr, newIdArr []int
		json.Unmarshal([]byte(val.GroupId), &idArr)
		for _, v := range idArr {
			exist := orm.NewOrm().
				QueryTable(models.ContractTBName()).
				Filter("student_id", v).
				Filter("surplus__gt", 0).
				Exist()
			if exist {
				newIdArr = append(newIdArr, v)
			}
		}
		if len(newIdArr) > 0 {
			jbyte, _ := json.Marshal(newIdArr)
			val.GroupId = string(jbyte)
			// 查找学生名
			stus := make([]student.Student, 0)
			groupName := ""
			orm.NewOrm().QueryTable(models.StudentTBName()).Filter("id__in", newIdArr).All(&stus)
			for _, s := range stus {
				groupName += s.Name + "|"
			}
			val.GroupName = groupName
			orm.NewOrm().Update(&val)
		} else {
			orm.NewOrm().Delete(&val)
		}
	}
	c.JsonResult(enums.JRCodeSucc, "ok", "")
}

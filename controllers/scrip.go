package controllers

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
)

type ScriptController struct {
	BaseController
}

func (c *ScriptController) GenClass() {
	m := []*class.Schedule{}
	orm.NewOrm().QueryTable(models.ScheduleTBName()).Filter("status", class.ScheduleStatusValid).All(&m)
	for _, val := range m {
		var classType int
		teacherId := val.TeacherId
		subjectId := val.SubjectId
		schoolId := val.SchoolId
		t := reflect.TypeOf(*val)
		v := reflect.ValueOf(*val)
		for k := 0; k < t.NumField(); k++ {
			if strings.Contains(t.Field(k).Name, "class") && !strings.Contains(t.Field(k).Name, "Length") {
				sid := v.Field(k).Interface()
				studentId, _ := sid.(int)
				if studentId > 0 {
					class := class.ClassRecord{}

					stu := student.Student{Id: studentId}
					tea := teacher.Teacher{Id: teacherId}
					err := orm.NewOrm().QueryTable(models.TeacherTBName()).Filter("id", teacherId).One(&tea, "id", "name", "status")
					if err != nil {
						continue
					}
					if tea.Status != teacher.TeacherStatusValid {
						continue
					}
					if strings.Contains(tea.Name, "T-") {
						classType = 2
					} else {
						err := orm.NewOrm().QueryTable(models.StudentTBName()).Filter("id", studentId).One(&stu, "id", "name", "type", "group_id")
						if err != nil {
							continue
						}
						classType = stu.Type
					}
					timeRange := utils.GetTimeRange(t.Field(k).Name)
					date := utils.GetDateByClass(t.Field(k).Name)
					lengthName := utils.GetLengthNameByClass(t.Field(k).Name)
					length := v.FieldByName(lengthName).Float()
					if length == 0.0 {
						length = 2.0
					}
					// 公共参数
					subject := teacher.Subject{Id: subjectId}
					class.Teacher = &tea
					class.Subject = &subject
					class.Student = &stu
					class.SchoolId = schoolId
					class.Date = date
					class.Time = timeRange
					class.Length = length
					class.Type = classType
					if classType == 1 { // 小班需查询班组成员
						jsonGroup := stu.GroupId
						byteGroup := []byte(jsonGroup)
						groupId := make([]string, 0)
						json.Unmarshal(byteGroup, &groupId)
						signalStu := student.Student{}
						for _, signal := range groupId {
							class.Id = 0
							signalStu.Id, _ = strconv.Atoi(signal)
							class.Student = &signalStu
							utils.JDebug(class)
							orm.NewOrm().ReadOrCreate(&class, "teacher_id", "student_id", "subject_id", "date", "time")
						}
					} else {
						utils.JDebug(class)
						orm.NewOrm().ReadOrCreate(&class, "teacher_id", "student_id", "subject_id", "date", "time")
					}
				}
			}
		}
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

// 更改学生状态
func (c *ScriptController) StudentStatus() {
	obj := []*student.Student{}
	orm.NewOrm().QueryTable(models.StudentTBName()).Filter("type", student.StudentType).All(&obj, "id", "balance1", "balance1_length", "balance2", "balance2_length", "balance3", "balance3_length", "status")
	for _, val := range obj {
		var status int
		if val.Balance1 <= 0 && val.Balance1Length <= 0 && val.Balance2 <= 0 && val.Balance2Length <= 0 && val.Balance3 <= 0 && val.Balance3Length <= 0 {
			cont := student.Contract{}
			err := orm.NewOrm().QueryTable(models.ContractTBName()).Filter("student_id", val.Id).One(&cont, "id")
			if err == nil {
				status = 2
			} else {
				status = 0
			}
		} else {
			status = 1
		}
		if val.Status != status {
			val.Status = status
			orm.NewOrm().Update(val, "status")
		}
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *ScriptController) TeamStatus() {
	obj := []*student.Student{}
	orm.NewOrm().QueryTable(models.StudentTBName()).Filter("type", student.TeamType).All(&obj, "id", "group_id", "status")
	for _, val := range obj {
		jsonGroup := val.GroupId
		byteGroup := []byte(jsonGroup)
		groupId := make([]string, 0)
		newGroupId := make([]string, 0)
		json.Unmarshal(byteGroup, &groupId)
		for _, stu := range groupId {
			cont := student.Contract{}
			err := orm.NewOrm().QueryTable(models.ContractTBName()).Filter("student_id", stu).Filter("type", student.ContractTypeXiao).Filter("surplus__gt", 0).One(&cont, "id")
			if err == nil { // 合同有余额
				newGroupId = append(newGroupId, stu)
			}
		}
		utils.JDebug(newGroupId)
		if len(newGroupId) > 0 {
			newJsonGroup, _ := json.Marshal(newGroupId)
			val.GroupId = string(newJsonGroup)
			orm.NewOrm().Update(val, "group_id")
		} else {
			val.Status = -1
			orm.NewOrm().Update(val, "status")
		}
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

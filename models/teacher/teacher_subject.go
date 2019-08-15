package teacher

import (
	"github.com/astaxie/beego/orm"
	"gocms/models"
)

func init() {
	orm.RegisterModel(new(TeacherSubject))
}

type TeacherSubject struct {
	Id          int    `json:"id"`
	TeacherId   int    `json:"teacher_id"`
	SubjectId   int    `json:"subject_id"`
	SubjectName string `json:"subject_name"`
}

func (m *TeacherSubject) TableName() string {
	return models.TeacherSubjectTBName()
}

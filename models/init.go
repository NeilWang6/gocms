package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// init 初始化
func init() {
	orm.RegisterModel(new(BackendUser), new(Resource), new(Role), new(RoleResourceRel), new(RoleBackendUserRel))
}

// TableName 下面是统一的表名管理
func TableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix")
	return prefix + name
}

// BackendUserTBName 获取 BackendUser 对应的表名称
func BackendUserTBName() string {
	return TableName("backend_user")
}

// ResourceTBName 获取 Resource 对应的表名称
func ResourceTBName() string {
	return TableName("resource")
}

// RoleTBName 获取 Role 对应的表名称
func RoleTBName() string {
	return TableName("role")
}

// RoleResourceRelTBName 角色与资源多对多关系表
func RoleResourceRelTBName() string {
	return TableName("role_resource_rel")
}

// RoleBackendUserRelTBName 角色与用户多对多关系表
func RoleBackendUserRelTBName() string {
	return TableName("role_backenduser_rel")
}
func StudentTBName() string {
	return TableName("student")
}
func SchoolTBName() string {
	return TableName("school")
}
func StudentSchoolTBName() string {
	return TableName("student_school")
}
func StudentAreaTBName() string {
	return TableName("student_area")
}
func TeacherTBName() string {
	return TableName("teacher")
}
func SubjectTBName() string {
	return TableName("subject")
}
func TeacherSubjectTBName() string {
	return TableName("teacher_subject")
}
func ScheduleTBName() string {
	return TableName("schedule")
}
func ContractTBName() string {
	return TableName("contract")
}
func ClassRecordTBName() string {
	return TableName("class_record")
}
func ContractPriceTBName() string {
	return TableName("contract_price")
}
func ExpendTBName() string {
	return TableName("expend")
}
func BannerTBName() string {
	return TableName("banner")
}
func NewsTBName() string {
	return TableName("news")
}
func ProjectTBName() string {
	return TableName("project")
}
func UsTBName() string {
	return TableName("us")
}
func StaffTBName() string {
	return TableName("staff")
}
func RecruitTBName() string {
	return TableName("recruit")
}
func HistoryTBName() string {
	return TableName("history")
}

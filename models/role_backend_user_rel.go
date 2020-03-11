package models

import "time"

// RoleBackendUserRel 角色与用户关系
type RoleBackendUserRel struct {
	Id          int
	Role        *Role        `orm:"rel(fk)"`  //外键
	BackendUser *BackendUser `orm:"rel(fk)" ` // 外键
	Created     time.Time    `orm:"auto_now_add;type(datetime)"`
}

// TableName 设置表名
func (a *RoleBackendUserRel) TableName() string {
	return RoleBackendUserRelTBName()
}

type RoleBackendUserRelQueryParam struct {
	BaseQueryParam
	ID            int
	BackendUserID int
	RoleID        int
	IDIn          []int
}

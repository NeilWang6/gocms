package models

// TableName 设置表名
func (a *Role) TableName() string {
	return RoleTBName()
}

// RoleQueryParam 用于搜索的类
type RoleQueryParam struct {
	BaseQueryParam
	NameLike string
}

// Role 用户角色 实体类
type Role struct {
	Id                 int    `form:"Id"`
	Name               string `form:"Name"`
	Seq                int
	RoleResourceRel    []*RoleResourceRel    `orm:"reverse(many)" json:"-"` // 设置一对多的反向关系
	RoleBackendUserRel []*RoleBackendUserRel `orm:"reverse(many)" json:"-"` // 设置一对多的反向关系
}

package models

// TableName 设置表名
func (a *Resource) TableName() string {
	return ResourceTBName()
}

// Resource 权限控制资源表
type Resource struct {
	Id              int
	Name            string    `orm:"size(64)"`
	Parent          *Resource `orm:"null;rel(fk)"` // RelForeignKey relation
	Rtype           int
	Seq             int
	Sons            []*Resource        `orm:"reverse(many)"` // fk 的反向关系
	SonNum          int                `orm:"-"`
	Icon            string             `orm:"size(32)"`
	LinkUrl         string             `orm:"-"`
	UrlFor          string             `orm:"size(256)" Json:"-"`
	HtmlDisabled    int                `orm:"-"`             //在html里应用时是否可用
	Level           int                `orm:"-"`             //第几级，从0开始
	RoleResourceRel []*RoleResourceRel `orm:"reverse(many)"` // 设置一对多的反向关系
}

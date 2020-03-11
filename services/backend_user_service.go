package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/cuua/gocms/models"
	"sync"
)

var BackendUserService = &backendUserService{
	mutex: &sync.Mutex{},
}

type backendUserService struct {
	mutex *sync.Mutex
}

// 创建
func (s *backendUserService) Create(m *models.BackendUser) (id int, err error) {
	err = db.Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func (s *backendUserService) First(m *models.BackendUser) (*models.BackendUser, error) {
	err := db.First(m).Error
	return m, err
}

// 根据主键更新
func (s *backendUserService) Save(m *models.BackendUser) error {
	return db.Save(m).Error
}

// BackendUserPageList 获取分页数据
func (s *backendUserService) BackendUserPageList(params *models.BackendUserQueryParam) ([]*models.BackendUser, int64) {
	query := orm.NewOrm().QueryTable(models.BackendUserTBName())
	data := make([]*models.BackendUser, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("username__istartswith", params.UserNameLike)
	query = query.Filter("realname__istartswith", params.RealNameLike)
	if len(params.Mobile) > 0 {
		query = query.Filter("mobile", params.Mobile)
	}
	if len(params.SearchStatus) > 0 {
		query = query.Filter("status", params.SearchStatus)
	}
	total, _ := query.Count()
	_, _ = query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// BackendUserOne 根据id获取单条
func (s *backendUserService) BackendUserOne(id int) (*models.BackendUser, error) {
	o := orm.NewOrm()
	m := models.BackendUser{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// BackendUserOneByUserName 根据用户名密码获取单条
func (s *backendUserService) BackendUserOneByUserName(username, userpwd string) (*models.BackendUser, error) {
	m := models.BackendUser{}
	err := orm.NewOrm().QueryTable(models.BackendUserTBName()).Filter("username", username).Filter("userpwd", userpwd).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

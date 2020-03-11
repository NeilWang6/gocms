package services

import (
	"github.com/cuua/gocms/models"
	"sync"
)

var RoleBackendUserRelService = &roleBackendUserRelService{
	mutex: &sync.Mutex{},
}

type roleBackendUserRelService struct {
	mutex *sync.Mutex
}

// 创建
func (s *roleBackendUserRelService) Create(m *models.RoleBackendUserRel) error {
	return db.Create(m).Error
}

func (s *roleBackendUserRelService) Find(filter models.RoleBackendUserRelQueryParam) []*models.RoleBackendUserRel {
	data := make([]*models.RoleBackendUserRel, 0)
	where := "1"
	whereArgs := []interface{}{}
	if filter.RoleID > 0 {
		where += " AND `role_id` = ?"
		whereArgs = append(whereArgs, filter.RoleID)
	}
	if filter.BackendUserID > 0 {
		where += " AND `backend_user_id` = ?"
		whereArgs = append(whereArgs, filter.BackendUserID)
	}
	db.Where(where, whereArgs...).Find(&data)
	return data
}

func (s *roleBackendUserRelService) Delete(filter models.RoleBackendUserRelQueryParam) (num int64, err error) {
	where := "1"
	whereArgs := []interface{}{}
	if len(filter.IDIn) > 0 {
		where += " AND `id` in (?)"
		whereArgs = append(whereArgs, filter.IDIn)
	}
	if filter.ID > 0 {
		where += " AND `id` = ?"
		whereArgs = append(whereArgs, filter.ID)
	}
	if filter.RoleID > 0 {
		where += " AND `role_id` = ?"
		whereArgs = append(whereArgs, filter.RoleID)
	}
	if filter.BackendUserID > 0 {
		where += " AND `backend_user_id` = ?"
		whereArgs = append(whereArgs, filter.BackendUserID)
	}
	err = db.Where(where, whereArgs...).Delete(models.RoleBackendUserRel{}).Error
	num = db.RowsAffected
	return
}

package services

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/cuua/gocms/models"
	"github.com/cuua/gocms/utils"
	"sync"
)

var ResourceService = &resourceService{
	mutex: &sync.Mutex{},
}

type resourceService struct {
	mutex *sync.Mutex
}

// ResourceOne 获取单条
func (s *resourceService) ResourceOne(id int) (*models.Resource, error) {
	o := orm.NewOrm()
	m := models.Resource{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// ResourceTreeGrid 获取treegrid顺序的列表
func (s *resourceService) ResourceTreeGrid() []*models.Resource {
	o := orm.NewOrm()
	query := o.QueryTable(models.ResourceTBName()).OrderBy("seq", "id")
	list := make([]*models.Resource, 0)
	_, _ = query.All(&list)
	return resourceList2TreeGrid(list)
}

// ResourceTreeGrid4Parent 获取可以成为某个节点父节点的列表
func (s *resourceService) ResourceTreeGrid4Parent(id int) []*models.Resource {
	tree := ResourceService.ResourceTreeGrid()
	if id == 0 {
		return tree
	}
	var index = -1
	//找出当前节点所在索引
	for i, _ := range tree {
		if tree[i].Id == id {
			index = i
			break
		}
	}
	if index == -1 {
		return tree
	} else {
		tree[index].HtmlDisabled = 1
		for _, item := range tree[index+1:] {
			if item.Level > tree[index].Level {
				item.HtmlDisabled = 1
			} else {
				break
			}
		}
	}
	return tree
}

func (s *resourceService) ResourceTreeGridByUserId(backuserid, maxrtype int) []*models.Resource {
	cachekey := fmt.Sprintf("rms_ResourceTreeGridByUserId_%v_%v", backuserid, maxrtype)
	var list []*models.Resource
	if err := utils.GetCache(cachekey, &list); err == nil {
		return list
	}
	o := orm.NewOrm()
	user, err := BackendUserService.BackendUserOne(backuserid)
	if err != nil || user == nil {
		return list
	}
	var sql string
	if user.IsSuper == true {
		//如果是管理员，则查出所有的
		sql = fmt.Sprintf(`SELECT id,name,parent_id,rtype,icon,seq,url_for FROM %s Where rtype <= ? Order By seq asc,Id asc`, models.ResourceTBName())
		_, _ = o.Raw(sql, maxrtype).QueryRows(&list)
	} else {
		//联查多张表，找出某用户有权管理的
		sql = fmt.Sprintf(`SELECT DISTINCT T0.resource_id,T2.id,T2.name,T2.parent_id,T2.rtype,T2.icon,T2.seq,T2.url_for
		FROM %s AS T0
		INNER JOIN %s AS T1 ON T0.role_id = T1.role_id
		INNER JOIN %s AS T2 ON T2.id = T0.resource_id
		WHERE T1.backend_user_id = ? and T2.rtype <= ?  Order By T2.seq asc,T2.id asc`, models.RoleResourceRelTBName(), models.RoleBackendUserRelTBName(), models.ResourceTBName())
		_, _ = o.Raw(sql, backuserid, maxrtype).QueryRows(&list)
	}
	result := resourceList2TreeGrid(list)
	_ = utils.SetCache(cachekey, result, 30)
	return result
}

// resourceList2TreeGrid 将资源列表转成treegrid格式
func resourceList2TreeGrid(list []*models.Resource) []*models.Resource {
	result := make([]*models.Resource, 0)
	for _, item := range list {
		if item.Parent == nil || item.Parent.Id == 0 {
			item.Level = 0
			result = append(result, item)
			result = resourceAddSons(item, list, result)
		}
	}
	return result
}

//resourceAddSons 添加子菜单
func resourceAddSons(cur *models.Resource, list, result []*models.Resource) []*models.Resource {
	for _, item := range list {
		if item.Parent != nil && item.Parent.Id == cur.Id {
			cur.SonNum++
			item.Level = cur.Level + 1
			result = append(result, item)
			result = resourceAddSons(item, list, result)
		}
	}
	return result
}

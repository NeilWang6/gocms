package controllers

import (
	"fmt"
	"strings"

	"github.com/cuua/gocms/enums"
	"github.com/cuua/gocms/models"
	"github.com/cuua/gocms/utils"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	ControllerName string             //当前控制名称
	ActionName     string             //当前action名称
	CurUser        models.BackendUser //当前用户信息
}

func (c *BaseController) Prepare() {
	//附值
	c.ControllerName, c.ActionName = c.GetControllerAndAction()
	//从Session里获取数据 设置用户信息
	c.adapterUserInfo()
}

// checkLogin判断用户是否登录，未登录则跳转至登录页面
// 一定要在BaseController.Prepare()后执行
func (c *BaseController) checkLogin() {
	if c.CurUser.Id == 0 {
		//登录页面地址
		urlstr := c.URLFor("HomeController.Login") + "?url="
		//登录成功后返回的址为当前
		returnURL := c.Ctx.Request.URL.Path
		//如果ajax请求则返回相应的错码和跳转的地址
		if c.Ctx.Input.IsAjax() {
			//由于是ajax请求，因此地址是header里的Referer
			returnURL := c.Ctx.Input.Refer()
			c.JsonResult(enums.JRCode302, "请登录", urlstr+returnURL)
		}
		c.Redirect(urlstr+returnURL, 302)
		c.StopRun()
	}
}

// 判断某 Controller.Action 当前用户是否有权访问
func (c *BaseController) CheckActionAuthor(ctrlName, ActName string) bool {
	if c.CurUser.Id == 0 {
		return false
	}
	//从session获取用户信息
	user := c.GetSession("backenduser")
	//类型断言
	v, ok := user.(models.BackendUser)
	if ok {
		//如果是超级管理员，则直接通过
		if v.IsSuper == true {
			return true
		}
		//遍历用户所负责的资源列表
		for i, _ := range v.ResourceUrlForList {
			urlfor := strings.TrimSpace(v.ResourceUrlForList[i])
			if len(urlfor) == 0 {
				continue
			}
			// TestController.Get,:last,xie,:first,asta
			strs := strings.Split(urlfor, ",")
			if len(strs) > 0 && strs[0] == (ctrlName+"."+ActName) {
				return true
			}
		}
	}
	return false
}

// checkLogin判断用户是否有权访问某地址，无权则会跳转到错误页面
//一定要在BaseController.Prepare()后执行
// 会调用checkLogin
// 传入的参数为忽略权限控制的Action
func (c *BaseController) CheckAuthor(ignores ...string) {
	//先判断是否登录
	c.checkLogin()
	//如果Action在忽略列表里，则直接通用
	for _, ignore := range ignores {
		if ignore == c.ActionName {
			return
		}
	}
	hasAuthor := c.CheckActionAuthor(c.ControllerName, c.ActionName)
	if !hasAuthor {
		utils.LogDebug(fmt.Sprintf("author control: path=%s.%s userid=%v  无权访问", c.ControllerName, c.ActionName, c.CurUser.Id))
		//如果没有权限
		if !hasAuthor {
			if c.Ctx.Input.IsAjax() {
				c.JsonResult(enums.JRCode401, "无权访问", "")
			} else {
				c.PageError("无权访问")
			}
		}
	}
}

//从session里取用户信息
func (c *BaseController) adapterUserInfo() {
	a := c.GetSession("backenduser")
	if a != nil {
		c.CurUser = a.(models.BackendUser)
		c.Data["backenduser"] = a
	}
}

//SetBackendUser2Session 获取用户信息（包括资源UrlFor）保存至Session
func (c *BaseController) setBackendUser2Session(userId int) error {
	m, err := models.BackendUserOne(userId)
	if err != nil {
		return err
	}
	//获取这个用户能获取到的所有资源列表
	resourceList := models.ResourceTreeGridByUserId(userId, 1000)
	for _, item := range resourceList {
		m.ResourceUrlForList = append(m.ResourceUrlForList, strings.TrimSpace(item.UrlFor))
	}
	c.SetSession("backenduser", *m)
	return nil
}

// 设置模板
// 第一个参数模板，第二个参数为layout
func (c *BaseController) SetTpl(template ...string) {
	var tplName string
	layout := "shared/layout_page.html"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		//不要Controller这个10个字母
		ctrlName := strings.ToLower(c.ControllerName[0 : len(c.ControllerName)-10])
		ActionName := strings.ToLower(c.ActionName)
		tplName = ctrlName + "/" + ActionName + ".html"
	}
	c.Layout = layout
	c.TplName = tplName
}
func (c *BaseController) JsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	r := &models.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

// 重定向
func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

// 重定向 去错误页
func (c *BaseController) PageError(msg string) {
	errorurl := c.URLFor("HomeController.Error") + "/" + msg
	c.Redirect(errorurl, 302)
	c.StopRun()
}

// 重定向 去登录页
func (c *BaseController) PageLogin() {
	url := c.URLFor("HomeController.Login")
	c.Redirect(url, 302)
	c.StopRun()
}

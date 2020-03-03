package student

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/cuua/gocms/controllers"
	"github.com/cuua/gocms/enums"
	"github.com/cuua/gocms/models"
	"github.com/cuua/gocms/models/student"
	"github.com/cuua/gocms/utils"
	"strconv"
	"strings"
	"time"
)

type StudentController struct {
	controllers.BaseController
}

func (c *StudentController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.CheckAuthor()

}

// 学生信息列表
func (c *StudentController) Index() {
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "student/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "student/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.CheckActionAuthor("StudentController", "Edit")
	c.Data["canDelete"] = c.CheckActionAuthor("StudentController", "Delete")
	c.Data["canAllocate"] = c.CheckActionAuthor("StudentController", "Allocate")
}

// 班组列表
func (c *StudentController) Team() {
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "student/team_headcssjs.html"
	c.LayoutSections["footerjs"] = "student/team_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.CheckActionAuthor("StudentController", "TeamEdit")
	c.Data["canDelete"] = c.CheckActionAuthor("StudentController", "TeamDelete")
}

// DataGrid 表格获取数据
func (c *StudentController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params student.StudentQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	params.Type = student.StudentType
	params.CurUser = c.CurUser // 用户传入用于数据权限控制
	data, total := student.StudentPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

//DataList 学生列表
func (c *StudentController) DataList() {
	var params = student.StudentQueryParam{}
	//获取form里的值
	if err := c.ParseForm(&params); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", "")
	}
	//params.Type = c.GetString("Type")
	params.CurUser = c.CurUser
	//获取数据列表和总数
	data := student.StudentDataList(&params)
	//定义返回的数据结构
	c.JsonResult(enums.JRCodeSucc, "", data)
}

// DataGrid 角色管理首页 表格获取数据
func (c *StudentController) TeamDataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params student.StudentQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	params.Type = student.TeamType // 查询班组
	params.CurUser = c.CurUser
	data, total := student.StudentPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

//Edit 添加、编辑角色界面
func (c *StudentController) Edit() {
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := student.Student{Id: Id}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.QueryTable(models.StudentTBName()).Filter("id", Id).RelatedSel().One(&m)
		if err != nil {
			c.PageError("数据无效，请刷新后重试")
		}
	}
	//校区信息
	if m.School == nil {
		m.School = &student.School{}
	}
	if m.StudentSchool == nil {
		m.StudentSchool = &student.StudentSchool{}
	}
	c.Data["m"] = m
	c.SetTpl("student/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "student/edit_footerjs.html"
}

//Edit 添加、编辑角色界面
func (c *StudentController) TeamEdit() {
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := student.Student{Id: Id}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.QueryTable(models.StudentTBName()).Filter("id", Id).One(&m)
		if err != nil {
			c.PageError("数据无效，请刷新后重试")
		}
	}
	//校区信息
	if m.School == nil {
		m.School = &student.School{Id: c.CurUser.SchoolId}
	}
	if m.StudentSchool == nil {
		m.StudentSchool = &student.StudentSchool{}
	}
	c.Data["m"] = m
	c.SetTpl("student/teamedit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "student/teamedit_footerjs.html"
}

//Save 添加、编辑页面 保存
func (c *StudentController) Save() {
	var err error
	m := student.Student{}
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	schoolId, _ := c.GetInt("SchoolId")
	studentSchoolId, _ := c.GetInt("StudentSchoolId")
	school := student.School{Id: schoolId}
	studentSchool := student.StudentSchool{Id: studentSchoolId}
	m.School = &school
	m.StudentSchool = &studentSchool
	o := orm.NewOrm()
	if m.Id == 0 {
		lastS := student.Student{}
		err = o.QueryTable(models.StudentTBName()).OrderBy("-id").One(&lastS)
		// 生成学生编号锁
		StudentCodeValid()
		//创建时间，编号创建
		m.CreatedAt = time.Now()
		m.Code = GenCode(lastS.Code)
		if m.Type == student.TeamType { // 班组添加
			groupIds := make([]string, 0)
			c.Ctx.Input.Bind(&groupIds, "GroupId")
			tmp, _ := json.Marshal(groupIds)
			m.GroupId = string(tmp)
			m.GroupName = GetStudentName(groupIds)
		}
		_, err = o.Insert(&m)
		if err == nil {
			c.JsonResult(enums.JRCodeSucc, "添加成功", m.Id)
		} else {
			c.JsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		}
	} else {
		m.UpdatedAt = time.Now()
		if m.Type == student.StudentType {
			_, err = o.Update(&m, "name", "sex", "guarder", "relate", "grade", "contact1", "school_id", "student_school_id", "adress", "note", "status", "updated_at")
		} else { // 班组更新
			groupIds := make([]string, 0)
			c.Ctx.Input.Bind(&groupIds, "GroupId")
			tmp, _ := json.Marshal(groupIds)
			m.GroupId = string(tmp)
			m.GroupName = GetStudentName(groupIds)
			_, err = o.Update(&m, "name", "group_id", "group_name", "status", "updated_at")
		}
		if err == nil {
			c.JsonResult(enums.JRCodeSucc, "编辑成功", m.Id)
		} else {
			c.JsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}
}

//Delete 批量删除
func (c *StudentController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	if num, err := student.StudentDelete(ids); err == nil {
		c.JsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.JsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

func (c *StudentController) Search() {
	code := c.GetString("Code")
	m := student.Student{}
	orm.NewOrm().QueryTable(models.StudentTBName()).Filter("code", code).One(&m, "id", "name", "school_id", "guarder", "contact1", "grade")
	c.JsonResult(enums.JRCodeSucc, "succ", m)
}

// 学生统计
func (c *StudentController) Total() {
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "student/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "student/total_footerjs.html"
}

type Flot struct {
	Label string `json:"label"`
	Data  int    `json:"data"`
	Color string `json:"color"`
}

// 新生统计
func (c *StudentController) TotalData() {
	dateStart := c.GetString("DateStart")
	dateEnd := c.GetString("DateEnd")
	schoolId := c.GetString("SchoolId")
	count := 0
	// 按状态分
	qb, _ := orm.NewQueryBuilder("mysql")
	qb = qb.Select("status", "count(*) as count").From(models.StudentTBName()).Where("type = ?").And("status >= ?")
	if dateStart != "" {
		qb = qb.And("DATE_FORMAT(created_at, '%Y-%m-%d') >= '" + dateStart + "'")
	}
	if dateEnd != "" {
		qb = qb.And("DATE_FORMAT(created_at, '%Y-%m-%d') <= '" + dateEnd + "'")
	}
	if schoolId != "" {
		qb = qb.And("school_id = " + schoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql := qb.GroupBy("status").OrderBy("count DESC").String()
	type statusStruct struct {
		student.Student
		Count int
	}
	m := []statusStruct{}
	orm.NewOrm().Raw(sql, student.StudentType, 0).QueryRows(&m)
	statusObj := make([]Flot, 0)
	tmp := Flot{}
	for i := 0; i < len(m); i++ {
		count += m[i].Count
		tmp.Label = student.StudentTypeMap[m[i].Status]
		tmp.Data = m[i].Count
		tmp.Color = utils.GetColor(i)
		if len(statusObj) <= 5 {
			statusObj = append(statusObj, tmp)
		}
	}
	// 按状态分
	// 按学校分
	qb, _ = orm.NewQueryBuilder("mysql")
	qb = qb.Select("school.name", "count(*) as count").From(models.StudentTBName() + " AS student").LeftJoin(models.StudentSchoolTBName() + " AS school").On("student.student_school_id = school.id").Where("type = ?").And("status >= ?").And("student_school_id > ?")
	if dateStart != "" {
		qb = qb.And("DATE_FORMAT(student.created_at, '%Y-%m-%d') >= '" + dateStart + "'")
	}
	if dateEnd != "" {
		qb = qb.And("DATE_FORMAT(student.created_at, '%Y-%m-%d') <= '" + dateEnd + "'")
	}
	if schoolId != "" {
		qb = qb.And("school_id = " + schoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql = qb.GroupBy("student_school_id").OrderBy("count DESC").String()
	type schoolStruct struct {
		student.StudentSchool
		Count int
	}
	sm := []schoolStruct{}
	orm.NewOrm().Raw(sql, student.StudentType, 0, 0).QueryRows(&sm)
	schoolObj := make([]Flot, 0)
	for key, val := range sm {
		tmp.Label = val.Name
		tmp.Data = val.Count
		tmp.Color = utils.GetColor(key)
		if len(schoolObj) <= 5 {
			schoolObj = append(schoolObj, tmp)
		}

	}
	// 按住址分
	qb, _ = orm.NewQueryBuilder("mysql")
	qb = qb.Select("count(*) as count", "adress").From(models.StudentTBName()).Where("type = ?").And("status >= ?").And("adress <> ''")
	if dateStart != "" {
		qb = qb.And("DATE_FORMAT(created_at, '%Y-%m-%d') >= '" + dateStart + "'")
	}
	if dateEnd != "" {
		qb = qb.And("DATE_FORMAT(created_at, '%Y-%m-%d') <= '" + dateEnd + "'")
	}
	if schoolId != "" {
		qb = qb.And("school_id = " + schoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql = qb.GroupBy("adress").OrderBy("count DESC").String()
	type adressStruct struct {
		student.Student
		Count int
	}
	am := []adressStruct{}
	orm.NewOrm().Raw(sql, student.StudentType, 0).QueryRows(&am)
	adressObj := make([]Flot, 0)
	for key, val := range am {
		tmp.Label = val.Adress
		tmp.Data = val.Count
		tmp.Color = utils.GetColor(key)
		if len(adressObj) <= 5 {
			adressObj = append(adressObj, tmp)
		}
	}
	// 按年级分
	qb, _ = orm.NewQueryBuilder("mysql")
	qb = qb.Select("count(*) as count", "grade").From(models.StudentTBName()).Where("type = ?").And("status >= ?")
	if dateStart != "" {
		qb = qb.And("DATE_FORMAT(created_at, '%Y-%m-%d') >= '" + dateStart + "'")
	}
	if dateEnd != "" {
		qb = qb.And("DATE_FORMAT(created_at, '%Y-%m-%d') <= '" + dateEnd + "'")
	}
	if schoolId != "" {
		qb = qb.And("school_id = " + schoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql = qb.GroupBy("grade").OrderBy("count DESC").String()
	gm := []adressStruct{}
	orm.NewOrm().Raw(sql, student.StudentType, 0).QueryRows(&gm)
	gradeObj := make([]Flot, 0)
	for key, val := range gm {
		tmp.Label = student.GetGradeName(val.Grade)
		tmp.Data = val.Count
		tmp.Color = utils.GetColor(key)
		if len(gradeObj) <= 5 {
			gradeObj = append(gradeObj, tmp)
		}
	}
	// 按合同类型
	qb, _ = orm.NewQueryBuilder("mysql")
	qb = qb.Select("count(*) as count", "type").From(models.ContractTBName()).Where("status = ?")
	if dateStart != "" {
		qb = qb.And("DATE_FORMAT(created_at, '%Y-%m-%d') >= '" + dateStart + "'")
	}
	if dateEnd != "" {
		qb = qb.And("DATE_FORMAT(created_at, '%Y-%m-%d') <= '" + dateEnd + "'")
	}
	sql = qb.GroupBy("type").OrderBy("count DESC").String()
	type contractStruct struct {
		student.Contract
		Count int
	}
	cm := []contractStruct{}
	orm.NewOrm().Raw(sql, student.ContractStatusValid).QueryRows(&cm)
	contractObj := make([]Flot, 0)
	for key, val := range cm {
		tmp.Label = student.GetContractTypeName(val.Type)
		tmp.Data = val.Count
		tmp.Color = utils.GetColor(key)
		if len(contractObj) <= 5 {
			contractObj = append(contractObj, tmp)
		}
	}
	// 按关系
	// 按年级分
	qb, _ = orm.NewQueryBuilder("mysql")
	qb = qb.Select("count(*) as count", "relate").From(models.StudentTBName()).Where("type = ?").And("status >= ?")
	if dateStart != "" {
		qb = qb.And("DATE_FORMAT(created_at, '%Y-%m-%d') >= '" + dateStart + "'")
	}
	if dateEnd != "" {
		qb = qb.And("DATE_FORMAT(created_at, '%Y-%m-%d') <= '" + dateEnd + "'")
	}
	if schoolId != "" {
		qb = qb.And("school_id = " + schoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql = qb.GroupBy("relate").OrderBy("count DESC").String()
	rm := []adressStruct{}
	orm.NewOrm().Raw(sql, student.StudentType, 0).QueryRows(&rm)
	relateObj := make([]Flot, 0)
	for key, val := range rm {
		tmp.Label = val.Relate
		tmp.Data = val.Count
		tmp.Color = utils.GetColor(key)
		if len(relateObj) <= 5 {
			relateObj = append(relateObj, tmp)
		}
	}
	ret := map[string]interface{}{"Status": statusObj, "School": schoolObj, "Adress": adressObj, "Grade": gradeObj, "Contract": contractObj, "Relate": relateObj, "Count": count}
	c.JsonResult(enums.JRCodeSucc, "succ", ret)
}
func (c *StudentController) New() {
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ControllerName + "." + c.ActionName)
	c.SetTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "student/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "student/new_footerjs.html"
}

// 新生统计数据
func (c *StudentController) NewData() {
	dateStart := c.GetString("DateStart")
	dateEnd := c.GetString("DateEnd")
	schoolId := c.GetString("SchoolId")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("DATE_FORMAT(created_at,'%Y%u') weeks", "count(*) count").From(models.StudentTBName()).Where("type = ?")
	if dateStart != "" {
		qb = qb.And("DATE_FORMAT(created_at, '%Y-%m-%d') >= '" + dateStart + "'")
	}
	if dateEnd != "" {
		qb = qb.And("DATE_FORMAT(created_at, '%Y-%m-%d') <= '" + dateEnd + "'")
	}
	if schoolId != "" {
		qb = qb.And("school_id = " + schoolId)
	}
	if c.CurUser.IsSuper == false {
		qb = qb.And("school_id = " + strconv.Itoa(c.CurUser.SchoolId))
	}
	sql := qb.GroupBy("weeks").String()
	type NewTotal struct {
		Weeks string
		Count int
	}
	m := make([]NewTotal, 0)
	ret := [][]interface{}{}
	orm.NewOrm().Raw(sql, student.StudentType).QueryRows(&m)
	for _, val := range m {
		tmp := []interface{}{}
		tmp = append(tmp, val.Weeks)
		tmp = append(tmp, val.Count)
		ret = append(ret, tmp)
	}
	c.JsonResult(enums.JRCodeSucc, "succ", ret)
}

// 学生编号锁
func StudentCodeValid() {
	redis := utils.GetRedis()
	for {
		_, err := redis.Get("valid").Result()
		if err != nil {
			break
		}
	}
	redis.SetNX("valid", "1", 3*time.Second)
}

// 生成学生编号
func GenCode(last string) string {
	length := 12
	code := "sh"
	if last == "" {
		return "sh0000000001"
	}
	index := strings.Trim(last, "sh")
	onlyIndex := strings.Trim(index, "0")
	lenIndex := len(onlyIndex)
	intIndex, _ := strconv.Atoi(index)
	for i := 0; i < length-2-lenIndex; i++ {
		code += "0"
	}
	code += strconv.Itoa(intIndex + 1)
	return code
}

// 班组学生名获取
func GetStudentName(idArr []string) string {
	student := []student.Student{}
	orm.NewOrm().QueryTable(models.StudentTBName()).Filter("id__in", idArr).All(&student)
	name := ""
	for _, val := range student {
		name += val.Name + "|"
	}
	name = strings.Trim(name, "|")
	return name
}

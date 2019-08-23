package class

import (
	"github.com/astaxie/beego/orm"
	"gocms/models"
)

const ScheduleStatusValid = 0
const ScheduleStatusStop = -1

func init() {
	orm.RegisterModel(new(Schedule))
}

type Schedule struct {
	Id             int     `json:"id"`
	SubjectId      int     `json:"subject_id"`
	TeacherId      int     `json:"teacher_id"`
	SchoolId       int     `json:"school_id"`
	Class10        int     `json:"class10"`
	Class20        int     `json:"class20"`
	Class30        int     `json:"class30"`
	Class40        int     `json:"class40"`
	Class50        int     `json:"class50"`
	Class60        int     `json:"class60"`
	Class70        int     `json:"class70"`
	Class80        int     `json:"class80"`
	Class90        int     `json:"class90"`
	Class100       int     `json:"class100"`
	Class110       int     `json:"class110"`
	Class120       int     `json:"class120"`
	Class130       int     `json:"class130"`
	Class140       int     `json:"class140"`
	Class150       int     `json:"class150"`
	Class160       int     `json:"class160"`
	Class170       int     `json:"class170"`
	Class180       int     `json:"class180"`
	Class190       int     `json:"class190"`
	Class200       int     `json:"class200"`
	Class210       int     `json:"class210"`
	Class220       int     `json:"class220"`
	Class230       int     `json:"class230"`
	Class240       int     `json:"class240"`
	Class250       int     `json:"class250"`
	Class260       int     `json:"class260"`
	Class270       int     `json:"class270"`
	Class280       int     `json:"class280"`
	Class290       int     `json:"class290"`
	Class300       int     `json:"class300"`
	Class310       int     `json:"class310"`
	Class320       int     `json:"class320"`
	Class330       int     `json:"class330"`
	Class340       int     `json:"class340"`
	Class350       int     `json:"class350"`
	Class360       int     `json:"class360"`
	Class370       int     `json:"class370"`
	Class380       int     `json:"class380"`
	Class390       int     `json:"class390"`
	Class400       int     `json:"class400"`
	Class410       int     `json:"class410"`
	Class420       int     `json:"class420"`
	Class430       int     `json:"class430"`
	Class440       int     `json:"class440"`
	Class450       int     `json:"class450"`
	Class460       int     `json:"class460"`
	Class470       int     `json:"class470"`
	Class480       int     `json:"class480"`
	Class490       int     `json:"class490"`
	Class500       int     `json:"class500"`
	Class510       int     `json:"class510"`
	Class520       int     `json:"class520"`
	Class530       int     `json:"class530"`
	Class540       int     `json:"class540"`
	Class550       int     `json:"class550"`
	Class560       int     `json:"class560"`
	Class10Length  float64 `json:"class10_length"`
	Class20Length  float64 `json:"class20_length"`
	Class30Length  float64 `json:"class30_length"`
	Class40Length  float64 `json:"class40_length"`
	Class50Length  float64 `json:"class50_length"`
	Class60Length  float64 `json:"class60_length"`
	Class70Length  float64 `json:"class70_length"`
	Class80Length  float64 `json:"class80_length"`
	Class90Length  float64 `json:"class90_length"`
	Class100Length float64 `json:"class100_length"`
	Class110Length float64 `json:"class110_length"`
	Class120Length float64 `json:"class120_length"`
	Class130Length float64 `json:"class130_length"`
	Class140Length float64 `json:"class140_length"`
	Class150Length float64 `json:"class150_length"`
	Class160Length float64 `json:"class160_length"`
	Class170Length float64 `json:"class170_length"`
	Class180Length float64 `json:"class180_length"`
	Class190Length float64 `json:"class190_length"`
	Class200Length float64 `json:"class200_length"`
	Class210Length float64 `json:"class210_length"`
	Class220Length float64 `json:"class220_length"`
	Class230Length float64 `json:"class230_length"`
	Class240Length float64 `json:"class240_length"`
	Class250Length float64 `json:"class250_length"`
	Class260Length float64 `json:"class260_length"`
	Class270Length float64 `json:"class270_length"`
	Class280Length float64 `json:"class280_length"`
	Class290Length float64 `json:"class290_length"`
	Class300Length float64 `json:"class300_length"`
	Class310Length float64 `json:"class310_length"`
	Class320Length float64 `json:"class320_length"`
	Class330Length float64 `json:"class330_length"`
	Class340Length float64 `json:"class340_length"`
	Class350Length float64 `json:"class350_length"`
	Class360Length float64 `json:"class360_length"`
	Class370Length float64 `json:"class370_length"`
	Class380Length float64 `json:"class380_length"`
	Class390Length float64 `json:"class390_length"`
	Class400Length float64 `json:"class400_length"`
	Class410Length float64 `json:"class410_length"`
	Class420Length float64 `json:"class420_length"`
	Class430Length float64 `json:"class430_length"`
	Class440Length float64 `json:"class440_length"`
	Class450Length float64 `json:"class450_length"`
	Class460Length float64 `json:"class460_length"`
	Class470Length float64 `json:"class470_length"`
	Class480Length float64 `json:"class480_length"`
	Class490Length float64 `json:"class490_length"`
	Class500Length float64 `json:"class500_length"`
	Class510Length float64 `json:"class510_length"`
	Class520Length float64 `json:"class520_length"`
	Class530Length float64 `json:"class530_length"`
	Class540Length float64 `json:"class540_length"`
	Class550Length float64 `json:"class550_length"`
	Class560Length float64 `json:"class560_length"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	Status         int     `json:"status"`
}
type ScheduleWithName struct {
	Schedule
	Class10Name  string `json:"class10_name"`
	Class20Name  string `json:"class20_name"`
	Class30Name  string `json:"class30_name"`
	Class40Name  string `json:"class40_name"`
	Class50Name  string `json:"class50_name"`
	Class60Name  string `json:"class60_name"`
	Class70Name  string `json:"class70_name"`
	Class80Name  string `json:"class80_name"`
	Class90Name  string `json:"class90_name"`
	Class100Name string `json:"class100_name"`
	Class110Name string `json:"class110_name"`
	Class120Name string `json:"class120_name"`
	Class130Name string `json:"class130_name"`
	Class140Name string `json:"class140_name"`
	Class150Name string `json:"class150_name"`
	Class160Name string `json:"class160_name"`
	Class170Name string `json:"class170_name"`
	Class180Name string `json:"class180_name"`
	Class190Name string `json:"class190_name"`
	Class200Name string `json:"class200_name"`
	Class210Name string `json:"class210_name"`
	Class220Name string `json:"class220_name"`
	Class230Name string `json:"class230_name"`
	Class240Name string `json:"class240_name"`
	Class250Name string `json:"class250_name"`
	Class260Name string `json:"class260_name"`
	Class270Name string `json:"class270_name"`
	Class280Name string `json:"class280_name"`
	Class290Name string `json:"class290_name"`
	Class300Name string `json:"class300_name"`
	Class310Name string `json:"class310_name"`
	Class320Name string `json:"class320_name"`
	Class330Name string `json:"class330_name"`
	Class340Name string `json:"class340_name"`
	Class350Name string `json:"class350_name"`
	Class360Name string `json:"class360_name"`
	Class370Name string `json:"class370_name"`
	Class380Name string `json:"class380_name"`
	Class390Name string `json:"class390_name"`
	Class400Name string `json:"class400_name"`
	Class410Name string `json:"class410_name"`
	Class420Name string `json:"class420_name"`
	Class430Name string `json:"class430_name"`
	Class440Name string `json:"class440_name"`
	Class450Name string `json:"class450_name"`
	Class460Name string `json:"class460_name"`
	Class470Name string `json:"class470_name"`
	Class480Name string `json:"class480_name"`
	Class490Name string `json:"class490_name"`
	Class500Name string `json:"class500_name"`
	Class510Name string `json:"class510_name"`
	Class520Name string `json:"class520_name"`
	Class530Name string `json:"class530_name"`
	Class540Name string `json:"class540_name"`
	Class550Name string `json:"class550_name"`
	Class560Name string `json:"class560_name"`
	TeacherName  string `json:"teacher_name"`
	SubjectName  string `json:"subject_name"`
}

func (m *Schedule) TableName() string {
	return models.ScheduleTBName()
}

type ScheduleQueryParam struct {
	models.BaseQueryParam
}

// ScheduleDataList 获取列表
func ScheduleDataList(params *ScheduleQueryParam) []*Schedule {
	params.Limit = -1
	params.Sort = "Seq"
	params.Order = "asc"
	data, _ := SchedulePageList(params)
	return data
}

// SchedulePageList 获取分页数据
func SchedulePageList(params *ScheduleQueryParam) ([]*Schedule, int64) {
	query := orm.NewOrm().QueryTable(models.ScheduleTBName())
	data := make([]*Schedule, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

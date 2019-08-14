package student

import (
	"github.com/astaxie/beego/orm"
	"sdrms/models"
)

const ContractPriceStatusUnlock = 0
const ContractPriceStatusLock = 1

func init() {
	orm.RegisterModel(new(ContractPrice))
}

type ContractPrice struct {
	Id        int `json:"id"`
	Type      int `json:"type"`
	Grade     int `json:"grade"`
	SchoolId  int `json:"school_id"`
	Lower     int `json:"lower"`
	Upper     int `json:"upper"`
	Status1   int `json:"status1"`
	Status2   int `json:"status2"`
	Status3   int `json:"status3"`
	Status10  int `json:"status10"`
	Status20  int `json:"status20"`
	Status30  int `json:"status30"`
	Status40  int `json:"status40"`
	Status50  int `json:"status50"`
	Status60  int `json:"status60"`
	Status70  int `json:"status70"`
	Status80  int `json:"status80"`
	Status90  int `json:"status90"`
	Status100 int `json:"status100"`
	Status110 int `json:"status110"`
	Status120 int `json:"status120"`
	Grade1    int `json:"grade1"`
	Grade2    int `json:"grade2"`
	Grade3    int `json:"grade3"`
	Grade10   int `json:"grade10"`
	Grade20   int `json:"grade20"`
	Grade30   int `json:"grade30"`
	Grade40   int `json:"grade40"`
	Grade50   int `json:"grade50"`
	Grade60   int `json:"grade60"`
	Grade70   int `json:"grade70"`
	Grade80   int `json:"grade80"`
	Grade90   int `json:"grade90"`
	Grade100  int `json:"grade100"`
	Grade110  int `json:"grade110"`
	Grade120  int `json:"grade120"`
}

func (m *ContractPrice) TableName() string {
	return models.ContractPriceTBName()
}

func (m *ContractPrice) Search(ctype, grade int, length float64) (price ContractPrice, err error) {
	err = orm.NewOrm().QueryTable(models.ContractPriceTBName()).Filter("type", ctype).Filter("lower__lte", length).Filter("upper__gt", length).One(&price)
	return
}

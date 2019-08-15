package utils

import (
	"fmt"
	"time"
)

const FormatDateTime = "2006-01-02 03:04:05"
const FormatDate = "2006-01-02"

func MonthRange() (firstDay, lastDay string) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	firstDay = firstOfMonth.Format("2006-01-02")
	lastDay = lastOfMonth.Format("2006-01-02")
	return
}

func WeekRange() (firstDay, lastDay string) {
	now := time.Now()
	//time.Weekday类型可以做运算，强制转int,会得到偏差数。
	//默认是 Sunday 开始到 Saturday 算 0,1,2,3,4,5,6
	offset1 := int(time.Monday - now.Weekday())
	offset2 := 7 - int(now.Weekday()-time.Sunday)
	if offset1 > 0 {
		offset1 = -6
	}
	if offset2 > 6 {
		offset2 = 0
	}
	fmt.Println("offset1%d,off2:%d\n", offset1, offset2)
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset1)
	weekEnd := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset2)
	firstDay = weekStart.Format("2006-01-02")
	lastDay = weekEnd.Format("2006-01-02")
	return
}

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// 获取上个月第一天
func GetPreFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, -1, -d.Day()+1)
	return GetZeroTime(d)
}

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 周日开始 计算下周一的时间
func GetNextNDate(d time.Time, n int) string {
	d = d.AddDate(0, 0, n)
	return d.Format("2006-01-02")
}

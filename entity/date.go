package entity

import (
	"fmt"
	"strconv" // 用于string到基本类型的转化
)

// Date struct
type Date struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
}

func (m *Date) initDate(tyear, tMonth, tDay, tHour, tMinute int) {
	m.Year = tyear
	m.Month = tMonth
	m.Day = tDay
	m.Hour = tHour
	m.Minute = tMinute
}

func (m Date) getYear() int {
	return m.Year
}

func (m *Date) setYear(y int) {
	m.Year = y
}

func (m Date) getMonth() int {
	return m.Month
}

func (m *Date) setMonth(mo int) {
	m.Month = mo
}

func (m Date) getDay() int {
	return m.Day
}

func (m *Date) setDay(d int) {
	m.Day = d
}

func (m Date) getHour() int {
	return m.Hour
}

func (m *Date) setHour(h int) {
	m.Hour = h
}

func (m Date) getMinute() int {
	return m.Minute
}

func (m *Date) setMinute(mi int) {
	m.Minute = mi
}

func (m Date) isValid() bool {
	if m.Year < 1000 || m.Year > 9999 {
		return false
	}
	if m.Month < 1 || m.Month > 12 {
		return false
	}
	if m.Day < 1 || m.Day > 31 {
		return false
	}
	if m.Hour < 0 || m.Hour > 23 {
		return false
	}
	if m.Minute < 0 || m.Minute > 59 {
		return false
	}
	if ((m.Month%2 == 1 && m.Month < 8) || (m.Month%2 == 0 && m.Month >= 8)) && (m.Day > 31 || m.Day < 1) {
		return false
	}
	if (m.Month == 4 || m.Month == 6 || m.Month == 9 || m.Month == 11) && (m.Day > 30 || m.Day < 1) {
		return false
	}
	if ((m.Year%4 == 0 && m.Year%100 != 0) || m.Year%400 == 0) && (m.Day > 29 || m.Day < 1) && m.Month == 2 {
		return false
	}
	if !((m.Year%4 == 0 && m.Year%100 != 0) || m.Year%400 == 0) && (m.Day > 28 || m.Day < 1) && m.Month == 2 {
		return false
	}
	return true
}

func stringToInt(s string) int {
	result, error := strconv.Atoi(s)

	if error != nil {
		fmt.Println("fail")
	}
	return result
}

//"0000-00-00/00:00"
func buildDateFromString(tDateString string) Date {
	var i int
	var x Date
	for i = 0; i < 4; i++ {
		if tDateString[i] > '9' || tDateString[i] < '0' {
			x.initDate(0, 0, 0, 0, 0)
			return x
		}
	}
	if tDateString[4] != '-' || tDateString[7] != '-' || len(tDateString) != 16 {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if tDateString[10] != '/' || tDateString[13] != ':' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if tDateString[14] > '9' || tDateString[14] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if tDateString[15] > '9' || tDateString[15] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if tDateString[5] > '9' || tDateString[5] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if tDateString[6] > '9' || tDateString[6] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if tDateString[8] > '9' || tDateString[8] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if tDateString[9] > '9' || tDateString[9] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if tDateString[11] > '9' || tDateString[11] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if tDateString[12] > '9' || tDateString[12] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	//0000-00-00/00:00
	x.setYear(stringToInt(tDateString[0:4]))
	x.setMonth(stringToInt(tDateString[5:7]))
	x.setDay(stringToInt(tDateString[8:10]))
	x.setHour(stringToInt(tDateString[11:13]))
	x.setMinute(stringToInt(tDateString[14:16]))
	if x.isValid() != false {
		return x
	} else {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}

}

// IntToString convert int to string
func IntToString(a int) string {
	// var res string
	return strconv.Itoa(a)
}

func dateToString(tDate Date) string {
	if tDate.isValid() == false {
		return "0000-00-00/00:00"
	}

	var re string = ""
	re += IntToString(tDate.Year) + "-"

	if tDate.Month < 10 {
		re += "0"
	}
	re += IntToString(tDate.Month) + "-"

	if tDate.Day < 10 {
		re += "0"
	}
	re += IntToString(tDate.Day) + "/"

	if tDate.Hour < 10 {
		re += "0"
	}
	re += IntToString(tDate.Hour) + ":"

	if tDate.Minute < 10 {
		re += "0"
	}
	re += IntToString(tDate.Minute)
	return re
}

func (m *Date) copyDate(t Date) {
	m.Year = t.Year
	m.Month = t.Month
	m.Day = t.Day
	m.Hour = t.Hour
	m.Minute = t.Minute
}

func (m Date) isTheSame(tDate Date) bool {
	return (tDate.getYear() == m.getYear() &&
		tDate.getMonth() == m.getMonth() &&
		tDate.getDay() == m.getDay() &&
		tDate.getHour() == m.getHour() &&
		tDate.getMinute() == m.getMinute())
}

func (m Date) isMoreThan(tDate Date) bool {
	if m.Year < tDate.Year {
		return false
	}
	if m.Year > tDate.Year {
		return true
	}
	if m.Month < tDate.Month {
		return false
	}
	if m.Month > tDate.Month {
		return true
	}
	if m.Day < tDate.Day {
		return false
	}
	if m.Day > tDate.Day {
		return true
	}
	if m.Hour < tDate.Hour {
		return false
	}
	if m.Hour > tDate.Hour {
		return true
	}
	if m.Minute < tDate.Minute {
		return false
	}
	if m.Minute > tDate.Minute {
		return true
	}
	return false
}

func (m Date) isLessThan(tDate Date) bool {
	if m.isMoreThan(tDate) != true && m.isTheSame(tDate) != true {
		return true
	}
	return false
}

// GreaterOrEqual => used to validation check of service layer
func (m Date) GreaterOrEqual(tDate Date) bool {
	if m.isLessThan(tDate) != true {
		return true
	}
	return false
}

// SmallerOrEqual => used to validation check of service layer
func (m Date) SmallerOrEqual(tDate Date) bool {
	if m.isMoreThan(tDate) != true {
		return true
	}
	return false
}

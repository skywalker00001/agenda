package entity

import (
	"fmt"
	"strconv"
)

// Date *
type Date struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
}

// NewDate constuct a date according to the params
// params: an int indicate the year of the date
// params: an int indicate the month of the date
// params: an int indicate the day of the date
// params: an int indicate the hour of the date
// params: an int indicate the minute of the date
func NewDate(year, month, day, hour, minute int) *Date {
	return &Date{
		Year:   year,
		Month:  month,
		Day:    day,
		Hour:   hour,
		Minute: minute,
	}
}

// GetYear *
func (d Date) GetYear() int {
	return d.Year
}

// GetMonth *
func (d Date) GetMonth() int {
	return d.Month
}

// GetDay *
func (d Date) GetDay() int {
	return d.Day
}

// GetHour *
func (d Date) GetHour() int {
	return d.Hour
}

// GetMinute *
func (d Date) GetMinute() int {
	return d.Minute
}

// SetYear *
func (d *Date) SetYear(year int) {
	d.Year = year
}

// SetMonth *
func (d *Date) SetMonth(month int) {
	d.Month = month
}

// SetDay *
func (d *Date) SetDay(day int) {
	d.Day = day
}

// SetHour *
func (d *Date) SetHour(hour int) {
	d.Hour = hour
}

// SetMinute *
func (d *Date) SetMinute(minute int) {
	d.Minute = minute
}

// IsValid *
func (d Date) IsValid() bool {
	// build a table represents for how many days per month
	dayOfMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (d.Year%4 == 0 && d.Year%100 != 0) || (d.Year%400 == 0) {
		dayOfMonth[1] = 29
	}

	if d.Month < 1 || d.Month > 12 {
		return false
	}

	if d.Day < 1 || d.Day > dayOfMonth[d.Month-1] {
		return false
	}

	if d.Hour < 0 || d.Hour > 23 {
		return false
	}

	if d.Minute < 0 || d.Minute > 59 {
		return false
	}

	return true
}

// StringToDate convert a date string to a Date type
// param: a string with format YYYY-MM-DD/HH:mm
// if the string doesn't fit the format,
// return an uncompleted Date struct
func StringToDate(dateString string) Date {
	d := Date{}

	// YYYY-MM-DD/HH:mm
	if len(dateString) != 16 {
		return d
	}
	if dateString[4] != '-' || dateString[7] != '-' || dateString[10] != '/' || dateString[13] != ':' {
		return d
	}
	// convert string year to int
	year, err := strconv.Atoi(dateString[0:4])
	if err != nil {
		return d
	}
	d.Year = year
	// convert string month to int
	month, err := strconv.Atoi(dateString[5:7])
	if err != nil {
		return d
	}
	d.Month = month
	// convert string day to int
	day, err := strconv.Atoi(dateString[8:10])
	if err != nil {
		return d
	}
	d.Day = day
	// convert string hour to int
	hour, err := strconv.Atoi(dateString[11:13])
	if err != nil {
		return d
	}
	d.Hour = hour
	// convert string minute to int
	minute, err := strconv.Atoi(dateString[14:16])
	if err != nil {
		return d
	}
	d.Minute = minute

	return d
}

// DateToString convert a Date struct to a string
// with format YYYY-MM-DD/HH:mm
func DateToString(date Date) string {
	if date.IsValid() {
		return fmt.Sprintf("%04d-%02d-%02d/%02d:%02d", date.Year, date.Month, date.Day, date.Hour, date.Minute)
	}
	return "0000-00-00/00:00"
}

// Assign *
func (d *Date) Assign(date Date) {
	d.Year = date.Year
	d.Month = date.Month
	d.Day = date.Day
	d.Hour = date.Hour
	d.Minute = date.Minute
}

// IsEqual *
func (d Date) IsEqual(date Date) bool {
	if d.Year == date.Year && d.Month == date.Month && d.Day == date.Day && d.Hour == date.Hour && d.Minute == date.Minute {
		return true
	}
	return false
}

// IsGreater *
func (d Date) IsGreater(date Date) bool {
	if d.Year > date.Year {
		return true
	} else if d.Year < date.Year {
		return false
	}
	// now d.Year == date.Year
	if d.Month > date.Month {
		return true
	} else if d.Month < date.Month {
		return false
	}
	// now d.Month == date.Month
	if d.Day > date.Day {
		return true
	} else if d.Day < date.Day {
		return false
	}
	// now d.Day == date.Day
	if d.Hour > date.Hour {
		return true
	} else if d.Hour < date.Hour {
		return false
	}
	// now d.Hour == date.Hour
	if d.Minute > date.Minute {
		return true
	} else if d.Minute < date.Minute {
		return false
	}
	// now d == date
	return false
}

// IsLess *
func (d Date) IsLess(date Date) bool {
	if d.Year < date.Year {
		return true
	} else if d.Year > date.Year {
		return false
	}
	// now d.Year == date.Year
	if d.Month < date.Month {
		return true
	} else if d.Month > date.Month {
		return false
	}
	// now d.Month == date.Month
	if d.Day < date.Day {
		return true
	} else if d.Day > date.Day {
		return false
	}
	// now d.Day == date.Day
	if d.Hour < date.Hour {
		return true
	} else if d.Hour > date.Hour {
		return false
	}
	// now d.Hour == date.Hour
	if d.Minute < date.Minute {
		return true
	} else if d.Minute > date.Minute {
		return false
	}
	// now d == date
	return false
}

// IsGreaterThanEqual *
func (d Date) IsGreaterThanEqual(date Date) bool {
	return !d.IsLess(date)
}

// IsLessThanEqual *
func (d Date) IsLessThanEqual(date Date) bool {
	return !d.IsGreater(date)
}

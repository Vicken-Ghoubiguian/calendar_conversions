package calendar_conversions

//
import (
    "time"
	"fmt"
	"strconv"
	"strings"
)

//
type Gregorian struct {

	//
	year int
	month time.Month
	monthNumber int
	day int
	hour int
	minute int
	second int
	microseconds int
}

//
func (current_gregorian *Gregorian) Initialize_gregorian_from_elements(year int, monthNumber int, day int, hour int, minute int, second int, microseconds int) {

	current_gregorian.year = year
	current_gregorian.month = time.Month(monthNumber)
	current_gregorian.monthNumber = monthNumber
	current_gregorian.day = day
	current_gregorian.hour = hour
	current_gregorian.minute = minute
	current_gregorian.second = second
	current_gregorian.microseconds = microseconds
}

//
func (current_gregorian *Gregorian) Initialize_gregorian_from_time(dt time.Time) {

	//
	current_gregorian.year = dt.Year()
	current_gregorian.month = dt.Month()
	current_gregorian.monthNumber = int(dt.Month())
	current_gregorian.day = dt.Day()
	current_gregorian.hour = dt.Hour()
	current_gregorian.minute = dt.Minute()
	current_gregorian.second = dt.Second()

	//
	pts := dt.String()
	pts_split := strings.Split(pts, ".")
	pts_split_2 := strings.Split(pts_split[1], " ")
	pts_split_3 := pts_split_2[0]

	//
	current_gregorian.microseconds, _ = strconv.Atoi(pts_split_3)
}

//
func(current_gregorian *Gregorian) Initialize_time(year int, monthNumber int, day int, hour int, minute int, second int, microseconds int) time.Time {

	return time.Date(year, time.Month(monthNumber), day, hour, minute, second, microseconds, time.Now().Location())
}

//
func (current_gregorian *Gregorian) Get_year() int {

	return current_gregorian.year
}

//
func (current_gregorian *Gregorian) Get_month() time.Month {

	return current_gregorian.month
}

//
func (current_gregorian *Gregorian) Get_monthNumber() int {

	return current_gregorian.monthNumber
}

//
func (current_gregorian *Gregorian) Get_day() int {

	return current_gregorian.day
}

//
func (current_gregorian *Gregorian) Get_hour() int {

	return current_gregorian.hour
}

//
func (current_gregorian *Gregorian) Get_minute() int {

	return current_gregorian.minute
}

//
func (current_gregorian *Gregorian) Get_second() int {

	return current_gregorian.second
}

//
func (current_gregorian *Gregorian) Get_microseconds() int {

	return current_gregorian.microseconds
}

//
func (current_gregorian *Gregorian) Format() string {

	year := strconv.Itoa(current_gregorian.year)

	month := strconv.Itoa(int(current_gregorian.month))
	if int(current_gregorian.month) < 10 {

		month = "0" + month
	}

	day := strconv.Itoa(current_gregorian.day)
	if current_gregorian.day < 10 {

		day = "0" + day
	}

	hour := strconv.Itoa(current_gregorian.hour)
	if current_gregorian.hour < 10 {

		hour = "0" + hour
	}

	minute := strconv.Itoa(current_gregorian.minute)
	if current_gregorian.minute < 10 {

		minute = "0" + minute
	}

	second := strconv.Itoa(current_gregorian.second)
	if current_gregorian.second < 10 {

		second = "0" + second
	}

	microseconds := strconv.Itoa(current_gregorian.microseconds)

	return fmt.Sprintf("%s-%s-%s %s:%s:%s.%s", year, month, day, hour, minute, second, microseconds)
}

//
func (current_gregorian *Gregorian) EqualTo(second_current_gregorian *Gregorian) bool {

    isEqual := false

	if (current_gregorian.year == second_current_gregorian.year) && (current_gregorian.month == second_current_gregorian.month) && (current_gregorian.monthNumber == second_current_gregorian.monthNumber) && (current_gregorian.day == second_current_gregorian.day) && (current_gregorian.hour == second_current_gregorian.hour) && (current_gregorian.minute == second_current_gregorian.minute) && (current_gregorian.second == second_current_gregorian.second) && (current_gregorian.microseconds == second_current_gregorian.microseconds) {

		isEqual = true
	}

    return isEqual
}
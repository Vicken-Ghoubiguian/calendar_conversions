package calendar_conversions

//
import (
	"fmt"
	"time"
	"math"
)

//
type JulianDay struct {

	//
	countOfDaysSinceJulianPeriod float64
}

// (Algorithm from Jean Meeus' 'Astronomical Algorithms', second edition, chapter 7 called 'Julian Day', pages 60 - 61)...
func (current_julian_day *JulianDay) Initialize_julian_day_from_gregorian_date(gregorian_date Gregorian) {

	//
	Y := float64(gregorian_date.Get_year())
	M := float64(gregorian_date.Get_monthNumber())
	D := float64(gregorian_date.Get_day())

	//
	A := math.Floor(Y/100)
	B := math.Floor(A/4)
	C := math.Floor(2 - A + B)
	E := math.Floor(365.25 * (Y + 4716))
	F := math.Floor(30.6001 * (M + 1))

	//
	H := float64(gregorian_date.Get_hour()) / float64(24.0)
    MIN := roundFloat((float64(gregorian_date.Get_minute()) / float64(24.0 * 60)), 5)
    SEC := roundFloat((float64(gregorian_date.Get_second()) / (24.0 * 60 * 60)), 5)
    MILLISEC := float64(gregorian_date.Get_microseconds() / (24.0 * 60 * 60 * 1000))

	//
	current_julian_day.countOfDaysSinceJulianPeriod = (E + F + D + C - 1524.5) + H + MIN + SEC + MILLISEC
}

// (Algorithm from Jean Meeus' 'Astronomical Algorithms', second edition, chapter 7 called 'Julian Day', pages 60 - 61)...
func (current_julian_day *JulianDay) Initialize_julian_day_from_julian_date(julian_date Julian) {

	//
	Y := float64(julian_date.Get_year())
	M := float64(julian_date.Get_monthNumber())
	D := float64(julian_date.Get_day())

	//
	E := math.Floor(365.25 * (Y + 4716))
	F := math.Floor(30.6001 * (M + 1))

	//
	H := float64(julian_date.Get_hour()) / float64(24.0)
	MIN := roundFloat((float64(julian_date.Get_minute()) / float64(24.0 * 60)), 5)
	SEC := roundFloat((float64(julian_date.Get_second()) / (24.0 * 60 * 60)), 5)
	MILLISEC := float64(julian_date.Get_microseconds() / (24.0 * 60 * 60 * 1000))

	//
	current_julian_day.countOfDaysSinceJulianPeriod = (E + F + D - 1524.5) + H + MIN + SEC + MILLISEC
}

//
func (current_julian_day *JulianDay) Determine_gregorian_date() Gregorian {

	//
	var gd Gregorian

	//
	/*Q := current_julian_day.countOfDaysSinceJulianPeriod + 0.5
	new_Q := int64(Q)
	Z, _ := math.Modf(Q)
	new_Z := int64(Z)
	W := (Z - 1867216.25)/36524.25
	X := W/4
	A := Z + 1 + W - X
  	B := A + 1524
	C := (B - 122.1)/365.25
	D := 365.25 * C
	E := (B - D)/30.6001
	F := 30.6001 * E*/

	//
	fmt.Print("")

	//
	//correspondingGregorianDate.day = B - D - F

	//
	return gd
}

//
func (current_julian_day *JulianDay) Determine_julian_date() Julian {

	//
	var jd Julian

	a := math.Floor(current_julian_day.Get_count_of_days_since_julian_period())
	b := a + 1524
	c := math.Floor((b - 122.1) / 365.25)
	d := math.Floor(365.25 * c)
	e := math.Floor((b - d) / 30.6001)
	
	//
	month := -1

	//
	if e < 14 {

		month = int(math.Floor(e - 1))

	} else {

		month = int(math.Floor(e - 13))
	}

	//
	year := -1

	//
	if month > 2 {

		year = int(math.Floor(c - 4716))

	} else {

		year = int(math.Floor(c - 4715))
	}

	//
	day := b - d - math.Floor(30.6001 * e)
	
	//
	jd.day = int(day)
	jd.month = time.Month(month)
	jd.monthNumber = int(month)
	jd.year = int(year)

	//
	jd.hour = 0
	jd.minute = 0
	jd.second = 0
	jd.microseconds = 0

	//
	return jd
}

//
func (current_julian_day *JulianDay) Get_count_of_days_since_julian_period() float64 {

	return current_julian_day.countOfDaysSinceJulianPeriod
}

//
func (current_julian_day *JulianDay) Add_value_to_count_of_days_since_julian_period(adding_value float64) {

	current_julian_day.countOfDaysSinceJulianPeriod += adding_value
}

//
func (current_julian_day *JulianDay) Format() string {

	return fmt.Sprintf("%f", current_julian_day.countOfDaysSinceJulianPeriod)
}

//
func (current_julian_day *JulianDay) EqualTo(second_current_julian_day *JulianDay) bool {

    isEqual := false

    if current_julian_day.countOfDaysSinceJulianPeriod == second_current_julian_day.countOfDaysSinceJulianPeriod {

        isEqual = true
    }

    return isEqual
}
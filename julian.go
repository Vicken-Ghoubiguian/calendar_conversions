package calendarConversionsGo

//
import (
    "time"
	//"math"
)

//
type Julian struct {

	year int64
	month time.Month
	day int64
}

//
func (current_julian *Julian) Initialize_julian(dt time.Time) {

	//
	var jd JulianDay

	//
	jd.Initialize_julian_day_from_time(dt)

	//
	/*jd.Add_value_to_count_of_days_since_julian_period(0.5)
	a := math.Floor(jd.Get_count_of_days_since_julian_period())
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
	current_julian.day = day
	current_julian.year = year*/
}

//
func (current_julian *Julian) Gregorian_to_julian(gregorianDate Gregorian) {


}

//
func (current_julian *Julian) Julian_to_Gregorian() {


}

//
func (current_julian *Julian) Format() string {

	return ""
}
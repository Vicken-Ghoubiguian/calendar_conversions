package calendarConversionsGo

//
import (
	"time"
	"strconv"
)

//
type JulianDay struct {

	countOfDaysSinceJulianPeriod int64
}

//
func (current_julian_day *JulianDay) Initialize_julian_day_from_time(dt time.Time) {

	/*Y := int64(dt.Year())
	M := 0
	D := int64(dt.Day())*/
}

//
func (current_julian_day *JulianDay) Initialize_julian_day_from_gregorian_date(gregorian_date *Gregorian) {


}

//
/*func (current_julian_day *JulianDay) Determine_gregorian_date() Gregorian {


}*/

//
func (current_julian_day *JulianDay) Get_count_of_days_since_julian_period(dt time.Time) int64 {

	return current_julian_day.countOfDaysSinceJulianPeriod
}

//
func (current_julian_day *JulianDay) Format() string {

	return strconv.Itoa(int(current_julian_day.countOfDaysSinceJulianPeriod))
}
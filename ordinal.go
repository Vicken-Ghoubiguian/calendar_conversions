package calendarConversionsGo

//
import (
	"fmt"
	"time"
)

//
type Ordinal_date struct {

	//
	year            int
	day_of_the_year int
}

//
func (current_ordinal_date *Ordinal_date) Initialize_ISO_week_based_calendar_date(dt time.Time) {

}

//
func (current_ordinal_date *Ordinal_date) ToString() string {

	fmt.Print("Hello, World !")
	return ""
}

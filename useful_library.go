package calendar_conversions

//
import (
	"math"
)

//
func roundFloat(val float64, precision uint) float64 {

	//
	ratio := math.Pow(10, float64(precision))

	//
	return math.Trunc(val*ratio) / ratio
}
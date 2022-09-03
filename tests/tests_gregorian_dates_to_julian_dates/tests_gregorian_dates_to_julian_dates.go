package main

import (
	"time"
	"fmt"
	"github.com/Vicken-Ghoubiguian/calendar_conversions"
)

func main() {

	//
	fmt.Println("===============================")

	//
	var november_7_1917_gregorian_to_julian calendar_conversions.Julian
	var november_7_1917_from_time_golang_to_julian calendar_conversions.Julian
	var november_7_1917_in_gregorian calendar_conversions.Gregorian

	//
	november_7_1917_in_gregorian.Initialize_gregorian_from_elements(1917, 11, 7)
	november_7_1917_in_time_golang := time.Date(1917, time.Month(11), 7, 0, 0, 0, 0, time.UTC)

	//
	november_7_1917_gregorian_to_julian.Gregorian_to_julian(november_7_1917_in_gregorian)
	november_7_1917_from_time_golang_to_julian.Initialize_julian(november_7_1917_in_time_golang)

	//
	fmt.Println(november_7_1917_from_time_golang_to_julian.Format())

	//
	fmt.Println("===============================")
}
package main

import (
	"encaps/example3/calendar"
	"fmt"
)

func main() {

	date := calendar.Date{}
	date.SetDay(28)
	date.SetMonth(05)
	date.SetYear(1994)
	fmt.Println(date)
}

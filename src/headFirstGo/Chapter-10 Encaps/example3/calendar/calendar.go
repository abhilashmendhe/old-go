package calendar

import "errors"

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) SetDay(day int) error {

	if day > 31 || day < 1 {
		return errors.New("Can't set the day!!")
	}
	d.day = day
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month > 12 || month < 1 {
		return errors.New("Can't set the month!!")
	}
	d.month = month
	return nil
}

func (d *Date) SetYear(year int) error {

	if year < 0 {
		errors.New("Can't set year!!")
	}
	d.year = year
	return nil
}

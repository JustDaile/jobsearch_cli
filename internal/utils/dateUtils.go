package utils

import (
	"fmt"
	"log"
	"time"
)

func CheckDateWithinDays(format, date string, days int) (bool, error) {
	t, e := time.Parse(format, date)
	if e != nil {
		return false, fmt.Errorf("unable to parse date %v as format %v - error: %v", date, format, e)
	}
	hoursSince := time.Since(t).Hours()
	log.Printf("%v days since checked date", int(hoursSince/24))
	return int(hoursSince/24) <= days, nil
}

func TimeMustParse(format, date string) time.Time {
	t, e := time.Parse(format, date)
	if e != nil {
		panic(e)
	}
	return t
}

package utils

import (
	"testing"
	"time"
)

type checkDateTest struct {
	Time      time.Time
	DaysSince int
	Expect    bool
}

var checkDateTests = []checkDateTest{
	{
		Time:      time.Now(),
		DaysSince: 0,
		Expect:    true,
	},
	{
		Time:      time.Now().AddDate(0, 0, -8),
		DaysSince: 7,
		Expect:    false,
	},
	{
		Time:      time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-5, 0, 0, 0, 0, time.Now().Location()),
		DaysSince: 5,
		Expect:    true,
	},
}

func TestDateUtils(t *testing.T) {
	stdFormat := "02/01/2006"

	for _, test := range checkDateTests {
		formatted := test.Time.Format(stdFormat)
		b, e := CheckDateWithinDays(stdFormat, formatted, test.DaysSince)
		if e != nil {
			t.Log(e)
			t.FailNow()
		}
		if b != test.Expect {
			t.Logf("expected %v when testing that %s is within range of %v days but was %v", test.Expect, formatted, test.DaysSince, b)
			t.Fail()
		} else {
			t.Logf("test success %s within range of %v days - %v", formatted, test.DaysSince, b)
		}
	}
}

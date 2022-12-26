package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	present := time.Now()
	p(present)

	DOB := time.Date(1993, 02, 28, 9, 04, 39, 213, time.Local)
	p(DOB)

	p("Year:", DOB.Year())
	p("Month:", DOB.Month())
	p("Day:", DOB.Day())
	p("Hour", DOB.Hour())
	p("Minute:", DOB.Minute())
	p("Second:", DOB.Second())
	p("Nanosecond:", DOB.Nanosecond())
	p("Location:", DOB.Location())
	p("Weekday:", DOB.Weekday())
	p("Before:", DOB.Before(present))
	p("After:", DOB.After(present))
	p("Equal:", DOB.Equal(present))

	diff := present.Sub(DOB)
	p("Diff:", diff)
	p("Diff hours:", diff.Hours())
	p("Diff minutes:", diff.Minutes())
	p("Diff seconds:", diff.Seconds())
	p("Diff nanoseconds:", diff.Nanoseconds())
	p("Add diff:", DOB.Add(diff))
	p("Add negative diff:", DOB.Add(-diff))

	p("Today:", present.Format("Mon, Jan 2, 2006 at 3:04pm"))
	someTime := time.Date(2017, time.March, 30, 11, 30, 55, 123456, time.Local)
	sameTime := someTime.Equal(present)
	p("someTime equals to now?", sameTime)
	diff2 := present.Sub(someTime)
	days := int(diff2.Hours() / 24)
	p("Diff:", diff)
	fmt.Printf("30th March 2017 was %d days ago\n", days)
}

package main

import "fmt"

type duration int64

const (
	nanosecond  duration = 1
	microsecond          = 1000 * nanosecond
	millisecond          = 1000 * microsecond
	second               = 1000 * millisecond
	minute               = 60 * second
	hour                 = 60 * minute
)

func (d *duration) setHours(h float64) {
	*d = duration(h) * hour
}

func (d duration) hours() float64 {
	hour := d / hour
	nsec := d % hour
	return float64(hour) + float64(nsec)*(1e-9/60/60)
}

func main() {
	fmt.Println(nanosecond, microsecond, millisecond, second, minute, hour)

	// Declare a variable of type duration set to
	// its zero value
	var dur duration

	// Change the value of dur to equal
	// five hours
	dur.setHours(5.2)

	// Display the new vlaue of dur.
	fmt.Println("Hours:", dur.hours(), dur)
}

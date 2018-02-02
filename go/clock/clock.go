package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func New(hour int, minute int) Clock {
	realMin := minute % 60
	if realMin < 0 {
		realMin = 60 + realMin
		hour--
	}
	realHour := (hour + (minute / 60)) % 24
	if realHour < 0 {
		realHour = 24 + realHour
	}
	return Clock{hour: realHour, minute: realMin}
}

func (c *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(min int) Clock {
	return New(c.hour, c.minute+min)
}

package clock

import (
	"fmt"
)

// Clock defines a clock.
type Clock struct {
	hour   int
	minute int
}

// New creates a new Clock.
func New(hour, minute int) Clock {
	if minute > 59 {
		for {
			if minute > 59 {
				minute -= 60
				hour++
			} else {
				break
			}
		}
	}

	if minute < 0 {
		for {
			if minute < 0 {
				minute += 60
				hour--
			} else {
				break
			}
		}
	}

	if hour >= 24 {
		for {
			if hour >= 24 {
				hour -= 24
			} else {
				break
			}
		}
	}

	if hour < 0 {
		for {
			if hour < 0 {
				hour += 24
			} else {
				break
			}
		}
	}

	return Clock{hour: hour, minute: minute}
}

// String returns the string representation of the Clock's time.
func (c Clock) String() string {
	if c.hour == 24 && c.minute == 00 {
		return fmt.Sprintf("00:00")
	}
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds minutes to the time.
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract subtracts minutes from the time.
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute+(minutes*-1))
}

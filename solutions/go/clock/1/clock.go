package clock

import "fmt"

type Clock int

const minutesInDay = 24 * 60

func normalize(m int) int {
	m = m % minutesInDay
	if m < 0 {
		m += minutesInDay
	}
	return m
}

func New(h, m int) Clock {
	total := h*60 + m
	return Clock(normalize(total))
}

func (c Clock) Add(m int) Clock {
	return Clock(normalize(int(c) + m))
}

func (c Clock) Subtract(m int) Clock {
	return Clock(normalize(int(c) - m))
}

func (c Clock) String() string {
	h := int(c) / 60
	m := int(c) % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}
package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    
    layout := "1/2/2006 15:04:05"

    t, _ := time.Parse(layout,date) // time.Time, error
    return t
}

func HasPassed(date string) bool {

    layout := "January 2, 2006 15:04:05"

    t, _ := time.Parse(layout, date)

    // check if appointment is before current time
    return t.Before(time.Now())
}


func IsAfternoonAppointment(date string) bool {

    layout := "Monday, January 2, 2006 15:04:05"

    t, _ := time.Parse(layout, date)

    hour := t.Hour()

    return hour >= 12 && hour < 18
}

func Description(date string) string {

    // input format
    inputLayout := "1/2/2006 15:04:05"

    t, _ := time.Parse(inputLayout, date)

    // required output format
    output := t.Format("Monday, January 2, 2006, at 15:04")

    return "You have an appointment on " + output + "."
}

func AnniversaryDate() time.Time {

    currentYear := time.Now().Year()

    return time.Date(
        currentYear,
        time.September,
        15,
        0,
        0,
        0,
        0,
        time.UTC,
    )
}

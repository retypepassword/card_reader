package main

import (
    "fmt"
    "time"
    "os"
)

func main() {
    for {
        var today time.Time = time.Now()
        var last_swipe time.Time = today
        var filename string = fmt.Sprintf("bis2a_oh_%s%v.csv", get_quarter(today), today.Year())

        f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            panic(err)
        }

        // Make a new file every quarter. Update last_swipe and today after the
        // loop so that the first swipe of this month doesn't get recorded in
        // last month's file.
        for ; get_quarter(today) == get_quarter(last_swipe); last_swipe, today = today, time.Now() {
            var today_str string = today.Format("01,02,2006,15:04")
            var student_id string
            fmt.Scan(&student_id)
            _, err := f.WriteString(fmt.Sprintf("%s,%s\n", today_str, student_id))
            if err != nil {
                panic(err)
            }
        }

        f.Close()
    }
}

// Given a date, returns the quarter of the year for the date as W, S, U, or F
// (winter, spring, summer, fall)
func get_quarter(date time.Time) string {
    switch date.Month() {
        case 1, 2, 3:
            return "W"
        case 4, 5, 6:
            return "S"
        case 7, 8:
            return "U"
        case 9, 10, 11, 12:
            return "F"
    }
    return ""
}

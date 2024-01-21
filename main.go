package main

import (
	"fmt"
	"time"
)

func main() {
    // Get the current time
    now := time.Now()

    // Get the current month
    month := now.Month()

    // Get the current year
    year := now.Year()

    // Get the day of the month
    firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

    // Get the last day of the month
    lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

    // Get the day of the week for the first day of the month
    firstDay := firstOfMonth.Weekday()

    // Get the number of days in the month
    daysInMonth := lastOfMonth.Day()
    
    // Print the calendar header
    fmt.Printf("Calendar for %s %d\n", month, year)

    // Print the days of the week
    fmt.Println("Su Mo Tu We Th Fr Sa")

    // Print the first week
    for i := 0; i < int(firstDay); i++ {
        fmt.Print("   ")
    }
    for i := 1; i <= 7-int(firstDay); i++ {
        fmt.Printf("%2d ", i)
    }
    fmt.Println()

    // Print the remaining weeks
    for i:=8-int(firstDay); i <= daysInMonth; i += 7 {
        for j := i; j < i+7; j++ {
            if j > daysInMonth {
                break
            }
            fmt.Printf("%2d ", j)
        }
        fmt.Println()
    }
}

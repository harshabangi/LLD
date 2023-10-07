package main

import (
	"fmt"
	"log"
	"time"
)

func NewDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func main() {
	m := map[time.Time][]string{}

	n1, err := time.Parse("03:04 PM", "12:00 PM")
	if err != nil {
		log.Fatal(err)
	}

	n2, err := time.Parse("03:04 PM", "01:00 PM")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n2.Sub(n1).Seconds())

	m[NewDate(2023, 2, 1)] = append(m[NewDate(2023, 2, 1)], "x")
	m[NewDate(2023, 2, 1)] = append(m[NewDate(2023, 2, 1)], "y")
	m[NewDate(2023, 2, 2)] = append(m[NewDate(2023, 2, 2)], "z")

	l := time.Time{}

	fmt.Println(l.IsZero())

	fmt.Println(m)

}

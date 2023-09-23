package main

import (
	"fmt"
	"github.com/harshabangi/LLD/apps/calculator/internal/service"
)

/*
Assumptions
1. Input is always valid
*/

func main() {
	s := service.NewCalculatorService()
	o1 := s.Calculate("2+1*2", "1")
	o2 := s.Calculate("2*1*2", "2")

	fmt.Printf("%+v", s.ListCalculations())
	fmt.Println(s.GetCalculationRecord("2"))

	fmt.Println(o1)
	fmt.Println(o2)
}

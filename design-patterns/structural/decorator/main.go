package main

import (
	"fmt"
)

/*
Decorator pattern is a structural design pattern that allows adding
new behaviours to objects dynamically by wrapping them in a special
wrapper objects called decorators.
*/

type Pizza interface {
	GetPrice() int64
}

type VegMania struct {
}

func (v *VegMania) GetPrice() int64 {
	return 2
}

type CheeseTopping struct {
	pizza Pizza
}

func (c *CheeseTopping) GetPrice() int64 {
	return 1 + c.pizza.GetPrice()
}

type TomatoTopping struct {
	pizza Pizza
}

func (t *TomatoTopping) GetPrice() int64 {
	return 1 + t.pizza.GetPrice()
}

func main() {
	v := &VegMania{}
	fmt.Printf("veg mania price: %d\n", v.GetPrice())

	vegManiaWithCheeseAndTomatoTopping := &TomatoTopping{&CheeseTopping{v}}

	fmt.Printf("veg mania with cheese and tomatao topping: %d\n", vegManiaWithCheeseAndTomatoTopping.GetPrice())
}

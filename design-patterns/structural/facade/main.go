package main

import (
	"fmt"
)

/*
The Facade pattern is used to provide a simplified and unified interface to a set of interfaces or subsystems.
A facade is an object that serves as a front-facing interface masking more complex underlying or structural code
*/

// SubsystemA represents a complex subsystem.
type SubsystemA struct{}

func (a *SubsystemA) OperationA() {
	fmt.Println("SubsystemA: Performing Operation A")
}

// SubsystemB represents another complex subsystem.
type SubsystemB struct{}

func (b *SubsystemB) OperationB() {
	fmt.Println("SubsystemB: Performing Operation B")
}

// Facade provides a simplified interface to the complex subsystems.
type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
	}
}

func (f *Facade) DoSomething() {
	fmt.Println("Facade: Starting the process")
	f.subsystemA.OperationA()
	f.subsystemB.OperationB()
	fmt.Println("Facade: Process completed")
}

func main() {
	facade := NewFacade()
	facade.DoSomething()
}

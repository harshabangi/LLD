package main

import (
	"container/heap"
	"fmt"
)

type Direction string

const (
	Up   Direction = "up"
	Down Direction = "down"
	Idle Direction = "idle"
)

type Location string

const (
	InsideElevator  Location = "inside_elevator"
	OutsideElevator Location = "outside_elevator"
)

type Requests []Request

func (u *Requests) Push(x any) {
	*u = append(*u, x.(Request))
}

func (u *Requests) Pop() any {
	a := *u
	n := len(*u)
	*u = a[:n-1]
	return a[n-1]
}

func (u *Requests) Len() int {
	return len(*u)
}

func (u *Requests) Less(i, j int) bool {
	switch (*u)[i].Direction {
	case Up:
		return (*u)[i].ToFloor < (*u)[j].ToFloor
	case Down:
		return (*u)[i].ToFloor > (*u)[j].ToFloor
	default:
		return false
	}
}

func (u *Requests) Swap(i, j int) {
	(*u)[i], (*u)[j] = (*u)[j], (*u)[i]
}

type Request struct {
	FromFloor int
	ToFloor   int
	Direction Direction
	Location  Location
}

type Elevator struct {
	CurrentFloor int
	Direction    Direction
	UpRequests   Requests
	DownRequests Requests
}

func NewElevator() *Elevator {
	ur := Requests{}
	dr := Requests{}

	heap.Init(&ur)
	heap.Init(&dr)

	return &Elevator{
		Direction:    Idle,
		UpRequests:   ur,
		DownRequests: dr,
	}
}

func (e *Elevator) WithCurrentFloor(floor int) *Elevator {
	e.CurrentFloor = floor
	return e
}

func (e *Elevator) Run() {
	for e.UpRequests.Len() > 0 || e.DownRequests.Len() > 0 {
		e.ProcessRequests()
	}
	e.Direction = Idle
	fmt.Println("Finished processing all the requests")
}

func (e *Elevator) ProcessRequests() {
	if e.Direction == Up || e.Direction == Idle {
		e.ProcessUpRequests()
		e.ProcessDownRequests()
	} else {
		e.ProcessDownRequests()
		e.ProcessUpRequests()
	}
}

func (e *Elevator) ProcessUpRequests() {
	for e.UpRequests.Len() > 0 {
		top := heap.Pop(&e.UpRequests)
		t := top.(Request)
		e.CurrentFloor = t.ToFloor
	}

	if e.DownRequests.Len() > 0 {
		e.Direction = Down
	} else {
		e.Direction = Idle
	}
}

func (e *Elevator) ProcessDownRequests() {
	for e.DownRequests.Len() > 0 {
		top := heap.Pop(&e.DownRequests)
		t := top.(Request)
		e.CurrentFloor = t.ToFloor
	}

	if e.UpRequests.Len() > 0 {
		e.Direction = Up
	} else {
		e.Direction = Idle
	}
}

func (e *Elevator) sendRequest(req Request) {

	if req.Direction == Up {
		heap.Push(&e.UpRequests, req)
		heap.Push(&e.UpRequests, Request{
			FromFloor: 0,
			ToFloor:   0,
			Direction: "",
			Location:  "",
		})
	}

	heap.Push(&e.UpRequests, req)
}

func main() {
	elevator := NewElevator().WithCurrentFloor(0)

	elevator.sendRequest(Request{
		FromFloor: 0,
		ToFloor:   2,
		Direction: Up,
	})

	elevator.sendRequest(Request{
		FromFloor: 3,
		ToFloor:   4,
		Direction: Up,
	})

	elevator.sendRequest(Request{
		FromFloor: 4,
		ToFloor:   1,
		Direction: Down,
	})

}

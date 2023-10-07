package main

import (
	"sync"
)

/*
The singleton pattern is a design pattern used in software engineering to ensure
that a class has only one instance and provides a global point of access to that instance
*/

var (
	mutex sync.Mutex
	once  sync.Once
)

type single struct{}

var singleInstance *single

// getSingleInstance1 using mutex
func getSingleInstance1() *single {
	if singleInstance == nil {
		mutex.Lock()
		singleInstance = &single{}
		mutex.Unlock()
	}
	return singleInstance
}

// getSingleInstance1 using once
func getSingleInstance2() *single {
	once.Do(func() {
		singleInstance = &single{}
	})
	return singleInstance
}

func main() {
	for i := 0; i < 100; i++ {
		go getSingleInstance1()
	}
}

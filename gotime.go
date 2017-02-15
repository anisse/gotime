package main

import (
	"fmt"
	"time"
)

func doSomething() {
	time.Sleep(10 * time.Millisecond) //OMIT
}

func example2() { //OMIT
	t1 := time.Now()
	doSomething()
	// Time jumps backwards (e.g after leap second adjustment)
	//...
	t2 := time.Now()
	fmt.Println("Elapsed", t2.Sub(t1))
}

func timeNow() { //OMIT
	fmt.Println("Go time", time.Now())
}

func main() {
	//example1 OMIT
	t1 := time.Now()
	doSomething()
	//...
	t2 := time.Now()
	fmt.Println("Elapsed", t2.Sub(t1))
}

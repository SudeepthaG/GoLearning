package errorhandling

import (
	"fmt"
	"runtime/debug"
)

// we use panic to prematurely terminate the program. When a function encounters a panic, its execution is stopped, any deferred functions are executed and then the control returns to its caller.
// This process continues until all the functions of the current goroutine have returned at which point the program prints the panic message, followed by the stack trace and then terminates.
// signature:func panic(interface{})

// It is possible to regain control of a panicking program using recover which is called from defer functions
// signTURE: func recover() interface{}

func PanicRecoveryPractice() {
	// panicExample()
	// panicWithDeferExample()
	// panicWithRecoverExample()
	// recoverFromInvalidIndex()
	// recoverWithStackTrace()
	recoverWithDifferentGoRoutine()

}

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func panicExample() {
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

func fullName2(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func panicWithDeferExample() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName2(&firstName, nil)
	fmt.Println("returned normally from main")
}

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName3(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func panicWithRecoverExample() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName3(&firstName, nil)
	fmt.Println("returned normally from main")
}

func recoverInvalidAccess() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
	}
}

func invalidSliceAccess() {
	defer recoverInvalidAccess()
	n := []int{5, 7, 4}
	fmt.Println(n[4])
	fmt.Println("normally returned from a")
}

// panic is auto caused by accessing an invalid index of a slice.
func recoverFromInvalidIndex() {
	invalidSliceAccess()
	fmt.Println("normally returned from main")

}

func recoverFullName2() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
		debug.PrintStack()
	}
}

func fullName4(firstName *string, lastName *string) {
	defer recoverFullName2()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

// If we recover from a panic, we lose the stack trace about the panic.
// We can print the stack trace using the PrintStack function of the Debug package
func recoverWithStackTrace() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName4(&firstName, nil)
	fmt.Println("returned normally from main")
}

// Recover works only when it is called from the same goroutine which is panicking. It's not possible to recover from a panic that has happened in a different goroutine
func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func sum(a int, b int) {
	defer recovery()
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	done := make(chan bool)
	go divide(a, b, done)
	<-done
}

func divide(a int, b int, done chan bool) {
	fmt.Printf("%d / %d = %d", a, b, a/b)
	done <- true

}

func recoverWithDifferentGoRoutine() {
	sum(5, 0)
	fmt.Println("normally returned from main")
}

package pointersstructsmethods

import "fmt"

func MethodsPractice() string {

	methodsTypes()

	return "\n\nend of methods topic\n"
}

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

type Employee struct {
	name string
	age  int
}

type Emp struct {
	name string
}

/*
Method with value receiver
*/

func (e Emp) changeName(newName string) { //we canoverload 'cos we have different struct types for calling
	e.name = newName
}

func (e Employee) changeName(newName string) {
	e.name = newName
}

func changeName(e Employee) {
	e.name = "dfdfd"
}

/*
Method with pointer receiver
*/
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func changeAge(e *Employee) {
	e.age = 23
}

func methodsTypes() {
	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	// methods on types are a way to achieve behaviour similar to classes. Methods allow a logical grouping of behaviour related to a type similar to classes.
	emp_e := Emp{
		name: "Andrew",
	}
	fmt.Printf("Employee name before change: %s", emp_e)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", emp_e)

	// Generally, pointer receivers can be used when changes made to the receiver inside the method should be visible to the caller.
	// Pointers receivers can also be used in places where it's expensive to copy a data structure.  if a pointer receiver is used,
	// the struct will not be copied and only a pointer to it will be used
	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	e.changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)

	fmt.Printf("\n\nEmployee name with value call before change: %s", e.name)
	changeName(e) //this is a function not a method. by using  methods, you are compartamentalizing the methods to specific data types only.
	fmt.Printf("\nEmployee age after change: %s", e.name)

	// Value receivers in methods vs Value arguments in functions
	// When a function has a value argument, it will accept only a value argument.When a method has a value receiver, it will accept both pointer and value receivers.
	p := &e
	p.changeName("avfcvd")
	e.changeName("dfgt")
	changeName(e)
	// changeName(p)	//error

	// Pointer receivers in methods vs Pointer arguments in functions-Similar to value arguments
	p.changeAge(22)
	e.changeAge(25)
	changeAge(p)
	// changeAge(e)   //error

	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2) //To define a method on a type, the definition of the receiver type and the definition of the method should be present in the same package.
	fmt.Println("\nSum is", sum)

}

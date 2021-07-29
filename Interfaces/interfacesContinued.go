package interfaces

import "fmt"

func InterFacesContinuePractice() {
	// pointerAndValueReceiversExample()
	// implementMultipleInterfacesExample()
	// embeddInterfacesExample()
	zeroValueOfInterface()

}

type Describer3 interface {
	Describe3()
}
type Person3 struct {
	name string
	age  int
}

func (p Person3) Describe3() { //implemented using value receiver-=value receiver can be called using both value and address
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe3() { //implemented using pointer receiver
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func pointerAndValueReceiversExample() {
	var d1 Describer3
	p1 := Person3{"Sam", 35}
	d1 = p1
	d1.Describe3()
	p3 := Person3{"James", 33}
	d1 = &p3
	d1.Describe3()

	var d3 Describer3
	a := Address{"Washington", "USA"}
	// Compilation error if the following line is uncommented cannot use a (type Address) as type Describer in assignment: Address does not implement Describer
	// That is Describe method has pointer receiver. It is similar to using int variable tp access Describe while Describe has no implemetation with int type
	//d3 = a
	// temp:=5
	// temp.Describe()
	d3 = &a //This works since Describer interface is implemented by Address pointer
	d3.Describe3()
	(&a).Describe3()
}

type SalaryCalculator2 interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

type Emp struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func implementMultipleInterfacesExample() {
	e := Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var s SalaryCalculator2 = e
	s.DisplaySalary()
	var l LeaveCalculator = e
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
	e.DisplaySalary()
	fmt.Println("\nLeaves left =", e.CalculateLeavesLeft())
}

//we can provide inheritance in go using embedded intercaes
type EmployeeOperations interface {
	SalaryCalculator2
	LeaveCalculator
}

func (e Emp) CalculateLeavesLeft() int { //this causes error if we use this along with (e Employee) DisplaySalary() in empOpn embedded interface
	fmt.Println("emp type")
	return e.totalLeaves - e.leavesTaken
}

func (e Emp) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func embeddInterfacesExample() {
	e := Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())

	e2 := Emp{
		firstName:   "Naveen2",
		lastName:    "Ramanathan",
		basicPay:    50000,
		pf:          2000,
		totalLeaves: 300,
		leavesTaken: 50,
	}
	var empOp2 EmployeeOperations = e2
	empOp2.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp2.CalculateLeavesLeft())
}

func zeroValueOfInterface() {
	var d1 Describer
	if d1 == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
	}
	// d1.Describe() //causes panic
}

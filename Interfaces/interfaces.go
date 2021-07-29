package interfaces

import "fmt"

func InterfacesPractice() {
	// interfaceExample()
	// practicalExample()
	// interfaceInternalRepresentationExample()
	// emptyInterFaceExample()
	// typeAssertionExample()
	// typeSwitchExample()
	CompareInterfaceTypeExample()
}

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func interfaceExample() {
	name := MyString("sam anderson")
	var v VowelsFinder
	v = name
	fmt.Printf("Vowels are %c", v.FindVowels())
}

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func practicalExample() {
	pemp1 := Permanent{
		empId:    1,
		basicpay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empId:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empId:    3,
		basicpay: 3000,
	}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
	fmt.Println()
	fmt.Println(pemp1.CalculateSalary())
	fmt.Println(cemp1.CalculateSalary())

	freelancer1 := Freelancer{
		empId:       4,
		ratePerHour: 70,
		totalHours:  120,
	}
	freelancer2 := Freelancer{
		empId:       5,
		ratePerHour: 100,
		totalHours:  100,
	}

	employees2 := []SalaryCalculator{pemp1, pemp2, cemp1, freelancer1, freelancer2}
	totalExpense(employees2)
	fmt.Println()
	fmt.Println(freelancer1.CalculateSalary())

}

type Worker interface {
	Work()
}

type Person struct {
	name string
}

func (p Person) Work() {
	fmt.Println(p.name, "is working")
}

func describe(w Worker) {
	fmt.Printf("Interface type %T value %v\n", w, w)
}

func interfaceInternalRepresentationExample() {
	p := Person{
		name: "Naveen",
	}
	var w Worker = p
	describe(w)
	w.Work()
}

func describe2(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

// An interface that has zero methods is called an empty interface and can be implemented by all types. Represented as interface{}
func emptyInterFaceExample() {
	s := "Hello World"
	describe2(s)
	i := 55
	describe2(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe2(strt)
}

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
	// p := i.(int) //program panics when string is given, so use above way of checking
	// fmt.Println(p)
	q := i
	fmt.Println(q)
}

// Type assertion is used to extract the underlying value of the interface. i.(T) is the syntax
func typeAssertionExample() {
	var s interface{} = 56
	assert(s)
	var i interface{} = "Steven Paul"
	assert(i)
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

// A type switch is used to compare the concrete type of an interface against multiple types specified in various case statements.
func typeSwitchExample() {
	findType("Naveen")
	findType(77)
	findType(89.98)
}

type Describer interface {
	Describe()
}
type Person2 struct {
	name string
	age  int
}

func (p Person2) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType2(i interface{}) {
	switch v := i.(type) {
	case Describer:
		fmt.Println("entered case")
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}

	switch i.(type) {
	case Describer:
		fmt.Println("\nEntered case without call")
	default:
		fmt.Printf("unknown type\n")
	}
}

func CompareInterfaceTypeExample() {
	findType2("Naveen")
	p := Person2{
		name: "Naveen R",
		age:  25,
	}
	findType2(p)

}

package functionspackages

import "fmt"

func FunctionUsage() string {
	fmt.Println("The signature of Go functions is func functionname(parametername type) returntype { \n //function body \n}")
	// The parameters and return type are optional in a function.

	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	area, perimeter = rectPropsWithNamedReturn(12.8, 6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	area, _ = rectProps(10, 5) //You can pass int values for float but not float for int
	fmt.Printf("Area %f ", area)

	return "\n\nend of functions practice\n\n"
}

func calculateBill(price int, no int) int {
	// func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// It is possible to return named values from a function. If a return value is named, it can be considered as
// being declared as a variable in the first line of the function.

func rectPropsWithNamedReturn(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

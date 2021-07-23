package mapsstrings

import (
	"fmt"
	"reflect"
)

func Maps() string {
	// MapCreationRetreiveValDeleteKeys()
	// MapsWithStruct()
	// CopyEditMaps()
	CompareMaps()

	return "end of maps function"
}

func MapCreationRetreiveValDeleteKeys() {
	employeeSalary := make(map[string]int)
	fmt.Println(employeeSalary)
	employeeSalary["steve"] = 12000
	employeeSalary["jamie"] = 15000
	employeeSalary["mike"] = 9000
	fmt.Println("employeeSalary map contents:", employeeSalary)

	employeeSal := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	employeeSal["mike"] = 9000
	fmt.Println("employeeSalary map contents:", employeeSal)

	// The zero value of a map is nil. If you try to add elements to a nil map, a run-time panic will occur
	// var employeeS map[string]int
	// employeeSalary["steve"] = 12000			//causes error-assignment to entry in nil map

	fmt.Println("Retreiving value")
	employee := "jamie"
	salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)

	fmt.Println("Salary of joe is", employeeSalary["joe"])

	newEmp := "joe"
	value, ok := employeeSalary[newEmp]
	if ok == true { //or u can use if ok{
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp, "not found")
	}

	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	}

	delete(employeeSalary, "steve")
	fmt.Println("map after deletion", employeeSalary)

	fmt.Println("length is", len(employeeSalary))

}

func MapsWithStruct() {

	type employee struct { //can be kept on top too for global scope
		salary  int
		country string
	}
	emp1 := employee{
		salary:  12000,
		country: "USA"} //if closing brackets are in next line u will need comma after usa.
	// This is because lexer auto adds semicolon at end of line. and this would result in end of stmt causing the lexer to throw an error.

	employeeInfo := map[string]employee{
		"Steve": emp1,
		"Jamie": employee{ //using employee here is redundant as we are already declaring employee as value tYpe during msp initialization. gives warning-redundant type from array, slice, or map composite literal
			salary:  14000,
			country: "Canada",
		},
		"Mike": {
			salary:  13000,
			country: "Idndia",
		},
	}

	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s Salary:$%d  Country: %s\n", name, info.salary, info.country)
	}

}

func CopyEditMaps() {

	// maps are reference types. When a map is assigned to a new variable, they both point to the same internal data structure.
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	fmt.Println("Original employee salary", employeeSalary)
	modified := employeeSalary
	modified["mike"] = 18000
	fmt.Println("Employee salary changed", employeeSalary)

	//To avoid refernce copies follow this:
	employeeSalaryCopy := make(map[string]int)
	for k, v := range employeeSalary {
		employeeSalaryCopy[k] = v
	}
	employeeSalaryCopy["mike"] = 133000
	fmt.Println("Employee salary changed", employeeSalary)
	fmt.Println("Employee salary changed", employeeSalaryCopy)

}

func CompareMaps() {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	map2 := map1

	// if map1 == map2 {    //gives error as map can only be compared to nil
	// }

	//we can compare each one's individual elements one by one.
	res := true
	for k, v := range map1 {
		if w, ok := map2[k]; !ok || v != w {
			res = false
		}
	}
	fmt.Println("value comparision: ", res)

	// The other way is using reflection
	// 	reflect.DeepEqual(map1, map2) returns true if both the maps satisfy the following conditions:
	// Both maps are nil or non-nil.
	// Both maps have the same length.
	// Both maps are the same map object or their corresponding keys map to deeply equal values.
	fmt.Println(reflect.DeepEqual(map1, map2))

	map3 := map1
	delete(map3, "two")
	fmt.Println(reflect.DeepEqual(map1, map3)) //returns true because of maps store reference values.so two delted in map1 &map3 and also map2
	fmt.Println(reflect.DeepEqual(map1, map2))
	fmt.Println(map1, "\n", map2, "\n", map3)

	mapCopy := make(map[string]int)
	for k, v := range map1 {
		mapCopy[k] = v
	}
	mapCopy["4dwr"] = 4
	fmt.Println(reflect.DeepEqual(map1, mapCopy)) //returns true because of maps store reference values.so two delted in map1 &map3 and also map2
	fmt.Println(map1, "\n", mapCopy)

}

package reflectionpractice

import (
	"fmt"
	"reflect"
)

func ReflectionPractice() {
	withoutReflectionsExample()
	reflectionExample()
	reflectionUseCaseExample()
}

type order struct {
	ordId      int
	customerId int
}

func createQuery1(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}

func withoutReflectionsExample() {
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	fmt.Println(createQuery1(o))
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("kind:", k)
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}

}

func reflectionExample() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery3(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}

func reflectionUseCaseExample() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery3(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery3(e)
	i := 90
	createQuery(i)

}

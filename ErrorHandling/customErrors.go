package errorhandling

import (
	"errors"
	"fmt"
	"math"
)

func CustomErrorsPractice() {
	basicCustomError()
	customErrorFormatted()
	customErrorWithStructField()
	customErrorWithStructMethod()
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}

func basicCustomError() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}

func circleArea2(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

func customErrorFormatted() {
	radius := -20.0
	area, err := circleArea2(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)

}

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func circleArea3(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func customErrorWithStructField() {
	radius := -20.0
	area, err := circleArea3(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}

type rectAreaError struct {
	err    string  //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

func (e *rectAreaError) Error() string {
	return e.err
}

func (e *rectAreaError) lengthNegative() bool {
	return e.length < 0
}

func (e *rectAreaError) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &rectAreaError{err, length, width}
	}
	return length * width, nil
}

func customErrorWithStructMethod() {
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*rectAreaError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)

			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)

			}
			return
		}
	}
	fmt.Println("area of rect", area)
}

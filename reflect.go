package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	Name string
	Id   int
}

func main() {
	var e = Employee{
		Name: "one",
		Id:   1,
	}

	x := reflect.TypeOf(e)
	fmt.Println(x)

	y := reflect.ValueOf(e)
	fmt.Println(y)

	w := y.FieldByName("Name")
	fmt.Println(w)

	z := reflect.ValueOf(&e).Elem()
	z.FieldByName("Name").SetString("two")
	fmt.Println(e.Name)
}

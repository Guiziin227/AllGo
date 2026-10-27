package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x any = "50"

	t := reflect.TypeOf(x)
	fmt.Printf("Tipo de x: %v\n", t)

	v := reflect.ValueOf(x)
	if v.CanInt() {
		xInt := v.Int()
		fmt.Printf("Valor de x: %v\n", xInt)
	} else {
		fmt.Println("Valor de x não é um inteiro")
	}

	var y any = 100

	fmt.Printf("Valor de y: %v\n", y)

	v2 := reflect.ValueOf(&y).Elem()
	v2.Set(reflect.ValueOf(200))
	fmt.Printf("Novo valor de y: %v\n", y)
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 42
	fmt.Println("Type:", reflect.TypeOf(x))
	fmt.Println("Value: ", reflect.ValueOf(x))
}

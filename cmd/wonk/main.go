package main

import (
	"fmt"
	"reflect"
)

func main() {

	fnValue := reflect.ValueOf(doAThing)

	out := fnValue.Call([]reflect.Value{})

	fmt.Printf("%d", len(out))
}

func doAThing() (*string, error) {
	str := "Hello, World!"
	return &str, nil
}
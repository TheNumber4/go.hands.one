package main

import (
	"log"
	"reflect"
)

type myStruct struct {
	message string
}

func magicMethod(data interface{}, function interface{}) {
	dv := reflect.ValueOf(data)
	fv := reflect.ValueOf(function)

	if fv.Type().Kind() != reflect.Func ||
		fv.Type().In(0) != dv.Type() {
		panic("That's a crap call !")
	}

	fv.Call([]reflect.Value{dv})
}

func main() {
	magicMethod(myStruct{message: "Hello"}, func(input myStruct) {
		log.Printf("Input: %s", input.message)
	})
	/* Crap call
	magicMethod(myStruct{message: "Hello"}, func(input int) {
		log.Printf("Input: %d", input)
	})
	*/
}

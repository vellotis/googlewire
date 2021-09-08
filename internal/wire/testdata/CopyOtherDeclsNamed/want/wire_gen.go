// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"fmt"
)

// Injectors from foo.go:

func injectedMessage() string {
	string2 := provideMessage()
	return string2
}

// foo.go:

func main() {
	fmt.Println(injectedMessage())
}

// provideMessage provides a friendly user greeting.
func provideMessage() string {
	return "Hello, World!"
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

// Injectors from wire.go:

func injectedMessage() string {
	s := provideS()
	wireA := s.A
	return wireA
}

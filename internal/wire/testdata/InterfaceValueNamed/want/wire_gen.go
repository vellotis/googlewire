// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"io"
	"strings"
)

// Injectors from wire.go:

func injectedReader() io.Reader {
	ioReader := _wireReaderValue
	return ioReader
}

var (
	_wireReaderValue = strings.NewReader("hello world")
)
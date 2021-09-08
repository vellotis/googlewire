// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

// Injectors from wire.go:

func injectBar() (Bar, func(), error) {
	fooA, cleanup := provideFoo()
	fooB, cleanup2 := provideFoo()
	bar, err := provideBar(fooA, fooB)
	if err != nil {
		cleanup2()
		cleanup()
		return 0, nil, err
	}
	return bar, func() {
		cleanup2()
		cleanup()
	}, nil
}

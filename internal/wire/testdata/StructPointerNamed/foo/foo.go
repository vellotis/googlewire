// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/google/wire"
)

func main() {
	fb := injectFooBar()
	e := injectEmptyStruct()
	fmt.Printf("%d %d %v\n", fb.FooA, fb.FooB, e)
}

type Foo int

type FooBar struct {
	FooA Foo `wire:"fooA"`
	FooB Foo `wire:"fooB"`
}

type Empty struct{}

var fooValue Foo = 21
func provideFoo() Foo {
	defer func() {
		fooValue*=2
	}()
	return fooValue
}

var Set = wire.NewSet(
	wire.Struct(new(FooBar), "*"),
	wire.Named("fooA", provideFoo),
	wire.Named("fooB", provideFoo),
)
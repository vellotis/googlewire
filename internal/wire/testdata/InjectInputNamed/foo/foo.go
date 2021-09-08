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
	fmt.Println(injectFoo(40))
}

type Foo int

var Set = wire.NewSet(
	wire.Named("fooB", provideFoo),
	wire.Named("fooResult", provideAddedFoo),
)

func provideFoo() Foo {
	return 2
}

func provideAddedFoo(fooA_wired Foo, fooB_wired Foo) Foo {
	return fooA_wired + fooB_wired
}

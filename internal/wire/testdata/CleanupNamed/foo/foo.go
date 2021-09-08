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
)

func main() {
	foo, cleanup := injectBar()
	fmt.Println(*foo)
	cleanup()
	fmt.Println(*foo)
}

type Foo int

func provideFooA() (*Foo, func()) {
	foo := new(Foo)
	*foo = 42
	return foo, func() { *foo = 0 }
}

func provideFooB(fooA_wired *Foo) (*Foo, func()) {
	fooB := new(Foo)
	*fooB = 77
	return fooB, func() {
		if *fooA_wired == 0 {
			panic("foo A cleaned up before foo B")
		}
		*fooA_wired = 0
	}
}

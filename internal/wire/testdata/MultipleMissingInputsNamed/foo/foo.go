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

import "fmt"

func main() {
	fmt.Println(injectMissingOutputType())
	fmt.Println(injectMultipleMissingTypes())
	fmt.Println(injectMissingRecursiveType())
}

type Foo int
type Bar int
type Baz int

func provideBaz(foo_wired Foo, bar_wired Bar) (baz_wired Baz) {
	return 0
}

type Zip int
type Zap int
type Zop int

func provideZip(foo_wired Foo) (zip_wired Zip) {
	return 0
}

func provideZap(zip_wired Zip) (zap_wired Zap) {
	return 0
}

func provideZop(zap_wired Zap) (zop_wired Zop) {
	return 0
}

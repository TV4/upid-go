# upid-go

[![Build Status](https://travis-ci.com/TV4/upid-go.svg?branch=master)](https://travis-ci.com/TV4/upid-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/TV4/upid-go)](https://goreportcard.com/report/github.com/TV4/upid-go)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/TV4/upid-go)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/TV4/upid-go#license)

`upid-go` is a Go package for encoding/decoding UPIDs. A UPID is a canonical
string that combines a user ID and a profile ID.

```bnf
<upid> ::= <user-id> | <user-id> "_" <profile-id>
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/TV4/upid-go"
)

func main() {
	inputUserID, inputProfileID := "user-id", "profile-id"
	up := upid.Encode(inputUserID, inputProfileID)
	fmt.Printf("Encode(%q, %q) -> %q\n", inputUserID, inputProfileID, up)

	inputUPID := "user-id_profile-id"
	outputUserID, outputProfileID := upid.Decode(inputUPID)
	fmt.Printf("Decode(%q) -> %q, %q\n", inputUPID, outputUserID, outputProfileID)
}
```
```
Encode("user-id", "profile-id") -> "user-id_profile-id"
Decode("user-id_profile-id") -> "user-id", "profile-id"
```

## License
Copyright (c) 2019 TV4

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

# Randstr
[![Build Status](https://travis-ci.org/m1ome/randstr.svg?branch=master)](https://travis-ci.org/m1ome/randstr)
[![Coverage Status](https://coveralls.io/repos/github/m1ome/randstr/badge.svg?branch=master)](https://coveralls.io/github/m1ome/randstr?branch=master)
[![GoDoc](https://godoc.org/github.com/m1ome/randstr?status.svg)](https://godoc.org/github.com/m1ome/randstr)
[![Go Report Card](https://goreportcard.com/badge/github.com/m1ome/randstr)](https://goreportcard.com/report/github.com/m1ome/randstr)

> Get a random string or byte

## Install
```bash
go get github.com/m1ome/randstr
```

## Usage
```go
package main

import (
    "fmt"

    "github.com/m1ome/randstr"
)

func main() {
    fmt.Printf("Random string: %s\n", randstr.GetString(10))
}

```
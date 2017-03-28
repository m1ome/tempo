# Tempo
> Get a random temporary file or directory path

[![Build Status](https://travis-ci.org/m1ome/tempo.svg?branch=master)](https://travis-ci.org/m1ome/tempo)
[![Coverage Status](https://coveralls.io/repos/github/m1ome/tempo/badge.svg?branch=master)](https://coveralls.io/github/m1ome/tempo?branch=master)
[![GoDoc](https://godoc.org/github.com/m1ome/tempo?status.svg)](https://godoc.org/github.com/m1ome/tempo)
[![Go Report Card](https://goreportcard.com/badge/github.com/m1ome/tempo)](https://goreportcard.com/report/github.com/m1ome/tempo)

## Install
```bash
go get github.com/m1ome/tempo
```

## Usage
```go
package main

import (
    "fmt"

    "github.com/m1ome/tempo"
)

func main() {
    fmt.Printf("Tempo directory: %s\n", tempo.Dir())
    fmt.Printf("Tempo file: %s\n", tempo.File(".png"))
}
```
# Transliteration

GO library for transliteration. 🇺🇦 🇬🇧 🔡

[![Build Status](https://img.shields.io/github/workflow/status/fre5h/transliteration-go/CI/main?style=flat-square)](https://github.com/fre5h/transliteration-go/actions?query=workflow%3ACI+branch%3Amain+)
[![CodeCov](https://img.shields.io/codecov/c/github/fre5h/transliteration-go.svg?style=flat-square)](https://codecov.io/github/fre5h/transliteration-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/fre5h/transliteration-go?style=flat-square)](https://goreportcard.com/report/github.com/fre5h/transliteration-go)
[![License](https://img.shields.io/github/license/fre5h/transliteration-go?style=flat-square)](https://pkg.go.dev/github.com/fre5h/transliteration-go)
[![Gitter](https://img.shields.io/badge/gitter-join%20chat-brightgreen.svg?style=flat-square)](https://gitter.im/fre5h/transliteration-go)
[![GoDoc](https://godoc.org/github.com/dgrijalva/jwt-go?status.svg)](https://godoc.org/github.com/dgrijalva/jwt-go)

## Requirements 🧐

* GO >= 1.16

## Available transliteration methods 🎁

<table>
    <thead>
        <tr>
            <th>From</th>
            <th>To</th>
            <th>Rules</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Ukrainian</td>
            <td>Latin</td>
            <td>
                Resolution of the Cabinet of Ministers of Ukraine №55 dated January 27, 2010
                <br />
                https://zakon.rada.gov.ua/laws/show/55-2010-%D0%BF#Text
            </td>
        </tr>
    </tbody>
</table>

## Using 👨‍🎓

###### main.go

```go
package main

import (
    "fmt"
    "os"
    "strings"

    "github.com/fre5h/transliteration-go"
)

func main() {
    inputString := strings.Join(os.Args[1:], " ")
    transliteratedString := transliteration.UkrToLat(inputString)
    
    fmt.Println("transliterated text: ", transliteratedString)
}
```

###### bash

```bash
$ go get "github.com/fre5h/transliteration-go"@v1.0.2
go: downloading github.com/fre5h/transliteration-go v1.0.2
go: added github.com/fre5h/transliteration-go v1.0.2

$ go run main.go Слава Україні!
transliterated text: Slava Ukraini!
```

### Some examples of *Ukrainian-to-Latin* transliteration ℹ️

<table>
    <thead>
        <tr>
            <th>Ukrainian text</th>
            <th>Transliterated text</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Володимир</td>
            <td>Volodymyr</td>
        </tr>
        <tr>
            <td>Богдан</td>
            <td>Bohdan</td>
        </tr>
        <tr>
            <td>Жанна</td>
            <td>Zhanna</td>
        </tr>
        <tr>
            <td>Наталія</td>
            <td>Nataliia</td>
        </tr>
        <tr>
            <td>Олексій</td>
            <td>Oleksii</td>
        </tr>
        <tr>
            <td>Уляна</td>
            <td>Uliana</td>
        </tr>
        <tr>
            <td>Юрій</td>
            <td>Yurii</td>
        </tr>
    </tbody>
</table>

## Contributing 🤝

See [CONTRIBUTING](https://github.com/fre5h/transliteration/blob/master/.github/CONTRIBUTING.md) file.

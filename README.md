# Transliteration

GO library for transliteration. 🇺🇦 🇬🇧 🔡

[![Build Status](https://img.shields.io/github/workflow/status/fre5h/transliteration-go/CI/main?style=flat-square)](https://github.com/fre5h/transliteration-go/actions?query=workflow%3ACI+branch%3Amain+)
[![CodeCov](https://img.shields.io/codecov/c/github/fre5h/transliteration-go.svg?style=flat-square)](https://codecov.io/github/fre5h/transliteration-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/fre5h/transliteration-go?style=flat-square)](https://goreportcard.com/report/github.com/fre5h/transliteration-go)
[![License](https://img.shields.io/github/license/fre5h/transliteration-go?style=flat-square)](https://pkg.go.dev/github.com/fre5h/transliteration-go)
[![Gitter](https://img.shields.io/badge/gitter-join%20chat-brightgreen.svg?style=flat-square)](https://gitter.im/fre5h/transliteration-go)
[![GoDoc](https://pkg.go.dev/badge/github.com/fre5h/transliteration-go)](https://pkg.go.dev/github.com/fre5h/transliteration-go)


## Requirements 🧐

* GO >= 1.18

## Available transliteration methods 🎁

| From      | To    | Rules                                                                                                                                     |
|-----------|-------|-------------------------------------------------------------------------------------------------------------------------------------------|
| Ukrainian | Latin | Resolution of the Cabinet of Ministers of Ukraine №55 dated January 27, 2010<br/> https://zakon.rada.gov.ua/laws/show/55-2010-%D0%BF#Text |

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
$ go get "github.com/fre5h/transliteration-go"@v1.1.5
go: downloading github.com/fre5h/transliteration-go v1.1.5
go: added github.com/fre5h/transliteration-go v1.1.5

$ go run main.go Слава Україні!
transliterated text: Slava Ukraini!
```

### Some examples of *Ukrainian-to-Latin* transliteration ℹ️

| Ukrainian text | Transliterated text |
|----------------|---------------------|
| Володимир      | Volodymyr           |
| Богдан         | Bohdan              |     
| Жанна          | Zhanna              |
| Наталія        | Nataliia            |
| Олексій        | Oleksii             |
| Уляна          | Uliana              |
| Юрій           | Yurii               |

## Contributing 🤝

See [CONTRIBUTING](https://github.com/fre5h/transliteration-go/blob/master/.github/CONTRIBUTING.md) file.

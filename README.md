#### Introduction

This package provides golang bindings to the [Citibike API](https://www.citibikenyc.com/system-data)


#### Usage

Install golang

```
brew install golang
```

Install this package:

```
go get https://github.com/josephmisiti/go-citibike
```

Use it in a program like this:

```go
package main

import(
    "fmt"
    "github.com/josephmisiti/go-citibike"
)

func main() {
    fmt.Println("------")
    API := new(citibike.API)
    stations, err := API.GetStations()
    if( err != nil ) {
        fmt.Println(err)
    }
    fmt.Println(stations)
}

```

#### Running the tests

```
make test
```
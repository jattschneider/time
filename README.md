# Go Time

Time is a Go library that implements time utilities (eg StopWatch).

## Getting Started

Just a quick example how to use the library:

#### main.go
```
package main

import (
	"fmt"

	"github.com/jattschneider/time"
)

func main() {
    sw := &time.StopWatch{}
    fmt.Println(sw.Summary())
    sw.Start("task 1")
    time.Sleep(2 * time.Nanosecond)
    sw.Stop()
    fmt.Println(sw.String())
}

```

```
$ go run main.go
```

### Installing

```
go get -v github.com/jattschneider/time
```

## Built With

* [Go](https://golang.org/) - The Go Programming Language

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/jattschneider/argonauts/tags). 

## Authors

* **José Augusto Schneider** - *Initial work* - [jattschneider](https://github.com/jattschneider)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

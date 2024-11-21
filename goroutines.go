package pkg

import (
	"fmt"
	"runtime"
)

var p = fmt.Println

func main() {
	p("main execution started")
	p("No. of CPUs:", runtime.NumCPU())
	p("No. of Goroutines:", runtime.NumGoroutine())

	p("OS:", runtime.GOOS)
	p("Arch:", runtime.GOARCH)

	p("GOMAXPROCS", runtime.GOMAXPROCS(0))

}

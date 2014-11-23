// Package hi provides a function for saying hello.
package hi

import (
	"fmt"
	"golang.org/x/mobile/example/libhello/clinfo"
)

func Hello(name string) string{
	fmt.Printf("Hello, %s!\n", name)
	return clinfo.GOCLInfo()
}

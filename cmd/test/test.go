package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	var f *os.File
	for _, c := range []io.Closer{f} {
		fmt.Println(c)
		fmt.Println("c is nil?", c == nil)
	}

	d, _ := os.Open("go.mod")
	defer d.Close()
	fmt.Println(io.ReadAll(d))
	for _, c := range []io.Closer{d} {
		fmt.Println(c)
		fmt.Println("c is nil?", c == nil)
	}
}

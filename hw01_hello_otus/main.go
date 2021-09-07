package main

import "fmt"

import "golang.org/x/example/stringutil"

var s = "Hello, OTUS!"

func main() {
	fmt.Printf("%s", stringutil.Reverse(s))
}

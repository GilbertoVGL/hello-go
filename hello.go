package main

import (
	"fmt"

	"example.com/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Printf("hello, world... again\n")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println("--------------------------------------------------")
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}

package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var src []string
	flag.Parse()
	src = flag.Args()

	if src[0] == "" {
		fmt.Println("0")
	} else {
		fmt.Println(len(strings.Split(src[0], " ")))
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	tmp, _ := reader.ReadString('\n')
	str := strings.Split(strings.Trim(tmp, "\n"), " ")

	if strings.Trim(tmp, "\n") == "" {
		fmt.Println(0)
	} else {
		fmt.Println(len(str))
	}
}

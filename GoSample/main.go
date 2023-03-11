package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world kento"
	v := strings.Fields(s)
	fmt.Println(v)
}

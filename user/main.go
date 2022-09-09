package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "1,2,3,4"
	ss := strings.Split(s, ",")
	fmt.Println(ss)
	fmt.Println(strconv.Atoi(ss[0]))
}

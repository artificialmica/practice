package main

import (
	"fmt"
	"practice"

)

func main() {
	s := "Hello World!"
	s = practice.StrRev(s)
	fmt.Println(s)
}

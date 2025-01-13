package main

import "github.com/01-edu/z01"

func main(){
	for i:=122; i >=97;i--{
		y:= rune(i)
		z01.PrintRune(y)
	}
	z01.PrintRune('\n')
}
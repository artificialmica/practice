package main

import "github.com/01-edu/z01"

func main(){

	
	for i:=97; i <=122;i++{
		y:= rune(i)
		z01.PrintRune(y)
	}
	z01.PrintRune('\n')
}

package main

import "fmt"

func main() {
	hex()
}
func hex()  {
	fmt.Printf("%b - %d - %x\n", 5, 5, 5)
	fmt.Printf("%b - %d - %X\n", 5, 5, 5)
	fmt.Printf("%b - %d - %#x\n", 5, 5, 5)
	fmt.Printf("%b - %d - %#X\n", 5, 5, 5)
	fmt.Printf("%b - %d - %x\n", 5, 5, 5)
	for i:=0;i<500 ;i++  {
		fmt.Printf("%b \t %o \t %d \t %x \t %q\n",i,i,i,i,i)
	}
}


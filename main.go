package main

import (
	"fmt"
)

func main() {

	var ans int
	fmt.Print("Create New Prof\t1\nLogin\t\t2\nDelete Account\t3\n")
	fmt.Print("Enter: ")
	fmt.Scanf("%d", &ans)

	if ans == 1 {
		Newpro()
	} else if ans == 2{
		Login()
	} else {
		Delete()
	}
}
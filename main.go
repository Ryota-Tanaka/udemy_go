package main

import "fmt"

func main() {
	//fmt.Println()
	calculation()
}

func calculation() {
	var (
		abcd = 23
		efgh = 21
		aiu  int
	)

	if abcd > 20 {
		aiu = abcd + efgh
	}

	fmt.Println(aiu)
}

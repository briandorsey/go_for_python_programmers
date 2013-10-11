package main

import "fmt"

func double(a [3]int) {
	for i := range a {
		a[i] *= 2
	}
}

func main() {
	a := [3]int{1, 2, 3}
	fmt.Println(a)

	double(a)
	fmt.Println(a)
}

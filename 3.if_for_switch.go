package main

import "fmt"

func conditional() {

	// if - else
	i := 10
	if i <= 10 {
		fmt.Println("i is less or equal to 10")
	} else {
		fmt.Println("i is large than 10")
	}

	// in one line
	if j := 10; j > 10 {
		fmt.Println("j > 10")
	} else {
		fmt.Println("j <= 10")
	}

	// switch
	k := 5
	switch {
	case k == 1:
		fmt.Println("k = 1")
	case k == 5:
		fmt.Println("k = 6")
	default:
		fmt.Println("default")
	}

	switch k {
	case 1:
		fmt.Println("k = 1")
	case 5:
		fmt.Println("k = 6")
	default:
		fmt.Println("default")
	}

	// for loop
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

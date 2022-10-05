package main

import "fmt"

func set() {

	// array
	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array)
	fmt.Println(array[3])
	fmt.Println(len(array), cap(array))

	array = [...]string{"a", "b", "c", "d", "e"}
	fmt.Println(array)
	fmt.Println(array[3])

	for i, v := range array {
		fmt.Println(i, v)
	}

	// slice
	slice := array[2:5]
	fmt.Println(slice)

	slice[2] = "z"
	fmt.Println(array)
	fmt.Println(slice)

	slice1 := make([]string, 4)
	fmt.Println(slice1)
	fmt.Println(len(slice1), cap(slice1))

	slice2 := make([]string, 4, 8)
	fmt.Println(slice2)
	fmt.Println(len(slice2), cap(slice2))

	slice2 = append(slice2, "a")
	fmt.Println(slice2)

	slice3 := []string{"a", "b", "c"}
	fmt.Println(slice3)
	fmt.Println(len(slice3), cap(slice3))

	slice3 = append(slice3, "d")
	fmt.Println(slice3)

	var slice4 []string
	fmt.Println(slice4)
	fmt.Println(len(slice4), cap(slice4))

	slice5 := make([]string, 0)
	fmt.Println(slice5)
	fmt.Println(len(slice5), cap(slice5))

	slice6 := make([]string, 3, 4)
	fmt.Println(slice6)
	fmt.Println(len(slice6), cap(slice6))

	slice7 := append(slice6, "a")
	fmt.Println(slice7)

	slice8 := append(slice6, "b")
	fmt.Println(slice7)
	fmt.Println(slice8)

	slice9 := append(slice6, "d", "e")
	fmt.Println(slice7)
	fmt.Println(slice8)
	fmt.Println(slice9)

	// map
	m := make(map[string]int)
	m["a"] = 1
	fmt.Println(m)

	m = map[string]int{}
	m["b"] = 2
	fmt.Println(m)

	num, ok := m["b"]
	if ok {
		fmt.Println(num)
	}

	delete(m, "b")
	fmt.Println(len(m))

	m["a"] = 1
	for k, v := range m {
		fmt.Println(k, v)
	}

	// string and []byte
	s := "hello"
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println(len(bs))
}

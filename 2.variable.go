package main

import (
	"fmt"
	"strconv"
	"strings"
)

func variable() {
	// int
	var i int = 10
	fmt.Println(i)

	var j = 20
	fmt.Println(j)

	var (
		x = 30
		y = 40
	)
	fmt.Println(x)
	fmt.Println(y)

	// float
	var f32 float32 = 2.2
	var f64 float64 = 10.3456
	fmt.Println("f32 is", f32, ", f64 is", f64)

	// bool
	var bl bool = true
	fmt.Println("bl is", bl)

	// string
	var s1 string = "hello"
	var s2 string = "world"
	fmt.Println(s1, s2)

	// default value
	var zi int
	var zf float64
	var zb bool
	var zs string
	fmt.Println(zi, zf, zb, zs)

	// short statement
	k := 10
	bf := false
	s3 := "word"
	fmt.Println(k, bf, s3)

	// pointer
	pi := &i
	fmt.Println(pi)
	fmt.Println(*pi)

	// change value
	i = 50
	fmt.Println(i)

	// const
	const name = "name"
	fmt.Println(name)

	// iota
	// before:
	const (
		one   = 1
		two   = 2
		three = 3
		four  = 4
	)
	fmt.Println(one, two, three, four)
	// after:
	const (
		one_ = iota + 1
		two_
		three_
		four_
	)
	fmt.Println(one_, two_, three_, four_)

	// conversion
	int2string := strconv.Itoa(i)
	fmt.Println(int2string)

	string2int, err := strconv.Atoi(int2string)
	fmt.Println(string2int, err)

	int2float := float64(i)
	fmt.Println(int2float)

	float2int := int(int2float)
	fmt.Println(float2int)

	// strings package
	s1 = "Hash go"
	fmt.Println(strings.HasPrefix(s1, "H"))
	fmt.Println(strings.Index(s1, "h"))
	fmt.Println(strings.ToUpper(s1))
}

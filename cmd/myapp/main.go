package main

import (
	"fmt"
	"strconv"
)

func convert(a int) string {
	b := strconv.Itoa(a)
	return b
}

func compare(c int) string {
	var d string
	if c%2 == 0 {
		d = "Even"
	} else {
		d = "Odd"
	}
	return d
}

func main() {
	a := 100
	e := 7
	fmt.Println(convert(a))
	fmt.Println(compare(e))
	arr := []int{8, 23, -3, 41, -5}
	min := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	fmt.Println("the lowest value is:", min)
}

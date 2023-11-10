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

	var T, N, K, Time int
	fmt.Println("How much tests: ")
	fmt.Scan(&T)

	fmt.Println("How much students: ")
	fmt.Scan(&N)

	fmt.Println("How much students we need: ")
	fmt.Scan(&K)
	for i := 1; i <= T; i++ {
		ComeEarly := 0
		for j := 1; j <= N; j++ {
			fmt.Println("When came student №", j)
			fmt.Scan(&Time)
			if Time <= 0 {
				ComeEarly++
			}
		}
		if ComeEarly >= K {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}

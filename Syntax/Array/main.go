package main

import "fmt"

func main() {
	// ① 배열선언과 원소 초기화를 따로
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println("a[0], a[1]:", a[0], a[1])
	fmt.Println("a:", a)

	// ② 배열선언과 초기화를 동시에
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("primes:", primes)

	// ③ 동적 배열
	var scores1 []int
	scores2 := []int{}
	scores3 := make([]int, 3, 4)

	scores1 = append(scores1, 1, 2, 3, 4, 5, 6)
	scores2 = append(scores2, []int{1, 2, 3, 4}...)

	fmt.Println(scores1)
	fmt.Println(scores2)
	fmt.Println(scores3)
}

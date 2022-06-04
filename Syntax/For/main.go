package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 세미콜론 없이, C의 while과 비슷하게 쓸 수 있습니다.
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// for{} 로 무한 루프를 만들어보세요!
	// for {
	// }

	for j, k := 1, 2; j < 20 || k < 10; {
		j++
		k++
	}

	
}

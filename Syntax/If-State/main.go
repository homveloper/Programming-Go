package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	} else {
		//else문은 if문이 닫히는(}) 줄과 함께 쓰여야 합니다.
		//한칸 위로 올려서 제출해보세요!
		return fmt.Sprint(math.Sqrt(x))
	}
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// v는 if문 내부에서만 쓸 수 있고, 여기부터는 쓸 수 없습니다.

	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

}

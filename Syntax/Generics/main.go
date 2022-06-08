package main

import (
	"fmt"

	"github.com/liyue201/gostl/ds/vector"
	"golang.org/x/exp/constraints"
)

func Min(x, y float64) float64 {
	if x > y {
		return y
	}
	return x
}

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func main() {
	vector := vector.New()
	vector.PushBack("dawd")
	vector.PushBack(2)

	fmt.Println(vector)
}

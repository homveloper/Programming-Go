package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// 구조체 인스턴스 선언 방법
var (
	//① 일반적인 선언방식입니다. X가1, Y가 2로 초기화됩니다.
	v1 = Vertex{1, 2}
	//② X만 값을 지정해주고, Y는 int에  zero value로 설정됩니다.
	v2 = Vertex{X: 1}
	//③ X, Y모두 int에 zero value로 설정됩니다.
	v3 = Vertex{}
	v4 Vertex
)

func main() {
	fmt.Println("v1.X값:", v1.X)
	v1.X = 4
	fmt.Println("v1.X = 4로 바꾼 v1.X값:", v1.X)

	//④ 구조체 포인터로도 구조체의 값을 바꿀 수 있습니다.
	var p = &v1
	p.X = 10
	fmt.Println("포인터로 바꾼 v1.X값:", v1.X)

	fmt.Println(v3)
}

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // i를 가리키는 포인터
	fmt.Println(*p) // 포인터를 통해 i 값을 읽습니다.
	*p = 21         // 포인터를 통해 i값을 설정합니다.
	fmt.Println(i)

	p = &j       // j를 가리킵니다.
	*p = *p / 37 // 포인터를 통해 j를 나눕니다.
	fmt.Println(j)

	p_p := p
	fmt.Println(*p_p)
}

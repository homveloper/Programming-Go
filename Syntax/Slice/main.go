package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("배열 names:", names)

	fmt.Println("①슬라이스 선언")
	// 슬라이스 선언방법
	// ① 일반적인 선언방법 : 변수 선언과 비슷합니다. 슬라이스타입은 []type입니다.
	var s1 []string = names[0:3]
	// ② 슬라이스도 var키워드와 타입 명시를 생략할 수 있습니다.
	s2 := names[0:2]
	s3 := names[1:]

	fmt.Println("names[0:3]:", s1)
	fmt.Println("names[0:2]:", s2)
	fmt.Println("names[1:]:", s3)

	//s1에서 값을 바꾸면 names, s1에서도 바뀐 값을 볼 수 있습니다.
	fmt.Println("②슬라이스로 값 변경")
	fmt.Println("s1[0]", s1[0])
	s1[0] = "XXX"
	fmt.Println("s1[0] = XXX 실행 후 s1:", s1)
	fmt.Println("s1[0] = XXX 실행 후 s2:", s2)
	fmt.Println("s1[0] = XXX 실행 후 names:", names)

	s2 = s1[0:2]
	fmt.Println("s2 = s1[0:2] 실행 후  s2:", s2)

	//②구조체 슬라이스 리터럴
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("구조체 슬라이스 리터럴:", s)

	// make()로 가변 길이 배열 만들기
	d_a := make([]int, 5)
	fmt.Printf("a := make([]int, 5)의\t %v", d_a)
}

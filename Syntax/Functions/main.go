package main

import "fmt"

//① 매개변수 타입, 리턴 타입은 이름 뒤에 지정해줍니다
func add1(x int, y int) int {
	return x + y
}

//② 매개변수 x, y가 같은 타입일 때에는 타입을 한 번만 명시해 줄 수 있습니다.
func add2(x, y int) int {
	return x + y
}

//③ return 뒤에 리턴 타입을 적어주는 방법
func divide1(dividend, divisor int) (int, int) {
	var quotient = (int)(dividend / divisor)
	var remainder = dividend % divisor
	return quotient, remainder
}

//④ return뒤에 리턴할 변수를 선언하는 방법. ①과는 달리 함수 내부에서 `quotient`를 `var`로 선언하지 않고 바로 씁니다.
func divide2(dividend, divisor int) (quotient, remainder int) {
	quotient = (int)(dividend / divisor)
	remainder = dividend % divisor
	return //return이라고만 적으면 미리 return값으로 정해 놓은 quotient와 remainder를 return합니다.
}

func main() {
	// ①, ② 함수 호출
	fmt.Println("add1(x int, y int)의 결과: ", add1(42, 13))
	fmt.Println("add2(x, y int)의 결과: ", add2(42, 13))

	//③로 한 번에 여러개의 결과를 return받는 부분
	var quotient, remainder int
	quotient, remainder = divide1(10, 3)
	fmt.Println("③의 결과:", quotient, remainder)

	//④로 한 번에 여러개의 결과를 return받는 부분
	quotient, remainder = divide2(10, 3)
	fmt.Println("④의 결과:", quotient, remainder)

	//③로 변수 선언과 동시에 초기화
	var quotient1, remainder1 int = divide1(10, 3)
	fmt.Println("③의 결과:", quotient1, remainder1)
}

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	//① map 사용
	//map[string] 타입 변수 선언
	var mymap map[string]Vertex
	mymap2 := make(map[string]Vertex)

	//make()로 맵 생성
	mymap = make(map[string]Vertex)
	mymap["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println("① mymap[\"Bell Labs\"]: ", mymap["Bell Labs"])
	fmt.Println(mymap2)

	var mymap3 = make(map[int]int)
	fmt.Println(mymap3)

	//② map literal 사용
	var mymap_literal = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println("② mymap_literal[\"Bell Labs\"]", mymap_literal["Bell Labs"])

	m := make(map[string]int)
	//① key-value 지정하기
	m["Answer"] = 42
	fmt.Println("m[\"Answer\"]값은:", m["Answer"])

	//② key-value 삭제하기
	delete(m, "Answer")
	fmt.Println("m[\"Answer\"]값은", m["Answer"])

	//③ key존재 확인하기
	v, ok := m["Answer"]
	fmt.Println("m[\"Answer\"]값은", v, "존재하나요?", ok)

	m = make(map[string]int)
	m["Answer"] = 22
	v, ok = m["Answer"]
	fmt.Println("m[\"Answer\"]값은", v, "존재하나요?", ok)
}

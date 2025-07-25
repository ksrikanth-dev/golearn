package main

import "fmt"

func main() {
	//var arr [3]int = {1,2,3}
	//x := [3]int {1,2,3}
	//	var matrix [3][3]int
	// matrix := [3][3]int{ //----> 3 times of arr
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	//
	// //opt1
	// for i, row := range matrix {
	// 	for j, value := range row {
	// 		fmt.Printf("matrix[%d][%d]: %d", i, j, value)
	// 	}
	// 	fmt.Printf("\n")
	// }
	//
	// //opt2
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Printf("matrix[%d][%d]: %d", i, j, matrix[i][j])
	// 	}
	// 	fmt.Printf("\n")
	// }

	// s := []int{1, 2}
	// fmt.Printf("address of s: %p", s)
	// fmt.Printf("address of s: %p", &s[0])
	// fmt.Printf("value of s: %v", s)
	// s = append(s, 1, 2, 3, 4, 5, 6)
	// fmt.Printf("\naddress of s: %p", s)
	// fmt.Printf("value of s: %v", s)

	m := map[string]int{"Alice": 25, "Bob": 30}

	s := []map[string]int{}
	s1 := []map[string]int{}

	for k, v := range m {
		s = append(s, map[string]int{k: v})
	}
	fmt.Printf("%v", s)

	keys := []string{}
	for key, _ := range m {
		keys = append(keys, key)
	}
	fmt.Println("keys: ", keys)
	for _, key := range keys {
		s1 = append(s1, map[string]int{key: m[key]})
	}
	fmt.Printf("\n%v", s1)
}

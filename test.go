package main

import "fmt"

func main() {
	// 創建一個整數指針，並分配內存並初始化為零值（即 int 的默認值 0）
	numPtr := new(int)
	fmt.Println(*numPtr) // Output: 0

	s := "hello,"
	m := " world"
	a := s + m
	fmt.Printf("%s\n", a)

	var v []int
	v = append(v, 1, 2, 3)
	fmt.Println(v)
	fmt.Println(len(v))

	var numbers []int

	// 使用 append 函數將元素添加到slice中
	numbers = append(numbers, 1)
	numbers = append(numbers, 2, 3, 4)

	// 使用 append 函數添加另一個slice的所有元素
	moreNumbers := []int{5, 6, 7}
	numbers = append(numbers, moreNumbers...)

	fmt.Println(numbers) // Output: [1 2 3 4 5 6 7]
}

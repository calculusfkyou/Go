package main

import "fmt"

func main() {
	// 創建一個整數指針，並分配內存並初始化為零值（即 int 的默認值 0）
	numPtr := new(int)
	fmt.Println(*numPtr) // Output: 0
}

package main

import "fmt"

func main() {
	// 標準的な for loop
	for i:=0;i<5;i++ {
		fmt.Println(i)
	}

	// while 風の利用法
	cnt := 0
	for cnt < 5 {
		fmt.Println(cnt)
		cnt++
	}

	// for による無限ループ
	loop := 0
	for {
		fmt.Println(loop)
		loop++
		if loop == 5 {break}
	}
}

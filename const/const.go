package main

import "fmt"

func main() {
	// const 定数名
	const INT_CONST int = 10
	const Int_const int = 11
	const int_const int = 12
	const STRING_CONST string = "hoge"

	// ブレースで一括設定可能

	const (
		INT1 int       = 13
		INT2           = 13
		INT3           = 13
		STRING1 string = "hoge1"
		STRING2        = "hoge2"
		STRING3        = "hoge3"
	)

	fmt.Println(INT_CONST)
	fmt.Println(Int_const)
	fmt.Println(int_const)
	fmt.Println(STRING_CONST)

	fmt.Println(INT1)
	fmt.Println(INT2)
	fmt.Println(INT3)
	fmt.Println(STRING1)
	fmt.Println(STRING2)
	fmt.Println(STRING3)
}

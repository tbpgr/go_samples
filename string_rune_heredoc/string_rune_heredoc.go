package main

import "fmt"

func main() {
	// String はダブルクォートを使用する
	fmt.Println("string")

	// Rune でコードポイントを利用する [...] は暗黙長の配列宣言
	//  := で変数の宣言と初期化を一気に行う
	runes := [...]rune{'A', 'B', 'C'}

	// for range で for each 風に処理を行う
	// i を利用しない場合は、 _ にする
	for i, value := range runes {
		fmt.Print(i)
		fmt.Print(",")
		fmt.Println(value)
	}

	// バッククォートでヒアドキュメント
	fmt.Println(`
here doc1
here doc2
here doc3
		`)
}

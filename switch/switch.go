package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randoms := []int{rand.Intn(10), rand.Intn(10), rand.Intn(10)}

	// 値による分岐（列挙可能）
	for _, n := range randoms {
		switch n {
		case 1,2,3:
			fmt.Println("1-3")
		case 4,5,6:
			fmt.Println("4-6")
		case 7,8,9:
			fmt.Println("7-9")
		default:
			fmt.Println("other")
		}
	}

	// 式による分岐
	for _, n := range randoms {
		switch {
		case n % 2 == 0:
			fmt.Printf("%v is even.\n", n)
		case  n % 2 == 1:
			fmt.Printf("%v is odd.\n", n)
		}
	}

	// fallthrough により、breakせずに処理を継続
	for _, n := range randoms {
		msg := ""
		switch {
		case n % 3 == 0:
			msg += "3 で割ると余り 0 "
			fallthrough
		case n % 2 == 0:
			msg += "偶数 "
			fallthrough
		default:
			fmt.Println(msg)
		}
	}
}

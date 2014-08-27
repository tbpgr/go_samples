package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randoms := []int{rand.Intn(10), rand.Intn(10), rand.Intn(10)}

	for _, n := range randoms {
		if isEven := (n % 2 == 0) ; isEven {
			fmt.Printf("isEven = %s\n", strconv.FormatBool(isEven))
		}
		// スコープ外のため、エラーが発生する「undefined: isEven」
		// fmt.Println(isEven)
	}
}

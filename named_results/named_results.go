package main

import "fmt"

func main() {
	ret1, ret2 := useNamedResults("msg1", "msg2")
	fmt.Println(ret1, ret2)
	fmt.Println(useNamedResults("msg3", "msg4"))
}

func useNamedResults(message1, message2 string) (ret1, ret2 string) {
	ret1 = message1 + "-hoge"
	ret2 = message2 + "-hige"
	return
}
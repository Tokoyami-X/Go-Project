package study

import "fmt"

func Pointer() {
	a := 100
	b := &a
	fmt.Println(b)
	fmt.Println(a)

	*b = 200
	fmt.Println(^a)
}

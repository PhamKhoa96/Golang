package mypackage

import (
	"fmt"
)

var a int
var b string

var c int

// thay vì  "var c int = 1", ta có thể ghi "c := 1"
func Printab (a int, b int) int{
	var a1 int
	fmt.Println(a1)	// giá trị chưa được gán giá trị thì sẽ nhận mặc định zero-value

	var a2 = a + b + a1
	 if a2 > 3{
		 fmt.Println("a2 lớn hơn 3")
	 } else {
		 fmt.Println("a2 nhỏ hơn 3")
	 }

	fmt.Println(testPrivateFunc("Hung"))

	return a2
}
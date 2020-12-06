// Package stringutil contains utility functions for working with strings.
package mypackage

import "fmt"

// Reverse returns its argument string reversed rune-wise left to right.
//Public func	<- có thể được dùng bởi các package khác (với điều kiện là phải import package vào)
func TestPublicFunc(s string) string {
    return "Hello from TestPublicFunc - " + testPrivateFunc(s)
}
//Private func	<- chỉ sử dụng bên trong các file ở củng package
func testPrivateFunc(s string) string {
    return "Hello from testPrivateMethod " + s
}

func Tong() int{
    var sum int
    for i := 0 ; i < 10; i++ {
        sum += i
    }
    return sum
}

func NhanDoi(){
    sum1 := 1
    for ;sum1 < 20; {
        sum1 += sum1
        fmt.Println(sum1)
    }
}
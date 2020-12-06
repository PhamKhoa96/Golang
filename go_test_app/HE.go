package main

import (
	"fmt"
	"go_test_app/mypackage"
	"runtime"
)

func checkIFStateMent() {

	if a, b := 3, 5; a -  b < 0 {
		fmt.Printf("%d nho hơn %d", a, b)
	} else {
		fmt.Printf("%d lon hơn %d", a, b)
	}
}

func checkSwitchCase() {
	var i int = 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}

func checkOS() {
	fmt.Print("Go runs on ")
	
	os := runtime.GOOS
    switch os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s.", os)
    }
}


func main() {
	fmt.Println(mypackage.TestPublicFunc("Khoa"))
	fmt.Println(mypackage.Printab(5,6))
	fmt.Println(mypackage.Tong())
	mypackage.NhanDoi()
	checkIFStateMent()
	checkSwitchCase()
	checkOS()
}




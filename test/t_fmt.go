package main

import (
	"fmt"
)

func main() {
	fmt.Print("this", 1, "is", "fine")

	ifaces := []interface{}{1, "good", "case", "here"}
	// OK
	fmt.Print(ifaces...)

	strs := []string{"bad", "case", "here"}
	// fixme cannot use strs (variable of type []string) as []interface{} value in argument to fmt.Print
	//fmt.Print(strs...)

	ifaces2 := make([]interface{}, 0)
	for _, str := range strs {
		ifaces2 = append(ifaces2, str)
	}
	// OK now
	fmt.Print(ifaces2...)
}

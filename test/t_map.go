package main

import "fmt"

func main() {
	m := map[string]string{"a": "sdfs"}
	fmt.Println(m["b"] == "")
	fmt.Println(m["a"])
}

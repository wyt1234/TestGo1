package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func ABflag(ud string) string {
	if ud == "" {
		return "a"
	}
	s := 1
	for _, ch2 := range ud {
		if unicode.IsDigit(ch2) {
			atoi, _ := strconv.Atoi(string(ch2))
			s += atoi
		} else {
			s += int(ch2)
		}
	}
	if s%2 == 0 {
		return "a"
	} else {
		return "b"
	}
}

func main2() {
	str := "123方法一UXSFSDFjksdfjh"
	s := 1
	for _, ch2 := range str {
		if unicode.IsDigit(ch2) {
			atoi, _ := strconv.Atoi(string(ch2))
			s += atoi
		} else {
			s += int(ch2)
		}
	}
	fmt.Println(s)

}

func main() {
	bflag := ABflag("123")
	fmt.Println(bflag)
	s := 1234
	fmt.Println(s % 2)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename1("a/b/c.go"))
	fmt.Println(basename1("c.d.go"))
	fmt.Println(basename1("abc"))
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename1(s string) string {
	slash := strings.LastIndex(s, "/")

	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

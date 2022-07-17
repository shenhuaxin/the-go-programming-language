package main

import (
	"golang.org/x/net/html"
	"html"
	"os"
)

func main() {

	doc, err := html.Parse(os.Stdin)
}

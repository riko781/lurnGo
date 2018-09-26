package main

import (
	"fmt"
	"html"
)

func main() {
	const s = `"Francki & Henrick Carneva" <tasty@example.com>`
	fmt.Println(html.EscapeString(s))
}

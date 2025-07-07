package main

import (
	"bytes"
	"fmt"
)

func intsToString(value []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, c := range value {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", c)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(intsToString(a))
}

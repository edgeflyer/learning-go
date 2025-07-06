// package main

// import (
// 	"bytes"
// 	"fmt"
// )

// func comma(s string) string {
// 	var buf bytes.Buffer
// 	mark := len(s) % 3
// 	buf.WriteString(s[:mark-1])
// 	if len(s) > 3 {
// 		for i, c := range s[mark-1:] { //这里c是rune类型
// 			buf.WriteRune(c)
// 			if i%3 == 0 && i != len(s)-mark {
// 				buf.WriteByte(',')
// 			}
// 		}

// 	}
// 	return buf.String()
// }

// func main() {
// 	fmt.Printf(comma("12345"))
// }

// package main

// // cahrcount计算unicode字符的个数
// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// 	"unicode"
// 	"unicode/utf8"
// )

// func main() {
// 	counts := make(map[rune]int)    // unicode字符数量
// 	var utflen [utf8.UTFMax + 1]int // UTF-8编码的长度
// 	invalid := 0                    // 非法UTF-8字符数量

// 	in := bufio.NewReader(os.Stdin)
// 	for {
// 		r, n, err := in.ReadRune() // 返回rune，nbytes，error,读取一个unicode字符

// 		if err == io.EOF { // 输入的是文件结尾，退出循环
// 			break
// 		}
// 		if err != nil { // 读取错误，结束程序
// 			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
// 			os.Exit(1)
// 		}
// 		if r == unicode.ReplacementChar && n == 1 { // 无法解析的unicode字符
// 			invalid++
// 			continue
// 		}
// 		counts[r]++ // 键值对
// 		utflen[n]++
// 	}
// 	fmt.Printf("rune\tcount\n")
// 	for c, n := range counts {
// 		fmt.Printf("%q\t%d\n", c, n)
// 	}
// 	fmt.Print("\nlen\tcount\n")
// 	for i, n := range utflen {
// 		if i > 0 {
// 			fmt.Printf("%d\t%d\n", i, n)
// 		}
// 	}
// 	if invalid > 0 {
// 		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
// 	}
// }

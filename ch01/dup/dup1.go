// package main

// import (
// 	"bufio" //高效的处理输入和输出，其中Scanner是扫描器类型
// 	"fmt"
// 	"os"
// )

// func main() {
// 	counts := make(map[string]int)      //make用来新建map
// 	input := bufio.NewScanner(os.Stdin) //新建bufio.Sacnner类型的变量。可以通过调用input.Scan()来获取一行标准输入，以行或者单词为单位断开，os.Stdin是标准输入流
// 	for input.Scan() {
// 		counts[input.Text()]++ //input.Text()返回当前行的文本，counts[input.Text()]++表示将当前行的计数加1,scan在读取到一行的时候返回true，没有内容时返回false
// 	}
// 	for line, n := range counts { //迭代map，每次迭代的顺序都是随机的，可以防止程序依赖某种特定的序列
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

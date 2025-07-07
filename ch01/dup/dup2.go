// // 打印输出多次出现的行的个数和文本
// // 从stdin或指定的文件列表读取
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	counts := make(map[string]int)
// 	files := os.Args[1:]
// 	fmt.Println(files)
// 	if len(files) == 0 { //没有指定文件
// 		countLines(os.Stdin, counts) //os.Stdin是标准输入流，通常是键盘输入
// 	} else {
// 		for _, arg := range files { //files
// 			f, err := os.Open(arg) //os.Open()返回两个值：第一个是打开的文件(*os.File)，之后被scanner读取。第二个是内置的error类型的值，如果err等于nil，标准文件打开成功
// 			if err != nil {        //如果打开文件失败，打印错误信息，并继续下一个文件，其中err的值描述错误的原因
// 				fmt.Fprintf(os.Stderr, "dup2: %v\n", err) //打印错误信息,%v可以使用默认格式显示任意类型的值。os.Stderr是标准错误流
// 				continue
// 			}
// 			countLines(f, counts)
// 			f.Close() //关闭文件，释放相应资源
// 		}
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// func countLines(f *os.File, counts map[string]int) { //map被传给函数的时候，接收的是引用的副本，在这里修改的值可以在函数调用者中显示出来
// 	input := bufio.NewScanner(f) //bufio.NewScanner是按行读取数据
// 	for input.Scan() {
// 		counts[input.Text()]++ //记录字符串出现的次数
// 	}
// }

// package main

// import (
// 	"flag"
// 	"fmt"
// 	"strings"
// )

// // 命令行标志
// var n = flag.Bool("n", false, "omit trailing newline")
// var sep = flag.String("s", " ", "separator")

// func main() {
// 	flag.Parse() //用于解析命令行参数，将命令行中传入的参数与定义的标志进行匹配
// 	fmt.Println(flag.Args())
// 	fmt.Print(strings.Join(flag.Args(), *sep)) //flag.Args返回解析后的非标志参数，strings.Join会将所有参数连接起来，连接时使用定义的分隔符*sep
// 	if !*n {
// 		fmt.Println()
// 	}
// }

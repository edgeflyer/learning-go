// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	var s, sep string                   //var声明了两个string类型的变量s和sep，没有明确初始化的时候会隐式的初始化为这个类型的空值
// 	for i := 1; i < len(os.Args); i++ { //:用于对段变量的声明，这种语句声明一个或多个变量，并根据初始化的值给予合适的类型，i++在这里是语句，不是c语言的表达式，仅支持后缀
// 		s += sep + os.Args[i] //对于数字提供逻辑运算，对于字符串+提供追加操作
// 		sep = ""
// 	}
// 	fmt.Println(s)
// }

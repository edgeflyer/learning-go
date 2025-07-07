// package main

// import "fmt"

// func solution(str []string) []string {
// 	if len(str) == 0 {
// 		return str
// 	}
// 	i := 1
// 	for j := 1; j < len(str); j++ {
// 		if str[j] != str[j-1] {
// 			str[i] = str[j]
// 			i++
// 		}
// 	}
// 	return str[:i]
// }

// func main() {
// 	str := []string{"hello", "hello", "world"}
// 	fmt.Print(solution(str))
// }
// go中可以直接使用==来比较字符串是否完全一样

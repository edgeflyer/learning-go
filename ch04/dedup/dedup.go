// package main

// import "fmt"

// func main() {
// 	// seen := make(map[string]bool) // 字符串集合
// 	// input := bufio.NewScanner(os.Stdin)
// 	// for input.Scan() {
// 	// 	line := input.Text()
// 	// 	if !seen[line] {
// 	// 		seen[line] = true
// 	// 		fmt.Println(line)
// 	// 	}
// 	// }
// 	// if err := input.Err(); err != nil {
// 	// 	fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
// 	// 	os.Exit(1)
// 	// }
// 	var m = make(map[string]int)
// 	list := []string{"hello", "world", "this"}
// 	m[k(list)]++
// 	fmt.Printf("类型：%T, value: %v", k(list), k(list))
// 	// for k, v := range m {
// 	// 	fmt.Printf("%v,,, %v", k, v)
// 	// }

// }

// func k(list []string) string { return fmt.Sprintf("%q", list) }

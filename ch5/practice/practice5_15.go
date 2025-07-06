// package main

// import "fmt"

// type Number interface {
// 	~int | ~int64 | float64
// }

// func sum[T Number](nums ...T) T {
// 	var total T
// 	for _, n := range nums {
// 		total += n
// 	}
// 	return total
// }

// func max[T Number](nums ...T) T {
// 	if len(nums) == 0 {
// 		panic("max: 至少需要一个参数") // 抛出一个运行时错误
// 	}
// 	m := nums[0]
// 	for _, n := range nums {
// 		if n > m {
// 			m = n
// 		}
// 	}
// 	return m
// }

// func min[T Number](nums ...T) T {
// 	if len(nums) == 0 {
// 		panic("min: 至少需要一个参数")
// 	}
// 	m := nums[0]
// 	for _, n := range nums[1:] {
// 		if n < m {
// 			m = n
// 		}
// 	}
// 	return m
// }

// func Join(sep string, elems ...string) string {
// 	result := ""
// 	for i, s := range elems {
// 		if i > 0 {
// 			result += sep
// 		}
// 		result += s
// 	}
// 	return result
// }

// func main() {
// 	fmt.Println(Join("-", "a", "b", "c"))
// }

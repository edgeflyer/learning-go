// //演示slice的就地修改算法
// package main

// // nonempty返回一个新的slice，slice中的元素都是非空字符串
// // 在函数的调用过程中，底层数组的元素发生了改变
// func nonempty(strings []string) []string {
// 	i := 0
// 	for _, s := range strings {
// 		if s != "" {
// 			strings[i] = s
// 			i++
// 		}
// 	}
// 	return strings[:i]
// }

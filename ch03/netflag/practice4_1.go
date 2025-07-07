// package main

// import (
// 	"ch3/popcount"
// 	"crypto/sha256"
// 	"fmt"
// )

// func main() {
// 	c1 := sha256.Sum256([]byte("Byx"))

// 	fmt.Printf("%x\n", c1)
// 	var sum = 0
// 	for _, val := range c1 {
// 		// fmt.Printf("%b\n", val)
// 		sum += popcount.PopCount(val)
// 	}
// 	fmt.Println(sum)
// }

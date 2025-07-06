// package main

// import (
// 	"fmt"
// 	"os"
// 	"time"
// )

// func main() {
// 	fmt.Println("Commencing countdown. Press return to abort.")

// 	//按回车键取消发射
// 	abort := make(chan struct{})
// 	go func() {
// 		os.Stdin.Read(make([]byte, 1)) // 读取单个字符
// 		abort <- struct{}{}
// 	}()

// 	// tick := time.Tick(1 * time.Second)
// 	// for countdown := 10; countdown > 0; countdown-- {
// 	// 	fmt.Println(countdown)
// 	// 	<-tick
// 	// }
// 	// launch()

// 	select {
// 	case <-time.After(0 * time.Second):
// 		wait := make(chan struct{})
// 		go func() {
// 			for count := 10; count > 0; count-- {
// 				fmt.Println(count)
// 				time.Sleep(1 * time.Second)
// 			}
// 			wait <- struct{}{}
// 			launch()
// 		}()
// 	case <-abort:
// 		fmt.Print("launch aborted!")
// 		return
// 	}
// }

// func launch() (int, error) {
// 	return fmt.Print("launch finished!")
// }

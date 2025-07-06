package main

import (
	"fmt"
	"log"
	"time"
)

// 先调用bigSlowOperation之后defer调用trace打印出来开始执行的时间后返回一个函数(函数入口时间)，defer把这个函数放到栈里面，执行完sleep之后执行返回的函数(函数出口时间)
func bigSlowOperation() {
	defer trace("bigSlowoperation")() // 为什么必须有这个括号:defer需要具体执行的操作，trace返回的是一个函数，所以需要()去执行函数
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func main() {
	// bigSlowOperaion()
	_ = double(4)
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	x++
	return x + x
}

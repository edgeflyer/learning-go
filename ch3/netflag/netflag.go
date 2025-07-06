// package main

// import "fmt"

// type Flags uint

// const (
// 	FlagUp           Flags = 1 << iota //向上 00000001
// 	FlagBroadcast                      //支持广播访问 00000010
// 	FlagLoopback                       // 是环回接口 00000100
// 	FlagPointToPoint                   // 属于点对点的链路 00001000
// 	FlagMulticast                      // 支持多路广播访问 00010000

// )

// func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
// func TurnDown(v *Flags)     { *v &^= FlagUp } // 按位清除，如果FlagUp的值为一，按位清楚v对应位置，为0保持不变
// func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
// func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

// func main() {
// 	var v Flags = FlagMulticast | FlagUp // 10001
// 	fmt.Printf("%b %t\n", v, IsUp(v))    // "10001 true"
// 	TurnDown(&v)
// 	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
// 	SetBroadcast(&v)
// 	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
// 	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
// }

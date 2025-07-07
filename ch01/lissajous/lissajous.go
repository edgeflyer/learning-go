// package main

// //导入包的时候可以使用路径最后一段来引用这个包
// import (
// 	"image"
// 	"image/color"
// 	"image/gif"
// 	"io"
// 	"log"
// 	"math"
// 	"math/rand"
// 	"net/http"
// 	"os"
// 	"time"
// )

// // var声明的变量作用于整个包，声明时可以指定类型或者根据初始值自动推导类型
// // 与:=相比，:=是简短声明，常用于函数内部，结合了声明和赋值两种操作，变量是局部的，且不能已经被声明
// var palette = []color.Color{
// 	color.Black,
// 	color.RGBA{0, 255, 0, 255},
// } //[]color.Color{}是复合字面量，用一些列元素的值初始化go的复合类型的紧凑表答方式

// const ( //用来给常量命名，其值在编译期间固定
// 	blackIndex = 0 //画板中的一个颜色
// 	greenIndex = 1 //画板中的下一个颜色
// )

// func main() {
// 	rand.Seed(time.Now().UTC().UnixNano())
// 	if len(os.Args) > 1 && os.Args[1] == "web" {
// 		//匿名函数，定义了一个HTTP请求处理函数
// 		handler := func(w http.ResponseWriter, r *http.Request) { //w代表HTTP响应入口的写入接口，r代表HTTP请求的内容
// 			lissajous(w)
// 		}
// 		http.HandleFunc("/", handler)                         //将http请求路径/绑定到handler函数，作用是当web服务器接到对/的请求时，会调用handler来进行处理
// 		log.Fatal(http.ListenAndServe("localhost:8080", nil)) //启动服务器并监听请求，如果报错log.Fatal会记录报错并终止程序。http.ListenAndServe用于启动并监听服务器，第一个参数是地址和端口（可以使用:8080监听所有IP地址的8080端口；第二个参数是http.ServeMux或其他的http.Handler，这里用nil表示使用默认的路由器http.DefaultServeMux，它会自动处理由 http.HandleFunc 注册的路由。
// 		return
// 	}
// 	lissajous(os.Stdout)
// }

// func lissajous(out io.Writer) {
// 	const (
// 		cycles  = 5
// 		res     = 0.001
// 		size    = 100
// 		nframes = 64
// 		delay   = 8
// 	)
// 	freq := rand.Float64() * 3.0
// 	anim := gif.GIF{LoopCount: nframes}
// 	phase := 0.0
// 	for i := 0; i < nframes; i++ {
// 		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
// 		img := image.NewPaletted(rect, palette)
// 		for t := 0.0; t < cycles*2*math.Pi; t += res {
// 			x := math.Sin(t)
// 			y := math.Sin(t*freq + phase)
// 			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
// 				greenIndex) //将颜色设置为绿色
// 		}
// 		phase += 0.1
// 		anim.Delay = append(anim.Delay, delay)
// 		anim.Image = append(anim.Image, img)
// 	}
// 	gif.EncodeAll(out, &anim)
// }

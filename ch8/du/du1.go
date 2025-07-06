package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// walkDir递归的遍历以dir为根目录的整个文件树
// 并在fileSizes上发送每个已找到的文件大小
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name()) // 构建子目录路径
			walkDir(subdir, fileSizes)                 // 递归调用walkDir，继续遍历子目录
		} else {
			fileSizes <- entry.Size() // 将文件的大小发到fileSizes通道
		}
	}
}

// dirents返回dir目录中的条目
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir) // 返回一个os.FileInof类型的slice，获得指定目录下的所有条目
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err) // 如果出错，打印错误信息
		return nil
	}
	return entries // 返回目录中的所有条目
}

func main() {
	//确定初始目录
	flag.Parse()         // 解析命令行中的标志，这个函数会处理通过命令行传递的参数
	roots := flag.Args() // 获取命令行中传递的额外参数，通常是目录路径，flag.Args返回一个字符串切片，包含所有的命令行参数（不包括标志部分）
	if len(roots) == 0 { // 如果没有路径，就当默认路径.
		roots = []string{"."}
	}

	// 遍历文件树
	fileSizes := make(chan int64) // 创建通道，用于传输文件大小
	go func() {                   // 匿名goroutine遍历目录
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes) // 所有目录遍历完，关闭通道，close通知接收端，文件大小的传输已完成
	}()
	// 输出结果
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

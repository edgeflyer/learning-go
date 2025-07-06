package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s) // 把字符串解析成为time.Duration类型，表示的是时间长度，单位通常是ns
	if err != nil {
		panic(s) // 打印无法解析的字符串
	}
	return d
}

// printTracks函数将播放列表输出为一个表格。
func printTracks(Track []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // 计算各列宽度并输出表格
}

// type byArtist []*Track

// func (x byArtist) Len() int           { return len(x) }
// func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
// func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// func main() {
// 	sort.Sort(byArtist(tracks))
// 	printTracks(tracks)
// }

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) } // 实现sort.Interface接口的Less方法，用的是定义的结构体中的less
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	sort.Sort(customSort{tracks, func(x, y *Track) bool { // 实现结构体中的less，同时实现了Less方法
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)
}

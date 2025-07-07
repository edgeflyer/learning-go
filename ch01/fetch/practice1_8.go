// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"strings"
// )

// func main() {
// 	prefix := "http://"

// 	for _, url := range os.Args[1:] {
// 		if !strings.HasPrefix(url, prefix) {

// 			fmt.Println("The Url is needed to add prefix")
// 			fmt.Println("Old url is: ", url)
// 			url = prefix + url
// 			fmt.Println("New url is: ", url)
// 		}

// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fatch: %v\n", err)
// 			os.Exit(1)
// 		}
// 		defer resp.Body.Close()

// 		fmt.Println("Enter filename:")
// 		input := bufio.NewScanner(os.Stdin)
// 		input.Scan()
// 		addr := input.Text()

// 		file, err := os.Create(addr)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
// 			os.Exit(1)
// 		}
// 		defer file.Close()

// 		n, err := io.Copy(file, resp.Body)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Writing filed: %v\n", err)
// 			os.Exit(1)
// 		}
// 		fmt.Println("Writing file successfully!")
// 		fmt.Printf("Downloaded %d bytes to file %s\n", n, addr)
// 	}

// }

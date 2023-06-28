package main

import (
	"flag"
	"fmt"
	"strings"
	"thinknetika_go/pkg/crawler/spider"
)

func main() {
	fmt.Println("start")
	var sFlag = flag.String("s", "", "filter message for crawler")
	flag.Parse()
	var url1 string = "https://go.dev"
	var url2 string = "https://golang.org"
	var depth int8 = 1

	s := spider.New()
	result_godev, err := s.Scan(url1, int(depth))
	if err != nil {
		fmt.Println(err)
	}
	result_golang, err := s.Scan(url2, int(depth))
	if err != nil {
		fmt.Println(err)
	}
	result_total := append(result_godev, result_golang...)
	for _, v := range result_total {
		if strings.Contains(v.URL, *sFlag) {
			fmt.Println(v.URL)
		}
	}

	fmt.Println("done")
}

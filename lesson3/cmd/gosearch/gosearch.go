package main

import (
	"flag"
	"fmt"
	"sort"
	"thinknetika_go/lesson2/pkg/crawler"
	"thinknetika_go/lesson2/pkg/crawler/spider"
	"thinknetika_go/lesson3/pkg/index"
)

func main() {
	fmt.Println("start")
	var sFlag = flag.String("s", "", "filter message for crawler")
	flag.Parse()
	urls := []string{"https://go.dev", "https://habr.com/"}
	var depth int = 2
	var total = []crawler.Document{}
	var docNum int = 0

	s := spider.New()
	for _, url := range urls {
		scanResult, err := s.Scan(url, int(depth))
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, v := range scanResult {
			total = append(total,
				crawler.Document{
					ID:    docNum,
					URL:   v.URL,
					Title: v.Title,
					Body:  v.Body,
				},
			)
			docNum = docNum + 1
		}
	}

	for _, v := range total {
		index.Add(v)
	}
	docIds := index.Index[*sFlag]
	for _, dId := range docIds {
		document := sort.Search(len(total), func(ind int) bool { return total[ind].ID >= dId })
		fmt.Println(total[document].URL, total[document].Title)
	}
	// index.ShowIndex()
	fmt.Println("done")
}

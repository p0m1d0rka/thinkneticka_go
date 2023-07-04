package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"
	"thinknetika_go/lesson2/pkg/crawler"
	"thinknetika_go/lesson2/pkg/crawler/spider"
	"thinknetika_go/lesson3/pkg/index"
)

func main() {
	fmt.Println("start")
	var sFlag = flag.String("s", "", "filter message for crawler")
	flag.Parse()
	urls := []string{"https://go.dev", "https://habr.com/"}
	var depth int8 = 2
	var total = []crawler.Document{}
	var currentDocumentNumber int = 0

	s := spider.New()
	for _, url := range urls {
		scanResult, err := s.Scan(url, int(depth))
		if err != nil {
			fmt.Println(err)
		}

		for _, v := range scanResult {
			total = append(total,
				crawler.Document{
					ID:    currentDocumentNumber,
					URL:   v.URL,
					Title: v.Title,
					Body:  v.Body,
				},
			)
			currentDocumentNumber = currentDocumentNumber + 1
		}
	}

	for _, v := range total {
		index.Indexing(v)
		if strings.Contains(v.URL, *sFlag) {
			fmt.Println(v.URL)
		}
	}
	docIds := index.Index[*sFlag]
	for _, dId := range docIds {
		document := sort.Search(len(total), func(ind int) bool { return total[ind].ID >= dId })
		fmt.Println(total[document].URL, total[document].Title)
	}
	// index.ShowIndex()
	fmt.Println("done")
}

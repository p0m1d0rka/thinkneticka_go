package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"thinknetika_go/lesson5/pkg/crawler"
	"thinknetika_go/lesson5/pkg/crawler/spider"
	"thinknetika_go/lesson5/pkg/index"
)

func saveDocsToFile(docs []crawler.Document, dataFileName string) error {
	docJsoned, err := json.Marshal(docs)
	if err != nil {
		fmt.Printf("Error while marshaling document")
		return errors.New("Error while marshaling document")
	}
	err = os.WriteFile(dataFileName, docJsoned, 0644)
	if err != nil {
		fmt.Printf("Error while write file")
		return errors.New("Error while write file")
	}
	return nil
}

func loadDocsFromFile(dataFileName string) ([]crawler.Document, error) {
	data, err := os.ReadFile(dataFileName)
	var total []crawler.Document
	if err != nil {
		fmt.Println("Error while reading documents from file")
		return []crawler.Document{}, errors.New("Error while reading documents from file")
	}
	e := json.Unmarshal(data, &total)
	if e != nil {
		fmt.Println("Error while unmarshaling")
		return []crawler.Document{}, errors.New("Error while unmarshaling")
	}
	return total, nil

}
func main() {
	fmt.Println("start")
	var sFlag = flag.String("s", "", "filter message for crawler")
	flag.Parse()
	urls := []string{"https://go.dev", "https://habr.com/"}
	var depth int = 2
	var total = []crawler.Document{}
	var docNum int = 0
	dataFileName := "documents.json"

	total, e := loadDocsFromFile(dataFileName)
	if e != nil {
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
		// созраняем результат в файл
		err := saveDocsToFile(total, dataFileName)
		if err != nil {
			fmt.Println("error while saving documents to file")
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

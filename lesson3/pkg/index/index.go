package index

import (
	"fmt"
	"sort"
	"strings"
	"thinknetika_go/lesson2/pkg/crawler"
)

var Index = make(map[string][]int)

func Indexing(doc crawler.Document) {
	title_words := strings.Fields(doc.Title)
	for _, word := range title_words {
		Index[word] = append(Index[word], doc.ID)
		sort.Sort(sort.IntSlice(Index[word]))

	}
}

func ShowIndex() {
	fmt.Println(Index)
}

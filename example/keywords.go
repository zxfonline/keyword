package main

import (
	"fmt"
	"strings"

	"github.com/zxfonline/csvconfig"
	"github.com/zxfonline/keyword"
)

func main() {
	csvconfig.Init("", "")
	err1 := csvconfig.Load([]string{"badwords"})
	if err1 != nil {
		panic(err1)
	}
	recs := csvconfig.GetAll("badwords")
	words := make([]string, 0, len(recs))
	for _, rec := range recs {
		fields := rec.Fields
		word := fields["word"]
		words = append(words, word)
	}
	kword := keyword.NewKeyWord()
	kword.Init(words)
	keyword.SetGlobal(kword)
	bol := haveBadWords("fuck")
	fmt.Println(bol)
}

func haveBadWords(str string) bool {
	lower := strings.ToLower(str)
	return keyword.Search(lower)
}

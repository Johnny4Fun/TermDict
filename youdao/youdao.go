package youdao

import (
	"fmt"
	"github.com/quii/learn-go-with-tests/mytest/TermDict"
	"log"
	"net/http"
)

type Dict struct {
}
func NewDict() *Dict{
	return &Dict{}
}
//
const queryURLPattern = "http://www.youdao.com/w/%s"

var pageParser = NewParser()

func (y *Dict) LookUp(word string) *TermDict.Word {
	url := fmt.Sprintf(queryURLPattern, word)
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		log.Fatalf("GET %q Error, %v", url, err)
		return nil
	}
	defer resp.Body.Close()

	wd, err := pageParser.Parse(resp.Body)
	if err != nil {
		log.Fatalf("parse resp body error, %v", err)
	}

	wd.Spell = word

	return wd
}
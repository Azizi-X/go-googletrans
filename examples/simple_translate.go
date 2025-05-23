//go:build ignore
// +build ignore

package main

import (
	"fmt"

	translator "github.com/dreamsxin/go-googletrans"
)

var content = `你好，世界！`

func main() {
	c := translator.TranslateConfig{
		//Proxy:       "http://127.0.0.1:1087",
		UserAgent:   []string{"Custom Agent"},
		ServiceUrls: []string{"translate.google.com.hk"},
	}
	t := translator.New(c)
	result, err := t.Translate(content, "auto", "en")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Text)
}

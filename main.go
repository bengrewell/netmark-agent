package main

import (
	"fmt"
	"github.com/hokaccha/go-prettyjson"
)

func main() {
	info := BasicInfo{}
	info.Update()

	b, err := prettyjson.Marshal(info)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(b))
}

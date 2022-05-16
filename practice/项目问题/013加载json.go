package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type JsonStruct struct {
	Appid  string   `json:"appid"`
	Secret string   `json:"secret"`
	Psas   []string `json:"psas"`
}

func main() {
	f, err := os.Open("D:\\goitem\\go-box\\practice\\项目问题\\013.json")
	if err != nil {
		fmt.Println(err)
	}

	r := io.Reader(f)
	var ret []JsonStruct

	if err = json.NewDecoder(r).Decode(&ret); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}
}

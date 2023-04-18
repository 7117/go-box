package main

import (
	"fmt"
	"reflect"
	"time"
)

type resumea struct {
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
}

func findDoca(student []resumea) []map[string]interface{} {

	var docs []map[string]interface{}

	for k, v := range student {
		typeOf := reflect.TypeOf(v)
		ValOf := reflect.ValueOf(v)

		doc := make(map[string]interface{})
		for i := 0; i < typeOf.NumField(); i++ {

			if typeOf.Field(i).Tag.Get("json") == "create_time" {
				CreateTimeVal := student[k].CreateTime.Format("2006-01-02 15:04:05")
				doc[typeOf.Field(i).Tag.Get("json")] = CreateTimeVal
				continue
			}

			doc[typeOf.Field(i).Tag.Get("json")] = ValOf.Field(i)
		}
		docs = append(docs, doc)
	}
	return docs
}

func main() {
	var student []resumea
	student = make([]resumea, 1)
	student[0].Name = "sundial"
	student[0].CreateTime = time.Now()
	fmt.Println("A", student)
	doc := findDoca(student)
	fmt.Println("END", doc)

}

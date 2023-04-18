package main

import (
	"fmt"
	"reflect"
	"time"
)

type resume struct {
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
}

func findDoc(student resume) map[string]interface{} {
	typeOf := reflect.TypeOf(student)
	ValOf := reflect.ValueOf(student)
	doc := make(map[string]interface{})

	for i := 0; i < typeOf.NumField(); i++ {

		if typeOf.Field(i).Tag.Get("json") == "create_time" {
			CreateTimeVal := student.CreateTime.Format("2006-01-02 15:04:05")
			doc[typeOf.Field(i).Tag.Get("json")] = CreateTimeVal
			continue
		}

		doc[typeOf.Field(i).Tag.Get("json")] = ValOf.Field(i)
	}

	return doc

}

func main() {
	var student resume
	student.Name = "sundial"
	student.CreateTime = time.Now()
	fmt.Println("A", student)
	doc := findDoc(student)
	fmt.Println("END", doc)

}

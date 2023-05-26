package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map

	//增加
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	//删除
	scene.Delete("london")
	//查询
	fmt.Println(scene.Load("london"))
	val, boolVal := scene.LoadAndDelete("greece")
	fmt.Println("LoadAndDelete:", val, boolVal)
	fmt.Println(scene.Load("greece"))
	scene.LoadOrStore("greece", "aaaa")
	fmt.Println(scene.Load("greece"))
	fmt.Println("LoadAndDelete:")

	// 遍历
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

	count := GetSyncMapCount(&scene)
	fmt.Println("长度", count)
	var sceneLen sync.Map
	countLen := GetSyncMapCount(&sceneLen)
	fmt.Println("长度", countLen)
}

func GetSyncMapCount(scene *sync.Map) int64 {
	count := 0
	scene.Range(func(k, v interface{}) bool {
		count++
		fmt.Println("iterate:", k, v)
		return true
	})
	return int64(count)
}

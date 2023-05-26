package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
)

func main() {
	普通()
	遍历和长度获取()
	序列化反序列化()
	GetShardUse()
}

func 普通() {
	m := cmap.New[string]()

	m.Set("foo", "bar")

	l := make(map[string]string)
	l["a"] = "a"
	l["b"] = "b"
	m.MSet(l)
	fmt.Println("m.MSet(l):", m.Items())

	bar, ok := m.Get("foo")
	fmt.Println(bar, ok)

	m.Remove("foo")
}

func 遍历和长度获取() {
	m := cmap.New[string]()

	m.Set("foo2", "bar2")
	m.Set("foo3", "bar3")
	m.Set("foo4", "bar4")
	fmt.Println(m.Count())

	keys := m.Keys()
	for _, v := range keys {
		val, ok := m.Get(v)
		fmt.Println(val, ok)
	}

	fmt.Println("开始使用items")

	for k, v := range m.Items() {
		fmt.Println(k, v)
	}

	//m.Clear()

	fmt.Println("m.IsEmpty()", m.IsEmpty())
	fmt.Println("m.Has(\"foo2\")", m.Has("foo2"))
	fmt.Println("m.IterCb", m.IterCb)
	a1, a2 := m.Pop("foo2")
	fmt.Println("m.Pop(\"foo2\")", a1, a2)

}

func 序列化反序列化() {

	m := cmap.New[string]()
	m.Set("foo1", "bar1")
	m.Set("foo2", "bar2")
	json, err := m.MarshalJSON()
	fmt.Println("序列化", json, err)
	err1 := m.UnmarshalJSON(json)
	fmt.Println("反序列化", err1, m.Items())

}

func GetShardUse() {
	m := cmap.New[string]()
	l := make(map[string]string)
	l["a"] = "a"
	l["b"] = "b"
	m.MSet(l)

	fmt.Println("m.MSet(l)", m.Items())
	fmt.Println("m.MSet(l)", m.Items())

	m.IterBuffered()

}

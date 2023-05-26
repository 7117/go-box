package main

//sync.Map在读多写少性能比较好，而concurrent-map 在key的hash度高的情况下性能比较好。
import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"sort"
	"strconv"
)

type Animal struct {
	name string
}

func main() {
	TestRemoveCb()
	TestIterator()
	TestBufferedIterator()
	TestIterCb()
	TestConcurrent()
	TestUpsert()
}

// TestRemoveCb
//此处函数的作用就是：删除的时候执行一个回调函数
//回调函数的参数是：要删除元素的 key val 是否存在的key
func TestRemoveCb() {
	fmt.Println("TestRemoveCb")

	m := cmap.New[Animal]()
	monkey := Animal{"monkey"}
	elephant := Animal{"elephant"}
	m.Set("monkey", monkey)
	m.Set("elephant", elephant)
	fmt.Println("TestRemoveCb", m.Items())

	var (
		mapKey   string
		mapVal   Animal
		wasFound bool
	)
	cb := func(key string, val Animal, exists bool) bool {
		fmt.Println("TestRemoveCb", key, val, exists)
		mapKey = key
		mapVal = val
		wasFound = exists
		return wasFound
	}

	result := m.RemoveCb("monkey", cb)
	fmt.Println("TestRemoveCb", result, mapKey, mapVal)
}

//TestIterator
//此处的函数作用是进行遍历循环，已经废弃了，使用新的items即可
func TestIterator() {
	fmt.Println("TestIterator")

	m := cmap.New[Animal]()

	for i := 0; i < 100; i++ {
		m.Set(strconv.Itoa(i), Animal{strconv.Itoa(i)})
	}
	fmt.Println("TestIterator", m.Items())

	counter := 0
	for k, item := range m.Items() {

		if (item == Animal{}) {
			fmt.Println("TestIterator", "Expecting an object.", k)
		}
		counter++
	}

	if counter != 100 {
		fmt.Println("TestIterator", "We should have counted 100 elements.")
	}
}

func TestBufferedIterator() {
	fmt.Println("TestBufferedIterator")

	m := cmap.New[Animal]()

	for i := 0; i < 100; i++ {
		m.Set(strconv.Itoa(i), Animal{strconv.Itoa(i)})
	}
	fmt.Println("TestBufferedIterator", m.IterBuffered())

	counter := 0
	for item := range m.IterBuffered() {
		val := item.Val

		if (val == Animal{}) {
			fmt.Println("TestBufferedIterator", "Expecting an object.")
		}
		counter++
	}

	if counter != 100 {
		fmt.Println("TestBufferedIterator", "We should have counted 100 elements.")
	}
}

func TestIterCb() {
	fmt.Println("TestIterCb")

	m := cmap.New[Animal]()

	// Insert 100 elements.
	for i := 0; i < 100; i++ {
		m.Set(strconv.Itoa(i), Animal{strconv.Itoa(i)})
	}

	counter := 0
	// Iterate over elements.
	m.IterCb(func(key string, v Animal) {
		counter++
	})
	fmt.Println("TestIterCb", counter)
}

func TestConcurrent() {
	m := cmap.New[int]()
	ch := make(chan int)
	const iterations = 1000
	var a [iterations]int

	go func() {
		for i := 0; i < iterations/2; i++ {
			m.Set(strconv.Itoa(i), i)
			val, _ := m.Get(strconv.Itoa(i))
			ch <- val
		}
	}()

	go func() {
		for i := iterations / 2; i < iterations; i++ {
			m.Set(strconv.Itoa(i), i)
			val, _ := m.Get(strconv.Itoa(i))
			ch <- val
		}
	}()
	fmt.Println("TestConcurrent", m.Items())

	counter := 0
	for elem := range ch {
		a[counter] = elem
		counter++
		if counter == iterations {
			break
		}
	}

	sort.Ints(a[0:iterations])

	if m.Count() != iterations {
		fmt.Println("TestConcurrent", "Expecting 1000 elements.")
	}

	for i := 0; i < iterations; i++ {
		if i != a[i] {
			fmt.Println("missing value", i)

		}
	}
}

func TestUpsert() {
	fmt.Println("TestUpsert")

	dolphin := Animal{"dolphin"}
	whale := Animal{"whale"}

	cb := func(exists bool, valueInMap Animal, newValue Animal) Animal {
		if !exists {
			return newValue
		}
		fmt.Println("TestUpsert1", newValue)
		valueInMap.name += newValue.name
		return valueInMap
	}

	m := cmap.New[Animal]()
	//TestUpsert map[marine:{dolphin}]
	//TestUpsert map[marine:{dolphinwhale}]
	m.Set("marine", dolphin)
	fmt.Println("TestUpsert", m.Items())
	m.Upsert("marine", whale, cb)
	fmt.Println("TestUpsert", m.Items())

	//m.Upsert("marine", whale, cb)
	//fmt.Println("TestUpsert", m.Items())
}

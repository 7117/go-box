package main

import (
	"testing"
	"time"
)

func Foreach(i uint64) uint64 {
	i++
	return i
}

func TestForeach10(t *testing.T) {

	var timer *time.Timer
	var i, tag uint64
	timer = time.NewTimer(1 * time.Second)
	go func() {
		<-timer.C
		tag = 1
	}()

	for {
		i = Foreach(i)
		if tag == 1 {
			t.Log("QPS:", i)
			break
		}
	}
}
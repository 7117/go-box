package goconvey

import (
	. "bou.ke/monkey"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAdd(t *testing.T) {
	Convey("1", t, func() {
		So(Add(1, 2), ShouldEqual, 3)
	})

	Convey("2", t, func() {
		//SkipSo函数表明相应的断言将不被执行
		SkipSo(Add(1, 2), ShouldEqual, 3)
	})

	//SkipConvey函数表明相应的闭包函数将不被执行
	SkipConvey("3", t, func() {
		So(Add(1, 2), ShouldEqual, 3)
	})

	Convey("4", t, func() {
		res := 8
		guard := Patch(Multi, func(_ int, _ int) int {
			return res
		})
		defer guard.Unpatch()
		So(Multi(3, 2), ShouldEqual, res)
	})
}

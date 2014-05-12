package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Check_New(t *testing.T) {
	HeadConvey("echo.websocket.org", t, func() {
		So(check("ws://echo.websocket.org"), ShouldEqual, nil)
	})
	HeadConvey("echo.websocket.org ssl", t, func() {
		So(check("wss://echo.websocket.org"), ShouldEqual, nil)
	})
}

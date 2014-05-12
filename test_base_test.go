package main

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
)

func HeadConvey(items ...interface{}) {
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("-- " + items[0].(string))
	fmt.Println("-------------------------------------------------------------------------")
	Convey(items...)
}

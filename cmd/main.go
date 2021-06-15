package main

import (
	"fmt"

	lk "github.com/digisan/logkit"
)

func main() {
	lk.Log2F(true, "temp")
	lk.Log("%s", "this is from INFO log")
	lk.Debug("%s", "this is from DEBUG log")
	lk.WarnOnErr("%s", "this is from WarnOnErr log")
	lk.WarnDetail(false)
	lk.Warn("%s", "this is from WARN without detail log")
	lk.WarnDetail(true)
	lk.Warn("%s", "this is from WARN with detail log")
	lk.FailOnErr("%s%v", "this is from FailOnErr log", fmt.Errorf(""))
}

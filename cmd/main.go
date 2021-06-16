package main

import (
	"fmt"

	lk "github.com/digisan/logkit"
)

func main() {

	lk.Log2F(true, "temp")

	iinfo := lk.Fac4IdxLog(0, lk.INFO, false)
	iinfo("%s", "this is from INFO log")
	iinfo("%s", "this is from INFO log")
	iinfo("%s", "this is from INFO log")

	idebug := lk.Fac4IdxLog(0, lk.DEBUG, false)
	idebug("%s", "this is from DEBUG log")
	idebug("%s", "this is from DEBUG log")
	idebug("%s", "this is from DEBUG log")

	iwarn := lk.Fac4IdxLog(0, lk.WARN, false)
	iwarn("%s", "this is from WARN log")
	iwarn("%s", "this is from WARN log")
	iwarn("%s", "this is from WARN log")

	ifail := lk.Fac4IdxLog(0, lk.FAIL, false)
	ifail("%s", "this is from FAIL log")
	ifail("%s", "this is from FAIL log")
	ifail("%s", "this is from FAIL log")

	fmt.Println("----------------------------------------")
	// ------------------------------ //

	lk.Log("%s %d", "this is from INFO log", 100)
	lk.Debug("%s", "this is from DEBUG log")
	lk.WarnOnErr("%s", "this is from WarnOnErr log")
	lk.WarnDetail(false)
	lk.Warn("%s", "this is from WARN without detail log")
	lk.WarnDetail(true)
	lk.Warn("%s", "this is from WARN with detail log")
	lk.FailOnErr("%s%v", "this is from FailOnErr log", fmt.Errorf(""))
}

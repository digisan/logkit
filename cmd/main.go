package main

import (
	"fmt"

	lk "github.com/digisan/logkit"
	"github.com/digisan/gotk/project"
)

func main() {

	pn, ok := project.PrjName("TEST-PNAME")
	lk.Log("%v - %v", pn, ok)

	pv, ok := project.GitVer("TEST-PVERSION")
	lk.Log("%v - %v", pv, ok)

	lk.Log("%s", "-------------------------")

	lk.Log2F(true, "temp")

	iinfo := lk.Fac4GrpIdxLogF("A", 0, lk.INFO, false)
	iinfo("%s %05d", "this is from INFO log", 100)
	iinfo("%s", "this is from INFO log")
	iinfo("%s", "this is from INFO log")

	idebug := lk.Fac4GrpIdxLogF("B", 0, lk.DEBUG, false)
	idebug("%s", "this is from DEBUG log")
	idebug("%s", "this is from DEBUG log")
	idebug("%s", "this is from DEBUG log")

	iwarn := lk.Fac4GrpIdxLogF("C", 0, lk.WARN, false)
	iwarn("%s", "this is from WARN log")
	iwarn("%s", "this is from WARN log")
	iwarn("%s", "this is from WARN log")

	ifail := lk.Fac4GrpIdxLogF("D", -1, lk.FAIL, false)
	ifail("%s", "this is from FAIL log")
	ifail("%s%v", "this is from FAIL log", fmt.Errorf(" --- STOP"))
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

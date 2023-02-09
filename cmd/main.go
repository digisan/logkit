package main

import (
	"fmt"

	"github.com/digisan/gotk/project"
	lk "github.com/digisan/logkit"
)

func main() {

	pn, ok := project.PrjName("TEST-PNAME")
	lk.Log("%v - %v", pn, ok)

	pv, ok := project.GitVer("TEST-PVERSION")
	lk.Log("%v - %v", pv, ok)

	lk.Log("%s", "-------------------------")

	lk.Log2F(true, "temp")

	lk.Log("this is from INFO log"+lk.LF+"this is from INFO log 2nd line %d", 100)

	iinfo := lk.Fac4GrpIdxLogF("A", 0, lk.INFO, false)
	iinfo("this is from INFO log 1"+lk.LF+"this is a new line %05d"+lk.LF+"this is another new line", 100)
	iinfo("this is from INFO log 2")
	iinfo("this is from INFO log 3")

	idebug := lk.Fac4GrpIdxLogF("B", 0, lk.DEBUG, false)
	idebug("this is from DEBUG log 1")
	idebug("this is from DEBUG log 2"+lk.LF+"this is a new line %05d"+lk.LF+"this is another new line", 200)
	idebug("%s", "this is from DEBUG log 3")

	iwarn := lk.Fac4GrpIdxLogF("C", 0, lk.WARN, false)
	iwarn("this is from WARN log 1")
	iwarn("this is from WARN log 2"+lk.LF+"this is a new line %05d"+lk.LF+"this is another new line", 300)
	iwarn("%s", "this is from WARN log 3")

	ifail := lk.Fac4GrpIdxLogF("D", -1, lk.FAIL, false)
	ifail("this is from FAIL log 1")
	ifail("this is from FAIL log 2 %v"+lk.LF+"this is a new line %d"+lk.LF+"this is the third line %d", fmt.Errorf(" --- STOP"), 400, 401)
	ifail("%s", "this is from FAIL log 3")

	fmt.Println("----------------------------------------")

	lk.Log("%s %d", "this is from INFO log", 100)
	lk.Debug("%s", "this is from DEBUG log")
	lk.WarnOnErr("%s", "this is from WarnOnErr log")
	lk.WarnDetail(false)
	lk.Warn("%s", "this is from WARN without detail log")
	lk.WarnDetail(true)
	lk.Warn("%s", "this is from WARN with detail log")
	lk.FailOnErr("%s%v", "this is from FailOnErr log", fmt.Errorf(""))
}

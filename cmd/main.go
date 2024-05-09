package main

import (
	"fmt"
	"time"

	"github.com/digisan/gotk/project"
	lk "github.com/digisan/logkit"
)

func main() {

	defer lk.Track(time.Now())

	lk.Log2F(true, true, "temp")
	// lk.Log2C(false)

	pn, ok := project.PrjName("TEST-PNAME")
	lk.Log("%v - %v", pn, ok)

	pv, ok := project.GitVer("TEST-PVERSION")
	lk.Log("%v - %v", pv, ok)

	lk.Log("%s", "-------------------------")

	lk.Log("this is from INFO log"+lk.LF+"this is from INFO log 2nd line %d", 100)

	return

	info := lk.Fn4GrpIdxLogF("A", 0, lk.INFO, false)
	info("this is from INFO log 1"+lk.LF+"this is a new line %05d"+lk.LF+"this is another new line", 100)
	info("this is from INFO log 2")
	info("this is from INFO log 3")

	debug := lk.Fn4GrpIdxLogF("B", 0, lk.DEBUG, false)
	debug("this is from DEBUG log 1")
	debug("this is from DEBUG log 2"+lk.LF+"this is a new line %05d"+lk.LF+"this is another new line", 200)
	debug("%s", "this is from DEBUG log 3")

	warn := lk.Fn4GrpIdxLogF("C", 0, lk.WARN, false)
	warn("this is from WARN log 1")
	warn("this is from WARN log 2"+lk.LF+"this is a new line %05d"+lk.LF+"this is another new line", 300)
	warn("%s", "this is from WARN log 3")

	fail := lk.Fn4GrpIdxLogF("D", -1, lk.FAIL, false)
	fail("this is from FAIL log 1")
	fail("this is from FAIL log 2 %v"+lk.LF+"this is a new line %d"+lk.LF+"this is the third line %d", fmt.Errorf(" --- STOP"), 400, 401)
	fail("%s", "this is from FAIL log 3")

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

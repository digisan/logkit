package logkit

import (
	"fmt"
	"os"
	"strings"
	"time"

	. "github.com/digisan/gotk/print"
)

var (
	fPt        = fmt.Print
	fSf        = fmt.Sprintf
	fEf        = fmt.Errorf
	sHasSuffix = strings.HasSuffix
)

const (
	tmFmt        = "2006/01/02 15:04:05 " // end with " " same as log.Printf
	logfile4test = "./a/b.log"
	LF           = "\n\t\t\t\t"
	longLF       = "\n\t\t\t\t\t\t"
)

type logCategory int

const (
	FILE  logCategory = 0
	INFO  logCategory = 1
	DEBUG logCategory = 2
	WARN  logCategory = 3
	FAIL  logCategory = 4
)

var (
	mLvlDesc map[logCategory]string = map[logCategory]string{
		FILE:  "",
		INFO:  "INFO",
		DEBUG: "DEBUG",
		WARN:  "WARN",
		FAIL:  "FAIL",
	}

	mLvlClr map[logCategory]func(a ...any) string = map[logCategory]func(a ...any) string{
		FILE:  func(a ...any) string { return fmt.Sprint(a...) },
		INFO:  G, // W
		DEBUG: B,
		WARN:  Y,
		FAIL:  R,
	}

	log2C                          = true
	log2F                          = false
	warnDetail                     = true
	mPathFile  map[string]*os.File = make(map[string]*os.File)
	nowStr                         = func() string { return time.Now().Format(tmFmt) }
)

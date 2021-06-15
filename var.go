package logkit

import (
	"fmt"
	"os"
	"strings"
	"time"

	clr "github.com/gookit/color"
)

var (
	fPt        = fmt.Print
	fPln       = fmt.Println
	fSf        = fmt.Sprintf
	fEf        = fmt.Errorf
	sHasSuffix = strings.HasSuffix
	sLastIndex = strings.LastIndex
	// sReplaceAll  = strings.ReplaceAll
	// sJoin        = strings.Join
	// sSplit       = strings.Split
	// sContains    = strings.Contains
	// scParseFloat = strconv.ParseFloat

)

const (
	tmFmt        = "2006/01/02 15:04:05 " // end with " " same as log.Printf
	logfile4test = "./a/b.log"
)

type logLevel int

const (
	FILE  logLevel = 0
	INFO  logLevel = 1
	DEBUG logLevel = 2
	WARN  logLevel = 3
	FAIL  logLevel = 4
)

var (
	mLvlDesc map[logLevel]string = map[logLevel]string{
		FILE:  "",
		INFO:  "INFO",
		DEBUG: "DEBUG",
		WARN:  "WARN",
		FAIL:  "FAIL",
	}

	W = clr.FgWhite.Render  // file
	G = clr.FgGreen.Render  // log
	B = clr.FgBlue.Render   // debug
	Y = clr.FgYellow.Render // warn
	R = clr.FgRed.Render    // fail

	mLvlClr map[logLevel]func(a ...interface{}) string = map[logLevel]func(a ...interface{}) string{
		FILE:  func(a ...interface{}) string { return fmt.Sprint(a...) },
		INFO:  G, // W
		DEBUG: B,
		WARN:  Y,
		FAIL:  R,
	}

	log2C                          = true
	log2F                          = false
	warnDetail                     = true
	mPathFile  map[string]*os.File = make(map[string]*os.File)
	nowstr                         = func() string { return time.Now().Format(tmFmt) }
)

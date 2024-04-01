package logkit

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"

	. "github.com/digisan/go-generics"
	fd "github.com/digisan/gotk/file-dir"
	"github.com/digisan/gotk/track"
)

// logger : write info into Console OR File
func logger(tl int, cat logCategory, format string, v ...any) {

	tc := "\n" + track.CallerDescription(tl) + "\n"
	const tcPrefix = "\n\t\t\t\t--> "
	const tcPrefix4f = "\n\t\t\t\t\t\t\t\t--> "
	const tcPrefix4c = "\n\t\t\t\t\t\t--> "
	tc = strings.Replace(tc, "\n", tcPrefix, 2)

	ev := Filter(v, func(i int, e any) bool { _, ok := e.(error); return ok && e != nil })
	hasErr := len(ev) > 0

	clrDesc := mLvlClr[cat](mLvlDesc[cat])
	v4c := append([]any{clrDesc}, v...)

	clrDesc = mLvlClr[FILE](mLvlDesc[cat])
	v4f := append([]any{clrDesc}, v...)

	nLF := strings.Count(format, LF)
	const lf4f = "\"\n\t\t\t\t\t\t\t\t\""

	switch cat {
	case INFO:
		if log2C {
			item := fSf("\t%s\t"+format+"\n", v4c...)
			fPt(nowStr() + item)
		}
		if log2F {
			item := fSf("\t%s\t\""+format+"\"\n", v4f...)
			item = strings.Replace(item, LF, lf4f, nLF)
			log.Printf("%s", item)
		}

	case DEBUG:
		if log2C {
			item := fSf("\t%s\t"+format+"%s", append(v4c, B(tc))...)
			fPt(nowStr() + item)
		}
		if log2F {
			item := fSf("\t%s\t\""+format+"\"%s", append(v4f, tc)...)
			item = strings.Replace(item, LF, lf4f, nLF)
			item = strings.Replace(item, tcPrefix, tcPrefix4f, 2)
			log.Printf("%s", item)
		}

	case WARN:
		if hasErr {
			if !warnDetail {
				if log2C {
					item := fSf("\t%s\t"+format+"\n", v4c...)
					fPt(nowStr() + item)
				}
				if log2F {
					item := fSf("\t%s\t\""+format+"\"\n", v4f...)
					item = strings.Replace(item, LF, lf4f, nLF)
					log.Printf("%s", item)
				}
			} else {
				if log2C {
					item := fSf("\t%s\t"+format+"%s", append(v4c, Y(tc))...)
					fPt(nowStr() + item)
				}
				if log2F {
					item := fSf("\t%s\t\""+format+"\"%s", append(v4f, tc)...)
					item = strings.Replace(item, LF, lf4f, nLF)
					item = strings.Replace(item, tcPrefix, tcPrefix4f, 2)
					log.Printf("%s", item)
				}
			}
		}

	case FAIL:
		if hasErr {

			var item string

			switch {
			case log2C && log2F:
				// console
				item = fSf("\t%s\t"+format+"%s", append(v4c, R(tc))...)
				fPt(nowStr() + item)

				// file
				item = fSf("\t%s\t\""+format+"\"%s", append(v4f, tc)...)
				item = strings.Replace(item, LF, lf4f, nLF)
				item = strings.Replace(item, tcPrefix, tcPrefix4f, 2)
				log.Printf("%s", item)

			case log2C:
				item = fSf("\t%s\t"+format+"%s", append(v4c, R(tc))...)
				item = strings.Replace(item, LF, longLF, nLF)
				item = strings.Replace(item, tcPrefix, tcPrefix4c, 2)
				log.Printf("%s", item)

			case log2F:
				item = fSf("\t%s\t\""+format+"\"%s", append(v4f, tc)...)
				item = strings.Replace(item, LF, lf4f, nLF)
				item = strings.Replace(item, tcPrefix, tcPrefix4f, 2)
				log.Printf("%s", item)
			}

			/// *** record fatal stack message to file ***
			///

			fatalDir := "./fatal"
			fd.MustCreateDir(fatalDir)
			fName := strings.TrimSpace(strings.ReplaceAll(nowStr(), "/", "-"))
			fPath := filepath.Join(fatalDir, fName+".log")

			f, err := os.OpenFile(fPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
			if err != nil {
				panic(err)
			}

			multiWriter := io.MultiWriter(os.Stderr, f)
			rd, wr, err := os.Pipe()
			if err != nil {
				panic(err)
			}
			os.Stderr = wr

			go func() {
				scanner := bufio.NewScanner(rd)
				for scanner.Scan() {
					multiWriter.Write([]byte(scanner.Text() + "\n"))
				}
			}()

			fmt.Fprintln(os.Stderr, item+"\n")
			debug.PrintStack()
			time.Sleep(time.Duration(1 * time.Second))

			/// ***

			panic("FAILED!")
		}
	}
}

// ------------------------------------------------------- //

func Log(format string, v ...any) {
	logger(0, INFO, format, v...)
}

// LogWhen : write info into Console OR File
func LogWhen(condition bool, format string, v ...any) {
	if condition {
		Log(format, v...)
	}
}

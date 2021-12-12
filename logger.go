package logkit

import (
	"log"

	"github.com/digisan/go-generics/obj"
)

// logger : write info into Console OR File
func logger(tl int, lvl logLevel, format string, v ...interface{}) {

	tc := TrackCaller(tl)

	ev := obj.FM(v, func(i int, e interface{}) bool { _, ok := e.(error); return ok && e != nil }, nil)
	hasErr := len(ev) > 0

	clrDesc := mLvlClr[lvl](mLvlDesc[lvl])
	v4c := append([]interface{}{clrDesc}, v...)

	clrDesc = mLvlClr[FILE](mLvlDesc[lvl])
	v4f := append([]interface{}{clrDesc}, v...)

	switch lvl {
	case INFO:
		if log2C {
			item := fSf("\t%s \t\""+format+"\"\n", v4c...)
			fPt(nowstr() + item)
		}
		if log2F {
			item := fSf("\t%s \t\""+format+"\"\n", v4f...)
			log.Printf("%s", item)
		}

	case DEBUG:
		if log2C {
			item := fSf("\t%s \t\""+format+"\"%s", append(v4c, tc)...)
			fPt(nowstr() + item)
		}
		if log2F {
			item := fSf("\t%s \t\""+format+"\"%s", append(v4f, tc)...)
			log.Printf("%s", item)
		}

	case WARN:
		if hasErr {
			if !warnDetail {
				if log2C {
					item := fSf("\t%s \t\""+format+"\"\n", v4c...)
					fPt(nowstr() + item)
				}
				if log2F {
					item := fSf("\t%s \t\""+format+"\"\n", v4f...)
					log.Printf("%s", item)
				}
			} else {
				if log2C {
					item := fSf("\t%s \t\""+format+"\"%s", append(v4c, tc)...)
					fPt(nowstr() + item)
				}
				if log2F {
					item := fSf("\t%s \t\""+format+"\"%s", append(v4f, tc)...)
					log.Printf("%s", item)
				}
			}
		}

	case FAIL:
		if hasErr {
			if log2C && log2F {
				item := fSf("\t%s \t\""+format+"\"%s", append(v4c, tc)...)
				fPt(nowstr() + item)
				item = fSf("\t%s \t\""+format+"\"%s", append(v4f, tc)...)
				log.Fatalf("%s", item)
			}
			if log2C {
				item := fSf("\t%s \t\""+format+"\"%s", append(v4c, tc)...)
				log.Fatalf("%s", item)
			}
			if log2F {
				item := fSf("\t%s \t\""+format+"\"%s", append(v4f, tc)...)
				log.Fatalf("%s", item)
			}
		}
	}
}

// ------------------------------------------------------- //

func Log(format string, v ...interface{}) {
	logger(0, INFO, format, v...)
}

// LogWhen : write info into Console OR File
func LogWhen(condition bool, format string, v ...interface{}) {
	if condition {
		Log(format, v...)
	}
}

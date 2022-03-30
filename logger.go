package logkit

import (
	"log"

	. "github.com/digisan/go-generics/v2"
)

// logger : write info into Console OR File
func logger(tl int, lvl logLevel, format string, v ...any) {

	tc := TrackCaller(tl)

	ov := Filter(&v, func(i int, e any) bool { _, ok := e.(error); return ok && e != nil })
	hasErr := len(v) > 0

	clrDesc := mLvlClr[lvl](mLvlDesc[lvl])
	v4c := append([]any{clrDesc}, ov...)

	clrDesc = mLvlClr[FILE](mLvlDesc[lvl])
	v4f := append([]any{clrDesc}, ov...)

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

func Log(format string, v ...any) {
	logger(0, INFO, format, v...)
}

// LogWhen : write info into Console OR File
func LogWhen(condition bool, format string, v ...any) {
	if condition {
		Log(format, v...)
	}
}

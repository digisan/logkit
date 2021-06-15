package logkit

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/digisan/gotk/io"
)

// WarnDetail :
func WarnDetail(enable bool) {
	warnDetail = enable
}

// Log2C :
func Log2C(enable bool) {
	log2C = enable
}

// Log2F :
func Log2F(enable bool, logfile string) {
	if enable {
		setLog(logfile)
	} else {
		resetLog()
	}
}

// FilePerm :
const FilePerm = 0666

// setLog :
func setLog(logfile string) {
	zone, offset := time.Now().Zone()
	if sHasSuffix(logfile, ".log") {
		logfile = logfile[:len(logfile)-4]
	}
	cat := "+"
	if offset < 0 {
		cat = "-"
	}

	logfile += fSf("@%s%s%4.1f%s", zone, cat, float32(offset/3600.0), ".log")
	if abspath, err := filepath.Abs(logfile); err == nil {
		io.MustAppendFile(abspath, nil, false)
		if f, err := os.OpenFile(abspath, os.O_RDWR|os.O_CREATE|os.O_APPEND, FilePerm); err == nil {
			mPathFile[abspath] = f
			log.SetFlags(log.LstdFlags) // log.SetFlags(log.LstdFlags | log.LUTC)
			log.SetOutput(f)
			log2F = true
		}
	}
}

// resetLog : call once at exit
func resetLog() error {
	for logPath, f := range mPathFile {
		// delete empty error log
		fi, err := f.Stat()
		if err != nil {
			return err
		}
		if fi.Size() == 0 {
			if err := os.Remove(logPath); err != nil {
				return err
			}
		}
		// close
		f.Close()
	}
	mPathFile = make(map[string]*os.File)
	log.SetOutput(os.Stdout)
	log2F = false
	return nil
}

package logkit

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	. "github.com/digisan/go-generics"
	fd "github.com/digisan/gotk/file-dir"
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
func Log2F(enable, new bool, logFile string) (string, error) {
	if enable {
		if new {
			if _, err := remove(logFile); err != nil {
				return "", err
			}
		}
		return setLog(logFile), nil
	}
	return "", resetLog()
}

func remove(logFile string) (string, error) {
	dir := filepath.Dir(logFile)
	logFile = strings.TrimSuffix(filepath.Base(logFile), ".log")
	des, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}
	for _, de := range des {
		if strings.Contains(de.Name(), logFile+"@") {
			fPath := filepath.Join(dir, de.Name())
			if err := os.RemoveAll(fPath); err != nil {
				return "", err
			}
			return fPath, nil
		}
	}
	return "", nil
}

// FilePerm :
const FilePerm = 0666

// setLog :
func setLog(logFile string) string {
	zone, offset := time.Now().Zone()
	if sHasSuffix(logFile, ".log") {
		logFile = logFile[:len(logFile)-4]
	}

	logFile += fSf("@%s%s%4.1f.log", zone, IF(offset < 0, "-", "+"), float32(offset/3600.0))
	if absPath, err := filepath.Abs(logFile); err == nil {
		fd.MustAppendFile(absPath, nil, false)
		if f, err := os.OpenFile(absPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, FilePerm); err == nil {
			mPathFile[absPath] = f
			log.SetFlags(log.LstdFlags) // log.SetFlags(log.LstdFlags | log.LUTC)
			log.SetOutput(f)
			log2F = true
		}
	}
	return logFile
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

func Fac4GrpIdxLogF(group string, start int, ll logLevel, warnOnErr bool) func(v ...any) {
	index := int64(start - 1)
	return func(v ...any) {
		prefix := group
		if start >= 0 {
			prefix = fmt.Sprintf("%s %06d", group, atomic.AddInt64(&index, 1))
		}
		fn := Log
		switch ll {
		case DEBUG:
			fn = DebugP1
		case WARN:
			fn = WarnP1 // must warn even if no error
			if warnOnErr {
				fn = WarnP1OnErr
			}
		case FAIL:
			fn = FailP1OnErr
		}
		fn("%s - "+v[0].(string), append([]any{prefix}, v[1:]...)...)
	}
}

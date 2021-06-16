package logkit

// Warn :
func Warn(format string, v ...interface{}) {
	logger(2, WARN, format+"%v", append(v, fEf(""))...)
}

// WarnOnErr : write error into Console OR File
func WarnOnErr(format string, v ...interface{}) {
	logger(2, WARN, format, v...)
}

// WarnOnErrWhen : write error into Console OR File
func WarnOnErrWhen(condition bool, format string, v ...interface{}) {
	if condition {
		logger(2, WARN, format, v...)
	}
}

// WarnP1 :
func WarnP1(format string, v ...interface{}) {
	logger(3, WARN, format+"%v", append(v, fEf(""))...)
}

// WarnP1OnErr : write error into Console OR File
func WarnP1OnErr(format string, v ...interface{}) {
	logger(3, WARN, format, v...)
}

// WarnP1OnErrWhen : write error into Console OR File
func WarnP1OnErrWhen(condition bool, format string, v ...interface{}) {
	if condition {
		logger(3, WARN, format, v...)
	}
}

package logkit

// FailOnErr : error holder use "%v"
func FailOnErr(format string, v ...interface{}) {
	logger(2, FAIL, format, v...)
}

// FailOnErrWhen :
func FailOnErrWhen(condition bool, format string, v ...interface{}) {
	if condition {
		logger(2, FAIL, format, v...)
	}
}

// FailP1OnErr : error holder use "%v"
func FailP1OnErr(format string, v ...interface{}) {
	logger(3, FAIL, format, v...)
}

// FailP1OnErrWhen :
func FailP1OnErrWhen(condition bool, format string, v ...interface{}) {
	if condition {
		logger(3, FAIL, format, v...)
	}
}

// FailPnOnErr : error holder use "%v"
func FailPnOnErr(n int, format string, v ...interface{}) {
	logger(2+n, FAIL, format, v...)
}

// FailPnOnErrWhen :
func FailPnOnErrWhen(condition bool, n int, format string, v ...interface{}) {
	if condition {
		logger(2+n, FAIL, format, v...)
	}
}

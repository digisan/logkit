package logkit

// Debug : write info into Console OR File
func Debug(format string, v ...any) {
	logger(2, DEBUG, format, v...)
}

// DebugWhen : write info into Console OR File
func DebugWhen(condition bool, format string, v ...any) {
	if condition {
		logger(2, DEBUG, format, v...)
	}
}

// DebugP1 :
func DebugP1(format string, v ...any) {
	logger(3, DEBUG, format, v...)
}

// DebugP1When : write info into Console OR File
func DebugP1When(condition bool, format string, v ...any) {
	if condition {
		logger(3, DEBUG, format, v...)
	}
}

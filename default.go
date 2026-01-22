package log

func GetDefaultLogger() *Logger {
	if defaultLogger == nil {
		defaultLogger = NewLogger("main", INFO)
	}
	return defaultLogger
}

func SetDefaultLogger(l *Logger) {
	defaultLogger = l
}

func Debug() *MsgChain {
	return GetDefaultLogger().Debug()
}

func Info() *MsgChain {
	return GetDefaultLogger().Info()
}

func Warning() *MsgChain {
	return GetDefaultLogger().Warning()
}

func Error() *MsgChain {
	return GetDefaultLogger().Error()
}

func Fatal() *MsgChain {
	return GetDefaultLogger().Fatal()
}

func Panic() *MsgChain {
	return GetDefaultLogger().Panic()
}

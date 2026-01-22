package log

const (
	DEBUG           = 4
	INFO            = 3
	WARNING         = 2
	ERROR           = 1
	FATAL           = 0
	MAXVERBOSELEVEL = DEBUG
)

func logLevelToString(level int) string {
	switch i := level; i {
	case 0:
		return "fatal"
	case 1:
		return "error"
	case 2:
		return "warning"
	case 3:
		return "info"
	case 4:
		return "debug"
	default:
		return "unknown"
	}
}

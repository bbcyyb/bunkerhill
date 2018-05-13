package logs

import (
	"time"
)

// RFC5424 log message levels.
const (
	LevelCritical = iota
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
)

const levelLoggerImpl = -1

const (
	AdapterConsole = "console"
)

type Logger interface {
	Init(config string) error
	Write(when time.Time, msg string, level int) error
	Destroy()
	Flush()
}

type newLoggerFunc func() Logger

var adapters = make(map[string]newLoggerFunc)
var levelPrefix = [LevelDebug + 1]string{"[C] ", "[E] ", "[W] ", "[I] ", "[D] "}

func Register(name string, log newLoggerFunc) {
	if log == nil {
		panic("logs: Register provide is nil")
	}

	if _, ok := adapters[name]; ok {
		panic("logs: Register called twice for provider " + name)
	}

	adapters[name] = log
}

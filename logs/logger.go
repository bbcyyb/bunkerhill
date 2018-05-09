package main

import (
	"encoding/json"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

// RFC5424 log message levels.
const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

const (
	AdapterConsole = "console"
)

//=================== Logger Begin  =========================

type Logger interface {
	Init(config string) error
	WriteMsg(when time.Time, msg string, level int) error
	Destroy()
	Flush()
}

type newLoggerFunc func() Logger

var adapters = make(map[string]newLoggerFunc)
var levelPrefix = [LevelDebug + 1]string{"[M] ", "[A] ", "[C] ", "[E] ", "[W] ", "[N] ", "[I] ", "[D] "}

func Register(name string, log newLoggerFunc) {
	if log == nil {
		panic("logs: Register provide is nil")
	}

	if _, ok := adapters[name]; ok {
		panic("logs: Register called twice for provider ", name)
	}

	adapters[name] = log
}

//=================== Logger End  =========================

type brush func(string) string

func newBrush(color string) brush {
	pre := "\033["
	reset := "\033[0m"
	return func(text string) string {
		return pre + color + "m" + text + reset
	}
}

var colors = []brush{
	newBrush("1;37"), // Emergency          white
	newBrush("1;36"), // Alert              cyan
	newBrush("1;35"), // Critical           magenta
	newBrush("1;31"), // Error              red
	newBrush("1;33"), // Warning            yellow
	newBrush("1;32"), // Notice             green
	newBrush("1;34"), // Informational      blue
	newBrush("1;44"), // Debug              Background blue
}

//=================== consoleWriter Begin  =========================

type consoleWriter struct {
	lg       *logWriter
	Level    int  `json:"level"`
	Colorful bool `json:"color"`
}

func NewConsole() Logger {
	cw := &consoleWriter{
		lg:       newLogWriter(os.Stdout),
		Level:    LevelDebug,
		Colorful: runtime.GOOS != "windows",
	}

	return cw
}

func (c *consoleWriter) Init(jsonConfig string) error {
	if len(jsonConfig) == 0 {
		return nil
	}

	err := json.Unmarshal([]byte(jsonConfig), c)
	if runtime.GOOS == "windows" {
		c.Colorful = false
	}

	return err
}

func (c *consoleWriter) WriteMsg(when time.time, msg string, level int) error {
	if level > c.Level {
		return nil
	}

	if c.Colorful {
		msg = colors[level](msg)
	}

	c.lg.println(when, msg)
	return nil
}

func (c *consoleWriter) Destroy() {

}

func (c *consoleWriter) Flush() {

}

//=================== consoleWriter End  =========================

//=================== loggerWriter Begin  =========================

type logWriter struct {
	sync.Mutex
	writer io.Writer
}

func newLogWriter(wr io.Writer) *logWriter {
	return &logWriter{writer: wr}
}

func (lg *logWriter) println(when time.Time, msg string) {
	lg.Lock()
	h, _ := formatTimeHeader(when)
	lg.writer.Write(append(append(h, msg...), '\n'))
	lg.Unlock()
}

//=================== loggerWriter End  =========================

const (
	y1  = `0123456789`
	y2  = `0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789`
	y3  = `0000000000111111111122222222223333333333444444444455555555556666666666777777777788888888889999999999`
	y4  = `0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789`
	mo1 = `000000000111`
	mo2 = `123456789012`
	d1  = `0000000001111111111222222222233`
	d2  = `1234567890123456789012345678901`
	h1  = `000000000011111111112222`
	h2  = `012345678901234567890123`
	mi1 = `000000000011111111112222222222333333333344444444445555555555`
	mi2 = `012345678901234567890123456789012345678901234567890123456789`
	s1  = `000000000011111111112222222222333333333344444444445555555555`
	s2  = `012345678901234567890123456789012345678901234567890123456789`
	ns1 = `0123456789`
)

func formatTimeHeader(when time.Time) ([]byte, int) {
	y, mo, d := when.Date()
	h, mi, s := when.Clock()
	ns := when.Nanosecond() / 1000000
	//len("2006/01/02 15:04:05.123 ")==24
	var buf [24]byte

	buf[0] = y1[y/1000%10]
	buf[1] = y2[y/100]
	buf[2] = y3[y-y/100*100]
	buf[3] = y4[y-y/100*100]
	buf[4] = '/'
	buf[5] = mo1[mo-1]
	buf[6] = mo2[mo-1]
	buf[7] = '/'
	buf[8] = d1[d-1]
	buf[9] = d2[d-1]
	buf[10] = ' '
	buf[11] = h1[h]
	buf[12] = h2[h]
	buf[13] = ':'
	buf[14] = mi1[mi]
	buf[15] = mi2[mi]
	buf[16] = ':'
	buf[17] = s1[s]
	buf[18] = s2[s]
	buf[19] = '.'
	buf[20] = ns1[ns/100]
	buf[21] = ns1[ns%100/10]
	buf[22] = ns1[ns%10]

	buf[23] = ' '

	return buf[0:], d
}

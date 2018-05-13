package logs

import (
	"encoding/json"
	"os"
	"runtime"
	"time"
)

type consoleLogger struct {
	writer   *logWriter
	Level    int  `json:"level"`
	Colorful bool `json:"color"`
}

func NewConsole() Logger {
	cl := &consoleLogger{
		writer:   newLogWriter(os.Stdout),
		Level:    LevelDebug,
		Colorful: runtime.GOOS != "windows",
	}

	return cl
}

func (c *consoleLogger) Init(jsonConfig string) error {
	if len(jsonConfig) == 0 {
		return nil
	}

	err := json.Unmarshal([]byte(jsonConfig), c)
	if runtime.GOOS == "windows" {
		c.Colorful = false
	}

	return err
}

func (c *consoleLogger) Write(when time.Time, msg string, level int) error {
	if level > c.Level {
		return nil
	}

	t := formatTimeShort(when)

	header := "[" + string(t) + "]"

	if c.Colorful {
		header = colors[level](header)
	}

	msg = header + msg

	c.writer.write(msg)
	return nil
}

func (c *consoleLogger) Destroy() {

}

func (c *consoleLogger) Flush() {

}

func init() {
	Register(AdapterConsole, NewConsole)
}

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

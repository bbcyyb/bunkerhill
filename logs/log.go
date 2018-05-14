package logs

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type LoggerWrapper struct {
	lock                sync.Mutex
	level               int
	init                bool
	enableFuncCallDepth bool
	loggerFuncCallDepth int
	asynchronous        bool
	msgChanLen          int64
	msgChan             chan *logMsg
	signalChan          chan string
	wg                  sync.WaitGroup
	outputs             []*nameLogger
}

const defaultAsyncMsgLen = 1e3

type nameLogger struct {
	Logger
	name string
}

type logMsg struct {
	level int
	msg   string
	when  time.Time
}

var logMsgPool *sync.Pool

func NewLogger(channelLens ...int64) *LoggerWrapper {
	lw := new(LoggerWrapper)
	lw.level = LevelDebug
	lw.loggerFuncCallDepth = 2
	lw.msgChanLen = append(channelLens, 0)[0]
	if lw.msgChanLen <= 0 {
		lw.msgChanLen = defaultAsyncMsgLen
	}
	lw.signalChan = make(chan string, 1)
	// Console Logger is set as default logger handler
	lw.setLogger(AdapterConsole)
	return lw
}

func (lw *LoggerWrapper) Async(msgLen ...int64) *LoggerWrapper {
	lw.lock.Lock()
	defer lw.lock.Unlock()
	if lw.asynchronous {
		return lw
	}

	lw.asynchronous = true
	if len(msgLen) > 0 && msgLen[0] > 0 {
		lw.msgChanLen = msgLen[0]
	}

	lw.msgChan = make(chan *logMsg, lw.msgChanLen)
	logMsgPool = &sync.Pool{
		New: func() interface{} {
			return &logMsg{}
		},
	}

	lw.wg.Add(1)
	go lw.startLogger()
	return lw
}

func (lw *LoggerWrapper) setLogger(adapterName string, configs ...string) error {
	config := append(configs, "{}")[0]
	for _, l := range lw.outputs {
		if l.name == adapterName {
			return fmt.Errorf("logs: duplicate adaptername %q (you have set this logger before)", adapterName)
		}
	}

	log, ok := adapters[adapterName]
	if !ok {
		return fmt.Errorf("logs: unknown adaptername %q (forgotten Register?)", adapterName)
	}

	lg := log()
	err := lg.Init(config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "logs.LoggerWrapper.SetLogger: "+err.Error())
		return err
	}

	lw.outputs = append(lw.outputs, &nameLogger{name: adapterName, Logger: lg})
	return nil
}

func (lw *LoggerWrapper) SetLogger(adapterName string, configs ...string) error {
	lw.lock.Lock()
	defer lw.lock.Unlock()
	if !lw.init {
		lw.outputs = []*nameLogger{}
		lw.init = true
	}

	return lw.setLogger(adapterName, configs...)
}

func (lw *LoggerWrapper) DelLogger(adapterName string) error {
	lw.lock.Lock()
	defer lw.lock.Unlock()
	outputs := []*nameLogger{}
	for _, lg := range lw.outputs {
		if lg.name == adapterName {
			lg.Destroy()
		} else {
			outputs = append(outputs, lg)
		}
	}

	if len(outputs) == len(lw.outputs) {
		return fmt.Errorf("logs: unknown adaptername %q (forgotten Register?)", adapterName)
	}

	lw.outputs = outputs
	return nil
}

func (lw *LoggerWrapper) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	if p[len(p)-1] == '\n' {
		p = p[0 : len(p)-1]
	}

	err = lw.writeMsg(levelLoggerImpl, string(p))
	if err == nil {
		return len(p), err
	}

	return 0, err
}

func (lw *LoggerWrapper) writeMsg(logLevel int, msg string, v ...interface{}) error {
	if !lw.init {
		lw.lock.Lock()
		lw.setLogger(AdapterConsole)
		lw.lock.Unlock()
	}

	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}

	when := time.Now()
	if lw.enableFuncCallDepth {
		_, file, line, ok := runtime.Caller(lw.loggerFuncCallDepth)
		if !ok {
			file = "???"
			line = 0
		}

		_, filename := path.Split(file)
		msg = "[" + filename + ":" + strconv.Itoa(line) + "]" + msg
	}

	if logLevel == levelLoggerImpl {
		logLevel = LevelCritical
	} else {
		msg = levelPrefix[logLevel] + msg
	}

	if lw.asynchronous {
		lm := logMsgPool.Get().(*logMsg)
		lm.level = logLevel
		lm.msg = msg
		lm.when = when
		lw.msgChan <- lm
	} else {
		lw.writeToLoggers(when, msg, logLevel)
	}

	return nil
}

func (lw *LoggerWrapper) writeToLoggers(when time.Time, msg string, level int) {
	for _, l := range lw.outputs {
		err := l.Write(when, msg, level)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to Write to adapter:%v,error:%v\n", l.name, err)
		}
	}
}

func (lw *LoggerWrapper) SetLevel(l int) *LoggerWrapper {
	lw.level = l
	return lw
}

func (lw *LoggerWrapper) SetLogFuncCallDepth(d int) *LoggerWrapper {
	lw.loggerFuncCallDepth = d
	return lw
}

func (lw *LoggerWrapper) GetLogFuncCallDepth() int {
	return lw.loggerFuncCallDepth
}

func (lw *LoggerWrapper) EnableFuncCallDepth(b bool) *LoggerWrapper {
	lw.enableFuncCallDepth = b
	return lw
}

func (lw *LoggerWrapper) startLogger() {
	gameOver := false
	for {
		select {
		case lm := <-lw.msgChan:
			lw.writeToLoggers(lm.when, lm.msg, lm.level)
			logMsgPool.Put(lm)
		case sg := <-lw.signalChan:
			lw.flush()
			if sg == "close" {
				for _, l := range lw.outputs {
					l.Destroy()
				}

				lw.outputs = nil
				gameOver = true
			}

			lw.wg.Done()
		}

		if gameOver {
			break
		}
	}
}

func FormatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}

func (lw *LoggerWrapper) Critical(format string, v ...interface{}) {
	if LevelCritical > lw.level {
		return
	}

	lw.writeMsg(LevelCritical, format, v...)
}

func (lw *LoggerWrapper) Error(format string, v ...interface{}) {
	if LevelError > lw.level {
		return
	}

	lw.writeMsg(LevelError, format, v...)
}

func (lw *LoggerWrapper) Wran(format string, v ...interface{}) {
	if LevelWarning > lw.level {
		return
	}

	lw.writeMsg(LevelWarning, format, v...)
}

func (lw *LoggerWrapper) Info(format string, v ...interface{}) {
	if LevelInfo > lw.level {
		return
	}

	lw.writeMsg(LevelInfo, format, v...)
}

func (lw *LoggerWrapper) Debug(format string, v ...interface{}) {
	if LevelDebug > lw.level {
		return
	}

	lw.writeMsg(LevelDebug, format, v...)
}

func (lw *LoggerWrapper) Flush() {
	if lw.asynchronous {
		lw.signalChan <- "flush"
		lw.wg.Wait()
		lw.wg.Add(1)
		return
	}

	lw.flush()
}

func (lw *LoggerWrapper) flush() {
	if lw.asynchronous {
		for {
			if len(lw.msgChan) > 0 {
				lm := <-lw.msgChan
				lw.writeToLoggers(lm.when, lm.msg, lm.level)
				logMsgPool.Put(lm)
				continue
			}
			break
		}
	}

	for _, l := range lw.outputs {
		l.Flush()
	}
}

func (lw *LoggerWrapper) Close() {
	if lw.asynchronous {
		lw.signalChan <- "close"
		lw.wg.Wait()
		close(lw.msgChan)
	} else {
		lw.flush()
	}
}

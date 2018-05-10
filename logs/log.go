package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
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

	lw.msgChan = make(chan *logMsg, bl.msgChanLen)
	logMsgPool = &sync.Pool{
		New: func() interface{} {
			return &logMsg{}
		},
	}

	lw.wg.Add(1)
	go lw.startLogger()
	return lw
}

func (lw *LoggerWrapper) setLoggerCore(adapterName string, configs ...string) error {
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
	def lw.lock.Unlock()
	if !lw.init {
		lw.outputs = []*nameLogger{}
		lw.init = true
	}

	return lw.setLoggerCore(adapterName, configs...)
}

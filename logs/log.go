package main

import (
	"encoding/json"
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
}

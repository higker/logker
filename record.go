// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/16 - 10:59 下午

package logker

import (
	"fmt"
	"github.com/fatih/color"
)

const (
	//[INFO] 2006-01-02 13:05.0006 MP - Position: test.go|main.test:21 - Message: news
	format     = "[%s] - Date: %s  %s - Message: %s"
	fileFormat = format + "\n"
)

// Logging record
type logRecord interface {
	OutPutMessage(v string)
}

func (c *console) OutPutMessage(model level, v string) {
	switch model.toStr() {
	case DEBUG.toStr():
		// blue color of log message.
		// format log message output console.
		color.Blue(format, DEBUG.toStr(), c.tz.NowTimeStr(), buildCallerStr(SKIP), v)
	case INFO.toStr():
		color.Green(format, INFO.toStr(), c.tz.NowTimeStr(), buildCallerStr(SKIP), v)
	case WARNING.toStr():
		color.Yellow(format, WARNING.toStr(), c.tz.NowTimeStr(), buildCallerStr(SKIP), v)
	case ERROR.toStr():
		color.Red(format, ERROR.toStr(), c.tz.NowTimeStr(), buildCallerStr(SKIP), v)
	default:
		// Log Level Type Error
		// Program automatically set to debug
		color.Red("-----------------------------------------------------------------")
		color.Red("！！！Log Level Type Error,Program automatically set to debug！！！")
		color.Red("-----------------------------------------------------------------")
		c.logLevel = DEBUG
		// recursion
		c.OutPutMessage(model, v)
	}
}

func (f *fileLog) OutPutMessage(model level, v string) {
	switch model.toStr() {
	case DEBUG.toStr():
		f.OutPut(DEBUG, v)
	case INFO.toStr():
		f.OutPut(INFO, v)
	case WARNING.toStr():
		f.OutPut(WARNING, v)
	case ERROR.toStr():
		f.OutPut(ERROR, v)
	default:
		// Log Level Type Error
		// Program automatically set to debug
		f.logLevel = DEBUG
		// recursion
		f.OutPutMessage(model, v)
	}
}

func (f *fileLog) OutPut(lev level, v string) {
	_, err := f.file.WriteString(fmt.Sprintf(fileFormat, lev.toStr(), f.tz.NowTimeStr(), buildCallerStr(SKIP), v))
	_ = f.file.Sync()
	if err != nil {
		_ = f.file.Close()
		panic("output message to log file fail. filePath:" + f.directory + "/" + f.fileName + ".log")
	}
	//fmt.Println(n)

}

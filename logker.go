// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/1 - 1:55 上午

package logker

import (
	"os"
)

// This is package version
const Version = "1.2.2"

/*
 ____ ____ ____ ____ ____ ____
||L |||o |||g |||K |||e |||r ||
||__|||__|||__|||__|||__|||__||
|/__\|/__\|/__\|/__\|/__\|/__\|
zh_CN:
LogKer是Golang语言的日志操作库.
1.控制台输出 (已经实现)
2.文件输出 (已经实现)
3.WebSocket输出 (未来将会支持)
4.网络kafka输出  (未来将会支持)
*/

// Build console Logger
/*
 lev : Logging level
 zone : Logging time zone
 formatting : Logging format template string
 at : AsyncTask Pointer
*/
func NewClog(lev level, zone logTimeZone, formatting string, at *AsyncTask) (Logger, error) {
	consoleLog := &console{
		// Level: logger Level
		logLevel: lev,
		// Zone : logger Time Zone
		timeZone:  zone,
		tz:        nil,
		asyncTask: at,
	}
	if err := verify(formatting); err != nil {
		return nil, err
	}
	consoleLog.formatting = formatting
	consoleLog.initTime()
	consoleLog.begin()
	return consoleLog, nil
}

// Build File logger
/*
 lev : Logging level
 zone : Logging time zone
 formatting : Logging format template string
 at : AsyncTask Pointer
 wheErr : Whether To Open A Separate ErrorFile
 dir : Logging file output directory
 fileName : Logging Name
 size : File size
 power : File system power
*/
func NewFlog(lev level, wheErr bool, zone logTimeZone, dir string, fileName string, size int64, power os.FileMode, formatting string, at *AsyncTask) (Logger, error) {
	fg := &fileLog{
		// logLevel:    lev,        logging level
		logLevel: lev,
		// wheError:    wheErr,     whether enable  error alone file
		wheError: wheErr,
		// directory:   dir,	    logging file save directory
		directory: dir,
		// fileName:    fileName,   logging save file name
		fileName: fileName,
		file:     nil,
		errFile:  nil,
		tz:       nil,
		// timeZone:    zone,	    load time zone format
		timeZone: zone,
		// power:       power,      file system power
		power: power,
		// fileMaxSize: size,       logging alone file max size
		fileMaxSize: size,
	}
	if err := verify(formatting); err != nil {
		return nil, err
	}
	fg.formatting = formatting
	fg.asyncTask = at
	fg.begin()
	fg.tz = &timeZone{TimeZoneStr: fg.timeZone}
	fg.file = fg.initFilePtr()
	if fg.isEnableErr() {
		fg.errFile = fg.initErrPtr()
	}
	return fg, nil
}

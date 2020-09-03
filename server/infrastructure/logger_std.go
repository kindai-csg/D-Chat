package infrastructure

import "log"

type LoggerStd struct{}

func NewLoggerStd() *LoggerStd {
	return &LoggerStd{}
}

func (handler *LoggerStd) Write(out interface{}) {
	log.Println(out)
}

func (handler *LoggerStd) Warnning(out interface{}) {
	log.Println(out)
}

func (handler *LoggerStd) Error(out interface{}) {
	log.Println(out)
}

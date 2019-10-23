// Package klog
package klog

const (
	FileLogger = 1
)

type Config struct {

}



type Logger interface {
	Config(conf Config) Logger
}

type fileLogger struct {

}

func (l *fileLogger) Config(conf Config) Logger {
	return l
}

type defaultLogger struct {

}

func NewFileLogger() Logger {
	return &fileLogger{}
}
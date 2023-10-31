package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type LogWriter struct {
	*log.Logger
	scope   string
	enabled bool
}

func (l *LogWriter) Enable() {
	l.SetOutput(os.Stdout)
	l.enabled = true
}

func (l *LogWriter) Disable() {
	l.SetOutput(io.Discard)
	l.enabled = false
}

func newLogger(scope string) *LogWriter {
	prefix := fmt.Sprintf("%s:", scope)
	logger := log.New(log.Writer(), prefix, log.Lshortfile)
	return &LogWriter{
		Logger:  logger,
		scope:   scope,
		enabled: true,
	}
}

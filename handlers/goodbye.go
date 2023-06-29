package handlers

import "log"

type Goodbye struct {
	l *log.Logger
}

func newGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

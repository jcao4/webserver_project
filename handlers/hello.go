package handlers

import "log"

type Hello struct {
	l *log.Logger
}

func newHello(l *log.Logger) *Hello {
	return &Hello{l}
}

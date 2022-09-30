package handlers

import (
	"container/list"
	"github.com/bindasov/spaceBattle/commands/exceptionstask/repeat"
)

func NewLogHandler(logCommand *repeat.LogCommand, queue *list.List) *logHandler {
	logHandler := &logHandler{
		logCommand: logCommand,
		queue:      queue,
	}
	return logHandler
}

type logHandler struct {
	logCommand *repeat.LogCommand
	queue      *list.List
}

func (eh *logHandler) Handle() {
	eh.queue.PushBack(eh.logCommand)
}

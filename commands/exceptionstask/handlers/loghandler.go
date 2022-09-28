package handlers

import (
	"container/list"
	"github.com/bindasov/spaceBattle/commands/exceptionstask/commands"
)

func NewLogHandler(logCommand *commands.LogCommand, queue *list.List) *logHandler {
	logHandler := &logHandler{
		logCommand: logCommand,
		queue:      queue,
	}
	return logHandler
}

type logHandler struct {
	logCommand *commands.LogCommand
	queue      *list.List
}

func (eh *logHandler) Handle() {
	eh.queue.PushBack(eh.logCommand)
}

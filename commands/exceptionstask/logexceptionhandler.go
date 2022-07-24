package exceptionstask

import (
	"container/list"
)

type LogExceptionHandler interface {
	Handle()
	CacheCommand(command logCommand)
	CacheQueue(queue *list.List)
}

func NewLogExceptionHandler() LogExceptionHandler {
	logExceptionHandler := &logExceptionHandler{}
	return logExceptionHandler
}

type logExceptionHandler struct {
	logCommand logCommand
	queue      *list.List
}

func (eh *logExceptionHandler) Handle() {
	eh.queue.PushBack(eh.logCommand)
}

func (eh *logExceptionHandler) CacheCommand(command logCommand) {
	eh.logCommand = command
}

func (eh *logExceptionHandler) CacheQueue(queue *list.List) {
	eh.queue = queue
}

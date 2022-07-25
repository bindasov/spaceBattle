package exceptionstask

import (
	"container/list"
)

type LogExceptionHandler interface {
	Handle()
	CacheCommand(command LogCommand)
	CacheQueue(queue *list.List)
}

func NewLogExceptionHandler() LogExceptionHandler {
	logExceptionHandler := &logExceptionHandler{}
	return logExceptionHandler
}

type logExceptionHandler struct {
	logCommand LogCommand
	queue      *list.List
}

func (eh *logExceptionHandler) Handle() {
	eh.queue.PushBack(eh.logCommand)
}

func (eh *logExceptionHandler) CacheCommand(command LogCommand) {
	eh.logCommand = command
}

func (eh *logExceptionHandler) CacheQueue(queue *list.List) {
	eh.queue = queue
}

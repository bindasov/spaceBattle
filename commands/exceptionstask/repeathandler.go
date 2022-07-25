package exceptionstask

import (
	"container/list"
)

type RepeatHandler interface {
	Handle()
	CacheCommand(command RepeatCommand)
	CacheQueue(queue *list.List)
}

func NewRepeatHandler() RepeatHandler {
	repeatHandler := &repeatHandler{}
	return repeatHandler
}

type repeatHandler struct {
	command RepeatCommand
	queue   *list.List
}

func (rh *repeatHandler) Handle() {
	rh.queue.PushBack(rh.command)
}

func (rh *repeatHandler) CacheQueue(queue *list.List) {
	rh.queue = queue
}

func (rh *repeatHandler) CacheCommand(command RepeatCommand) {
	rh.command = command
}

package exceptionstask

import "container/list"

type RepeatHandler interface {
	Handle()
	CacheQueue(queue *list.List)
}

func NewRepeatExceptionHandler(command RepeatCommand) RepeatHandler {
	repeatHandler := &repeatHandler{
		command: command,
	}
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

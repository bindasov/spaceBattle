package exceptionstask

import (
	"container/list"
	"github.com/bindasov/spaceBattle/commands"

	"github.com/bindasov/spaceBattle/commands/exceptionstask/base"
	"github.com/bindasov/spaceBattle/commands/exceptionstask/handlers"
	"github.com/bindasov/spaceBattle/commands/exceptionstask/repeat"
	"github.com/bindasov/spaceBattle/logger"
)

func RunRepeat() {
	// preparing

	config := &logger.Config{Level: "info"}
	log, _ := logger.InitLogger(config)

	queue := list.New()
	cmdErrHandler := make(map[error]base.Handler)

	// creating base command

	command := base.NewCommand()

	// pushing command to queue

	queue.PushBack(command)

	// creating commands

	logCommand := repeat.NewLogCommand(log)
	repeatCommand := repeat.NewRepeatCommand(command)

	// creating error handlers

	logExceptionHandler := handlers.NewLogHandler(logCommand, queue)
	repeatHandler := handlers.NewRepeatHandler(repeatCommand, queue)

	// mapping errors to handlers

	cmdErrHandler[base.CommandError] = repeatHandler
	cmdErrHandler[base.RepeatCommandError] = logExceptionHandler

	// executing commands from queue

	for e := queue.Front(); e != nil; e = e.Next() {
		v := (e.Value).(commands.Command)
		err := v.Execute()
		if _, ok := cmdErrHandler[err]; ok {
			logCommand.SetError(err)
			cmdErrHandler[err].Handle()
		}
	}
}

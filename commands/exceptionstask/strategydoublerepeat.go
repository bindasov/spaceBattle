package exceptionstask

import (
	"container/list"
	"github.com/bindasov/spaceBattle/commands"

	"github.com/bindasov/spaceBattle/commands/exceptionstask/base"
	"github.com/bindasov/spaceBattle/commands/exceptionstask/repeat"
	"github.com/bindasov/spaceBattle/logger"
)

func RunDoubleRepeat() {
	// preparing

	config := &logger.Config{Level: "info"}
	log, _ := logger.InitLogger(config)

	queue := list.New()

	// creating base command

	command := base.NewCommand()

	// creating commands

	logCommand := repeat.NewLogCommand(log)
	doubleRepeatCommand := repeat.NewDoubleRepeatCommand(command)

	// pushing double repeat command to queue

	queue.PushBack(doubleRepeatCommand)

	// executing commands from queue

	for e := queue.Front(); e != nil; e = e.Next() {
		v := (e.Value).(commands.Command)
		err := v.Execute()

		if _, ok := (e.Value).(*repeat.DoubleRepeatCommand); err != nil && ok {
			logCommand.SetError(base.DoubleRepeatCommandError)
			_ = logCommand.Execute()
		}
	}
}

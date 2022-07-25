package exceptionstask

import (
	"github.com/bindasov/spaceBattle/logger"
)

type LogCommand interface {
	Execute()
	SetError(err error)
}

func NewLogCommand(log logger.Logger) LogCommand {
	logCommand := &logCommand{
		log: log,
	}
	return logCommand
}

type logCommand struct {
	log logger.Logger
	err error
}

func (lc *logCommand) Execute() {
	lc.log.Errorf(lc.err.Error())
}

func (lc *logCommand) SetError(err error) {
	lc.err = err
}

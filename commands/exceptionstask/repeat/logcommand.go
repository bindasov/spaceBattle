package repeat

import "github.com/bindasov/spaceBattle/logger"

func NewLogCommand(log logger.Logger) *LogCommand {
	logCommand := &LogCommand{
		log: log,
	}
	return logCommand
}

type LogCommand struct {
	log logger.Logger
	err error
}

func (lc *LogCommand) Execute() error {
	lc.log.Errorf(lc.err.Error())
	return nil
}

func (lc *LogCommand) SetError(err error) {
	lc.err = err
}

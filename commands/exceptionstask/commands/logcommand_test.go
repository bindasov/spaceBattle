package commands

import (
	"bytes"
	"errors"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/bindasov/spaceBattle/logger"
)

func TestLogCommand_Execute(t *testing.T) {
	config := &logger.Config{Level: "info"}
	commandLogger, _ := logger.InitLogger(config)
	logCommand := NewLogCommand(commandLogger)
	commandError := errors.New("error")
	logCommand.SetError(commandError)

	oldOut := log.StandardLogger().Out
	buf := bytes.Buffer{}
	log.SetOutput(&buf)
	_ = logCommand.Execute()
	log.SetOutput(oldOut)
}

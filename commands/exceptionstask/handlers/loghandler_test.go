package handlers

import (
	"container/list"
	"github.com/bindasov/spaceBattle/commands/exceptionstask/repeat"
	"github.com/bindasov/spaceBattle/logger"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogHandler_Handle(t *testing.T) {
	type deps struct {
		queue      *list.List
		logCommand *repeat.LogCommand
		handler    *logHandler
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "success",
			handler: func(t *testing.T, deps *deps) {
				deps.handler.Handle()

				require.Equal(t, deps.logCommand, deps.queue.Back().Value)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			config := &logger.Config{Level: "info"}
			commandLogger, _ := logger.InitLogger(config)
			logCommand := repeat.NewLogCommand(commandLogger)
			queue := &list.List{}
			handler := NewLogHandler(logCommand, queue)

			deps := &deps{
				queue:      queue,
				logCommand: logCommand,
				handler:    handler,
			}

			tc.handler(t, deps)
		})
	}
}

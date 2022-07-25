package exceptionstask

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bindasov/spaceBattle/commands/exceptionstask/mocks"
)

func TestLogCommand_Execute(t *testing.T) {
	type deps struct {
		handler LogExceptionHandler
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "success",
			handler: func(t *testing.T, deps *deps) {
				command := &mocks.LogCommand{}
				queue := &list.List{}
				deps.handler.CacheCommand(command)
				deps.handler.CacheQueue(queue)

				deps.handler.Handle()

				require.Equal(t, command, queue.Back().Value)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handler := NewLogExceptionHandler()

			deps := &deps{
				handler: handler,
			}

			tc.handler(t, deps)
		})
	}
}

package handlers

import (
	"container/list"
	"github.com/bindasov/spaceBattle/commands/exceptionstask/base"
	"github.com/bindasov/spaceBattle/commands/exceptionstask/repeat"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepeatHandler_Handle(t *testing.T) {
	type deps struct {
		queue         *list.List
		repeatCommand *repeat.RepeatCommand
		handler       *RepeatHandler
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "success",
			handler: func(t *testing.T, deps *deps) {
				deps.handler.Handle()

				require.Equal(t, deps.repeatCommand, deps.queue.Back().Value)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			command := base.NewCommand()
			queue := &list.List{}
			repeatCommand := repeat.NewRepeatCommand(command)
			handler := NewRepeatHandler(repeatCommand, queue)

			deps := &deps{
				queue:         queue,
				repeatCommand: repeatCommand,
				handler:       handler,
			}

			tc.handler(t, deps)
		})
	}
}

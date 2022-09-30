package multithreading

import (
	"github.com/bindasov/spaceBattle/commands"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStartCommand_Execute(t *testing.T) {
	tests := []struct {
		name    string
		handler func(*testing.T)
	}{
		{
			name: "coroutine is running",
			handler: func(t *testing.T) {
				queue := make(chan commands.Command)

				startActor := NewStart(queue)
				go func() {
					_ = startActor.Execute()
				}()

				sumCommand := NewSum(2, 3)
				queue <- sumCommand
				require.Equal(t, 5, sumCommand.GetResult())
			},
		},
		{
			name: "soft stop",
			handler: func(t *testing.T) {
				queue := make(chan commands.Command)

				startActor := NewStart(queue)
				go func() {
					_ = startActor.Execute()
				}()

				sumCommand := NewSum(2, 3)
				softStop := NewSoftStop(queue)

				go func() {
					_ = startActor.Execute()
				}()
				queue <- sumCommand
				queue <- softStop
				_, ok := <-queue
				require.Equal(t, false, ok)
				require.Equal(t, 5, sumCommand.GetResult())
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.handler(t)
		})
	}
}

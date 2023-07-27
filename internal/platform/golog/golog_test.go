package golog

import (
	"context"
	"io"
	"os"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	// disable init once during tests
	initOnce = false

	type testInput struct {
		ctx     context.Context
		message string
		opts    []Options
	}

	type test struct {
		name   string
		setup  func(t *testing.T) (*os.File, *os.File)
		input  testInput
		assert func(t *testing.T, input testInput, read, write *os.File)
	}

	tests := []test{
		// nolint:dupl
		{
			name: "should log using local info logger",
			setup: func(t *testing.T) (read *os.File, write *os.File) {
				read, write, err := os.Pipe()
				if err != nil {
					t.Fatalf("failed to create pipe file. Cause: %v", err)
				}
				logWriter = write
				return read, write
			},
			input: testInput{
				ctx:     context.Background(),
				message: "LOCAL INFO LOGGER",
				opts:    []Options{},
			},
			assert: func(t *testing.T, input testInput, read, write *os.File) {
				Log().Info(input.ctx, input.message, input.opts...)
				write.Close()
				out, err := io.ReadAll(read)
				read.Close()

				switch {
				case err != nil:
					t.Fatalf("failed to read from stdout. Cause: %s", err)
				case !strings.Contains(string(out), "[  INFO  ]"):
					t.Fatalf("the expected output must have: '[  INFO  ]' but got: '%s'", string(out))
				}
			},
		},
		// nolint:dupl
		{
			name: "should log using local warn logger",
			setup: func(t *testing.T) (read *os.File, write *os.File) {
				read, write, err := os.Pipe()
				if err != nil {
					t.Fatalf("failed to create pipe file. Cause: %v", err)
				}
				logWriter = write
				return read, write
			},
			input: testInput{
				ctx:     context.Background(),
				message: "LOCAL WARN LOGGER",
				opts:    []Options{},
			},
			assert: func(t *testing.T, input testInput, read, write *os.File) {
				Log().Warn(input.ctx, input.message, input.opts...)
				write.Close()
				out, err := io.ReadAll(read)
				read.Close()

				switch {
				case err != nil:
					t.Fatalf("failed to read from stdout. Cause: %s", err)
				case !strings.Contains(string(out), "[  WARN  ]"):
					t.Fatalf("the expected output must have: '[  WARN  ]' but got: '%s'", string(out))
				}
			},
		},
		// nolint:dupl
		{
			name: "should log using local debug logger",
			setup: func(t *testing.T) (read *os.File, write *os.File) {
				read, write, err := os.Pipe()
				if err != nil {
					t.Fatalf("failed to create pipe file. Cause: %v", err)
				}
				logWriter = write
				return read, write
			},
			input: testInput{
				ctx:     context.Background(),
				message: "LOCAL DEBUG LOGGER",
				opts:    []Options{},
			},
			assert: func(t *testing.T, input testInput, read, write *os.File) {
				Log().Debug(input.ctx, input.message, input.opts...)
				write.Close()
				out, err := io.ReadAll(read)
				read.Close()

				switch {
				case err != nil:
					t.Fatalf("failed to read from stdout. Cause: %s", err)
				case !strings.Contains(string(out), "[  DEBUG ]"):
					t.Fatalf("the expected output must have: '[  DEBUG ]' but got: '%s'", string(out))
				}
			},
		},
		// nolint:dupl
		{
			name: "should log using local error logger",
			setup: func(t *testing.T) (read *os.File, write *os.File) {
				read, write, err := os.Pipe()
				if err != nil {
					t.Fatalf("failed to create pipe file. Cause: %v", err)
				}
				logWriter = write
				return read, write
			},
			input: testInput{
				ctx:     context.Background(),
				message: "LOCAL ERROR LOGGER",
				opts:    []Options{},
			},
			assert: func(t *testing.T, input testInput, read, write *os.File) {
				Log().Error(input.ctx, input.message, input.opts...)
				write.Close()
				out, err := io.ReadAll(read)
				read.Close()

				switch {
				case err != nil:
					t.Fatalf("failed to read from stdout. Cause: %s", err)
				case !strings.Contains(string(out), "[  ERROR ]"):
					t.Fatalf("the expected output must have: '[  ERROR ]' but got: '%s'", string(out))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			read, write := tt.setup(t)
			tt.assert(t, tt.input, read, write)
		})
	}
}

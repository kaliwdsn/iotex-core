// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

// A warrper for Zerolog (https://github.com/rs/zerolog)
//
// Package log provides a global logger for zerolog.
// derived from https://github.com/rs/zerolog/blob/master/log/log.go
// putting here to get a better integration

package logger

import (
	"context"
	"flag"
	"io"
	"os"

	"github.com/rs/zerolog"
)

var (
	logger   *zerolog.Logger
	levelStr string
	pathStr  string
)

func init() {
	flag.StringVar(&levelStr, "log-level", zerolog.InfoLevel.String(), "Log level")
	flag.StringVar(&pathStr, "log-path", "", "Log path")

}

// Logger gets the logger client instance
func Logger() *zerolog.Logger {
	if logger == nil {
		level, err := zerolog.ParseLevel(levelStr)
		if err != nil {
			panic(err)
		}
		var l zerolog.Logger
		if pathStr == "" {
			l = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).Level(level).With().Timestamp().Logger()
		} else {
			file, err := os.Create(pathStr)
			if err != nil {
				panic(err)
			}
			l = zerolog.New(file).Level(level).With().Timestamp().Logger()
		}
		logger = &l
	}
	return logger
}

// SetLogger sets the global logger client pointer to arbitrary logger client instance
func SetLogger(l *zerolog.Logger) {
	logger = l
}

// Output duplicates the global logger and sets w as its output.
func Output(w io.Writer) zerolog.Logger {
	return Logger().Output(w)
}

// With creates a child logger with the field added to its context.
func With() zerolog.Context {
	return Logger().With()
}

// Level creates a child logger with the minimum accepted level set to level.
func Level(level zerolog.Level) zerolog.Logger {
	return Logger().Level(level)
}

// Sample returns a logger with the s sampler.
func Sample(s zerolog.Sampler) zerolog.Logger {
	return Logger().Sample(s)
}

// Hook returns a logger with the h Hook.
func Hook(h zerolog.Hook) zerolog.Logger {
	return Logger().Hook(h)
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() *zerolog.Event {
	return Logger().Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() *zerolog.Event {
	return Logger().Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() *zerolog.Event {
	return Logger().Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error() *zerolog.Event {
	return Logger().Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() *zerolog.Event {
	return Logger().Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic() *zerolog.Event {
	return Logger().Panic()
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func WithLevel(level zerolog.Level) *zerolog.Event {
	return Logger().WithLevel(level)
}

// Log starts a new message with no level. Setting zerolog.GlobalLevel to
// zerolog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func Log() *zerolog.Event {
	return Logger().Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	Logger().Print(v...)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	Logger().Printf(format, v...)
}

// Ctx returns the logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}

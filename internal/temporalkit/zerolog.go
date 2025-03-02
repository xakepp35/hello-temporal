package temporalkit

import (
	"github.com/rs/zerolog"
)

type Zerolog struct {
	log *zerolog.Logger
}

func NewZerolog(log *zerolog.Logger) *Zerolog {
	return &Zerolog{
		log: log,
	}
}

func (s *Zerolog) Debug(msg string, keyvals ...any) {
	if s.log == nil {
		return
	}
	event := s.log.Debug()
	for i := 0; i < len(keyvals); i += 2 {
		event = event.Any(keyvals[i].(string), keyvals[i+1])
	}
	event.Msg(msg)
}

func (s *Zerolog) Info(msg string, keyvals ...any) {
	if s.log == nil {
		return
	}
	event := s.log.Info()
	for i := 0; i < len(keyvals); i += 2 {
		event = event.Any(keyvals[i].(string), keyvals[i+1])
	}
	event.Msg(msg)
}

func (s *Zerolog) Warn(msg string, keyvals ...any) {
	if s.log == nil {
		return
	}
	event := s.log.Warn()
	for i := 0; i < len(keyvals); i += 2 {
		event = event.Any(keyvals[i].(string), keyvals[i+1])
	}
	event.Msg(msg)
}

func (s *Zerolog) Error(msg string, keyvals ...any) {
	if s.log == nil {
		return
	}
	event := s.log.Error()
	for i := 0; i < len(keyvals); i += 2 {
		event = event.Any(keyvals[i].(string), keyvals[i+1])
	}
	event.Msg(msg)
}

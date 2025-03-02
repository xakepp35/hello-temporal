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

func (s *Zerolog) Debug(msg string, keyvals ...interface{}) {
	if s.log == nil {
		return
	}
	for i := 0; i < len(keyvals); i += 2 {
		s.log.Debug().Any(keyvals[0].(string), keyvals[1])
	}
}

func (s *Zerolog) Info(msg string, keyvals ...interface{}) {
	if s.log == nil {
		return
	}
	for i := 0; i < len(keyvals); i += 2 {
		s.log.Info().Any(keyvals[0].(string), keyvals[1])
	}
}

func (s *Zerolog) Warn(msg string, keyvals ...interface{}) {
	if s.log == nil {
		return
	}
	for i := 0; i < len(keyvals); i += 2 {
		s.log.Warn().Any(keyvals[0].(string), keyvals[1])
	}
}

func (s *Zerolog) Error(msg string, keyvals ...interface{}) {
	if s.log == nil {
		return
	}
	for i := 0; i < len(keyvals); i += 2 {
		s.log.Error().Any(keyvals[0].(string), keyvals[1])
	}
}

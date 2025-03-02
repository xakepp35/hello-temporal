package activities

import (
	"context"

	"github.com/rs/zerolog/log"
)

// HelloWorld: печатает Hello, world
func HelloWorld(ctx context.Context) error {
	log.Info().Msg("Hello, world")
	return nil
}

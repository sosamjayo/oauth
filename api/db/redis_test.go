package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRedisClient(t *testing.T) {
	t.Run("Test NewRedisClient", func(t *testing.T) {
		ctx := context.Background()
		_, err := NewRedisClient(ctx)
		assert.NoError(t, err)
	})
}

package db

import "github.com/redis/go-redis/v9"

type ClientDB interface {
}

// ClientRedis is a struct that contains a redis client
// and implements the ClientDB interface
type ClientRedis struct {
	Client *redis.Client
}

func NewClientRedis(c *redis.Client) ClientDB {
	return &ClientRedis{
		Client: c,
	}
}

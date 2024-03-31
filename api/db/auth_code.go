package db

import "github.com/redis/go-redis/v9"

type AuthCodeDB interface {
}

// AuthCodeRedis is a struct that contains a redis client
// and implements the AuthCodeDB interface
type AuthCodeRedis struct {
	Client *redis.Client
}

func NewAuthCodeRedis(c *redis.Client) AuthCodeDB {
	return &AuthCodeRedis{
		Client: c,
	}
}

package db

import "github.com/redis/go-redis/v9"

type AccessTokenDB interface {
}

// AccessTokenRedis is a struct that contains a redis client
// and implements the AccessTokenDB interface
type AccessTokenRedis struct {
	Client *redis.Client
}

func NewAccessTokenRedis(c *redis.Client) AccessTokenDB {
	return &AccessTokenRedis{
		Client: c,
	}
}

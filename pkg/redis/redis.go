package redis

import (
	"github.com/go-redis/redis/v8"
)

func New(cfg *Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address[0],
		Password: cfg.Password, // no password set
		DB:       cfg.Db,       // use default DB
	})
	return rdb, nil
}

func NewCluster(cfg *Config) (*redis.ClusterClient, error) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Address,
		Password: cfg.Password,
	})
	return rdb, nil
}

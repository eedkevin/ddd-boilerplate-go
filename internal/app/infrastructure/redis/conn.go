package redis

import (
	"context"
	"log"
	"strings"

	"ddd-boilerplate/internal/app/adapter/cfg"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

//Client is the alias of redis.Client
type Client = redis.Client

//ClusterClient is the alias of redis.ClusterClient
type ClusterClient = redis.ClusterClient

//Connection creates a redis client
func Connection(cfg cfg.RedisConfig) *Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Username: cfg.User,
		Password: cfg.Pass,
		DB:       cfg.DB,
	})
	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatal("Error on connection to redis", err)
	}
	return client
}

//ConnectionCluster creates a redis cluster client
func ConnectionCluster(cfg cfg.RedisConfig) *ClusterClient {
	addrs := strings.Split(cfg.Cluster.Addrs, ",")
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addrs,
		Username: cfg.User,
		Password: cfg.Pass,
	})
	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatal("Error on connection to redis cluster", err)
	}
	return client
}

//Pop returns a random uuid string
func Pop() string {
	return uuid.NewString()
}

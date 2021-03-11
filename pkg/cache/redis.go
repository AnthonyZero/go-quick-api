package cache

import (
	"time"

	"github.com/anthonyzero/go-quick-api/configs"
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
)

var repo *RedisRepo

//RedisRepo repo
type RedisRepo struct {
	client *redis.Client
}

//InitRedis init redis pool
func InitRedis() error {
	client, err := connectRedis()
	if err != nil {
		return err
	}
	repo = &RedisRepo{
		client: client,
	}
	return nil
}

func connectRedis() (*redis.Client, error) {
	cfg := configs.Get().Redis
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Pass,
		DB:           cfg.Db,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, errors.Wrap(err, "ping redis error")
	}

	return client, nil
}

//Client 获取redis client
func Client() *redis.Client {
	return repo.client
}

// Set set some <key,value> into redis
func Set(key, value string, ttl time.Duration) error {
	if err := repo.client.Set(key, value, ttl).Err(); err != nil {
		return errors.Wrapf(err, "redis set key: %s err", key)
	}

	return nil
}

// Get get some key from redis
func Get(key string) (string, error) {
	value, err := repo.client.Get(key).Result()
	if err != nil {
		return "", errors.Wrapf(err, "redis get key: %s err", key)
	}

	return value, nil
}

// TTL get some key from redis
func TTL(key string) (time.Duration, error) {
	ttl, err := repo.client.TTL(key).Result()
	if err != nil {
		return -1, errors.Wrapf(err, "redis get key: %s err", key)
	}

	return ttl, nil
}

// Expire expire some key
func Expire(key string, ttl time.Duration) bool {
	ok, _ := repo.client.Expire(key, ttl).Result()
	return ok
}

// ExpireAt expire some key at some time
func ExpireAt(key string, ttl time.Time) bool {
	ok, _ := repo.client.ExpireAt(key, ttl).Result()
	return ok
}

// Del del some key from redis
func Del(keys ...string) bool {
	if len(keys) == 0 {
		return true
	}

	value, _ := repo.client.Del(keys...).Result()
	return value > 0
}

//Incr incr key
func Incr(key string) int64 {
	value, _ := repo.client.Incr(key).Result()
	return value
}

// Close close redis client
func Close() error {
	return repo.client.Close()
}

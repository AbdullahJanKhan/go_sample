package redis

import (
	"time"

	"github.com/abdullahjankhan/go_sample/config"
	"github.com/abdullahjankhan/go_sample/models"
	"github.com/go-redis/redis"
)

type Client interface {
	Set(key string, values interface{}) error
	Get(key string) (string, error)
	Del(key string) error
}

type client struct {
	client *redis.Client
}

func NewClient(conf *config.GlobalConfig) Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       1,
	})
	if redisClient == nil {
		panic("failed to connect to redis")
	}

	return &client{
		client: redisClient,
	}
}

func (r *client) Set(key string, values interface{}) error {
	err := r.client.Set(key, values, time.Minute*2).Err()
	if err != nil {
		return &models.StandardError{
			Code:        models.INTERNAL_SERVER_ERROR,
			ActualError: nil,
			Line:        "Set():19",
			Message:     models.REDIS_SET_ERROR_MESSAGE,
		}
	}
	return nil
}

func (r *client) Get(key string) (string, error) {
	res := r.client.Get(key)
	if res.Err() == redis.Nil {
		return "", nil

	} else if res.Val() != "" {
		return res.Val(), nil
	}
	return "", &models.StandardError{
		Code:        models.INTERNAL_SERVER_ERROR,
		ActualError: nil,
		Line:        "Get():19",
		Message:     models.REDIS_GET_ERROR_MESSAGE,
	}
}

func (r *client) Del(key string) error {
	err := r.client.Del(key).Err()
	if err != nil {
		return &models.StandardError{
			Code:        models.INTERNAL_SERVER_ERROR,
			ActualError: nil,
			Line:        "Del():64",
			Message:     models.REDIS_DEL_ERROR_MESSAGE,
		}
	}
	return nil
}

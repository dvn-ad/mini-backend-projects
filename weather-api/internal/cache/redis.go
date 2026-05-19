package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct{
	client *redis.Client
}

func New(redisURL string)(*Cache,error){
	opts,err:=redis.ParseURL(redisURL)
	if err!=nil{
		opts=&redis.Options{Addr: redisURL}
	}
	client:=redis.NewClient(opts)

	ctx,cancel:=context.WithTimeout(context.Background(),3*time.Second)
	defer cancel()

	if err:=client.Ping(ctx).Err();err!=nil{
		return nil, err
	}
	return &Cache{client: client}, nil
}

func(c *Cache) Get(ctx context.Context, key string)(string,error){
	return c.client.Get(ctx,key).Result()
}

func(c *Cache) Set(ctx context.Context, key string,value string,ttl time.Duration) error{
	return c.client.Set(ctx,key,value,ttl).Err()
}

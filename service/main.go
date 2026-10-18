package service

import (
	"context"
	redis "github.com/redis/go-redis/v9"
)

type URLService struct {
	redisClient *redis.Client
}

func NewService(redisClient *redis.Client) *URLService {
	return &URLService{
		redisClient: redisClient,
	}
}
func (s *URLService) ShortenURL(longURL string) string {
	return "http://short.url/abc123"
}

func(s *URLService) GetLongURL(ctx context.Context, shortURL string) string{
	val,err:=s.redisClient.Get(ctx,shortURL).Result()
	if err==nil{
		return val
	}
	return ""
}
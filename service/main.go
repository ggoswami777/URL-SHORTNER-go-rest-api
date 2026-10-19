package service

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type URLService struct {
	redisClient *redis.Client
}

type ShortURLRedis struct{
	ShortURL string
	LongURL string
	CreatedAt int64
	
}

func NewService(redisClient *redis.Client) *URLService {
	return &URLService{
		redisClient: redisClient,
	}
}
func URLShortnerAlgorithm(longURL string) string{
	return "abc123"
}
func (s *URLService) ShortenURL(ctx context.Context,longURL string) (string,error){
	shortUrl:=URLShortnerAlgorithm(longURL)
	shortURLRedis:=&ShortURLRedis{
		ShortURL: shortUrl,
		LongURL: longURL,
		CreatedAt: time.Now().Unix(),
	}

	err:=s.redisClient.Set(ctx,shortUrl,shortURLRedis,0).Err()
	if err!=nil{
		return "",err
	}

	return "http://short.url/abc123",nil
}

func(s *URLService) GetLongURL(ctx context.Context, shortURL string) string{
	val,err:=s.redisClient.Get(ctx,shortURL).Result()
	if err==nil{
		return val
	}
	return ""
}
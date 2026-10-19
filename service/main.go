package service

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type URLService struct {
	redisClient *redis.Client
}

type ShortURLRedis struct{
	Shorten string
	LongURL string
	CreatedAt int64
	
}

func NewService(redisClient *redis.Client) *URLService {
	return &URLService{
		redisClient: redisClient,
	}
}

const charset="abcdefghijklmnopqrstuvwxyz1234567890"
var seededRand=rand.New(rand.NewSource(time.Now().UnixNano()))
func URLShortnerAlgorithm() string{
	const length=6
	b:=make([]byte,length)
	for i:=range b{
		b[i]=charset[seededRand.Intn(len(charset))]
	}
	return string (b)
}
func (s *URLService) ShortenURL(ctx context.Context,longURL string) (string,error){
	shorten:=URLShortnerAlgorithm()
	
	shortURLRedis:=&ShortURLRedis{
		Shorten:shorten,
		LongURL: longURL,
		CreatedAt: time.Now().Unix(),
	}
	
	// marshal and store in redis
	shortURLMarshal,err:=json.Marshal(shortURLRedis)
	if err!=nil{
		return "",err
	}
	err=s.redisClient.Set(ctx,shorten,shortURLMarshal,0).Err()
	if err!=nil{
		return "",err
	}
	shortURL:= "http://localhost:3000/"+shorten
	return shortURL,nil
}

func(s *URLService) GetLongURL(ctx context.Context, shortURL string) (string,error){
	val,err:=s.redisClient.Get(ctx,shortURL).Result()
	
	if err==redis.Nil{
		return "",nil
	}else if err!=nil{
		return "",err
	}
	var shortURLRedis ShortURLRedis

	err=json.Unmarshal([]byte(val),&shortURLRedis)
	if err==nil{
		return shortURLRedis.LongURL,nil
	}
	return "",err
}
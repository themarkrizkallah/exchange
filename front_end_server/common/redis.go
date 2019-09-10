package common

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/go-redis/redis"

	"front_end_server/env"
)

var redisClient *redis.Client

// GenerateRandomBytes returns securely generated random bytes.
func generateRandomBytes(n uint) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(n uint) (string, error) {
	b, err := generateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func InitRedis() *redis.Client {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     env.RedisHostName + ":" + env.RedisPort,
		Password: env.RedisPassword,
		DB:       0, // use default DB
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Panicln("Error setting up Redis:", err)
	}

	return redisClient
}

func GetRedisClient() *redis.Client {
	return redisClient
}
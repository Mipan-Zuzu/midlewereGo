package utils

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"time"
	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

var RDB *redis.Client

func GenerateToken(userID uint) (string, error) {
	if os.Getenv("JWT_SECRET") == "" {
		panic("missing jwt secret token")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func ComperePassword(findUser string, inputuser string) error  {
	data := bcrypt.CompareHashAndPassword([]byte(findUser), []byte(inputuser))
	return data
}

func RedishConnection () *redis.Client{

	RDB = redis.NewClient(&redis.Options{
		Addr: os.Getenv("UPSTASH_REDIS_REST_URL"),
		Password: os.Getenv("UPSTASH_REDIS_REST_TOKEN"),
		TLSConfig: &tls.Config{},
	})
	
	ctx := context.Background()
	_,err := RDB.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect redis : %v", err))
	}
	return RDB
}
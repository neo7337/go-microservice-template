package cache

import (
	"context"
	"fmt"

	"strconv"

	"github.com/neo7337/go-microservice-template/internal/config"
	"github.com/redis/go-redis/v9"
	"oss.nandlabs.io/golly/l3"
	"oss.nandlabs.io/golly/lifecycle"
)

var rdb *redis.Client

var logger = l3.Get()

func GetConnection(manager lifecycle.ComponentManager) *lifecycle.SimpleComponent {
	return &lifecycle.SimpleComponent{
		CompId: "redis-connection",
		StartFunc: func() error {
			return InitRedis()
		},
		AfterStart: func(err error) {
			if err != nil {
				logger.Error("Failed to initialize Redis connection", "error", err)
				manager.StopAll()
			} else {
				logger.Info("Redis connection initialized successfully")
			}
		},
	}
}

func InitRedis() (err error) {
	configDetails := config.GetConfig()
	addr := configDetails.Cache.Config.Host + ":" + strconv.Itoa(configDetails.Cache.Config.Port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,                                // Redis server address
		Password: configDetails.Cache.Config.Password, // no password set
		DB:       configDetails.Cache.Config.Db,       // use default DB
	})
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return
	}
	logger.InfoF("Redis connected successfully: %s", pong)
	return
}

func GetRedisClient() *redis.Client {
	if rdb == nil {
		logger.Error("Redis client is not initialized. Call InitRedis first.")
		return nil
	}
	return rdb
}

func GetFlightOfferFromRedis(provider, flightOfferId string) (string, error) {
	key := provider + ":" + flightOfferId
	fmt.Println("Retrieving flight offer from Redis with key:", key)
	members, err := rdb.SMembers(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			logger.InfoF("Flight offer %s not found in Redis for provider %s", flightOfferId, provider)
			return "", nil
		}
		logger.ErrorF("Error retrieving flight offer %s from Redis for provider %s: %v", flightOfferId, provider, err)
		return "", err
	}

	if isNumber(members[0]) {
		return members[1], nil
	} else {
		return members[0], nil
	}
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

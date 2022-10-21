package bootstrappers

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	limiter "github.com/ulule/limiter/v3"
	LimiterMiddleware "github.com/ulule/limiter/v3/drivers/middleware/gin"
	RedisStore "github.com/ulule/limiter/v3/drivers/store/redis"
)

var RateLimiterMiddleware gin.HandlerFunc

func BootstrapRateLimiter() {

	// Define a limit rate to 4 requests per hour.
	rate, err := limiter.NewRateFromFormatted("20-M")
	if err != nil {
		log.Fatal(err)
		return
	}

	redisDBNo, error := strconv.Atoi(os.Getenv("REDIS_DATABASE_NO"))

	if error != nil {
		log.Fatal("Rate Limiter bootstrap failed ", error)
		return
	}

	// Create a redis client.
	option := redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDBNo,
	}

	client := redis.NewClient(&option)

	// Create a store with the redis client.
	store, err := RedisStore.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix:   "spotube-matcher",
		MaxRetry: 3,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create a new middleware with the limiter instance.
	RateLimiterMiddleware = LimiterMiddleware.NewMiddleware(limiter.New(store, rate))
}

## 获取reids连接
```

import (
	"github.com/go-redis/redis"
)

func GetRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	return rdb
}
```

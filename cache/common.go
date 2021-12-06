package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"strconv"
)

// RedisClient Redis缓存客户端单例
var (
	RedisClient  *redis.Client
	RedisDb    			string
	RedisAddr  			string
	RedisPw    			string
	RedisDbName    		string
)

// Redis 在中间件中初始化redis链接  防止循环导包，所以放在这里
func init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("Redis 配置文件读取错误，请检查文件路径:", err)
	}
	LoadRedisData(file)
	Redis()
}

//Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64) 		//TODO 这里记得了！！
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		//Password: conf.RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}


func LoadRedisData(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").MustString("redis")
	RedisAddr = file.Section("redis").Key("RedisAddr").MustString("127.0.0.1:6379")
	RedisPw = file.Section("redis").Key("RedisPw").MustString("root")
	RedisDbName = file.Section("redis").Key("RedisDbName").MustString("2")
}
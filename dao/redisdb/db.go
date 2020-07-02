package redisdb

import (
	"sz_resume_202005/utils/setting"
	"sz_resume_202005/utils/zlog"
	"time"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

//Init 初始化redis
func Init() {
	err := initClient()
	if err != nil {
		zlog.Panic(err)
	}

}

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     setting.RedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	return
}

//Set 设置字符串  t=0代表没有过期时间
func Set(key, value string, t int) error {
	err := rdb.Set(key, value, time.Duration(t)*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

//Get 获取字符串
func Get(key string) (val string, exist bool, err error) {
	val, err = rdb.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", false, nil
		}
		return "", false, err
	}
	return val, true, nil
}

//GetTTL 获取过期时间
func GetTTL(key string) (tm time.Duration, err error) {
	tm, err = rdb.TTL(key).Result()
	return
}

//SetNX 不存在才设置
func SetNX(key string, value string, t int) (val bool, err error) {
	rdb.SetNX("counter", 0, time.Duration(t)*time.Second).Result()
	val, err = rdb.SetNX(key, value, time.Duration(t)*time.Second).Result()
	return
}

//Incr 指定key自增
func Incr(key string) (result int64, err error) {
	return rdb.Incr(key).Result()
}

//Del 删除多个键
func Del(keys ...string) (int64, error) {
	return rdb.Del(keys...).Result()

}

//Exist 指定key是否存在
func Exist(keys ...string) (int64, error) {
	return rdb.Exists(keys...).Result()
}

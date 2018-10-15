package gameutil

import (
	"github.com/go-redis/redis"
	"github.com/kingbuffalo/seelog"
	"buffalo/king/common/startcfg"
)

var client *redis.Client = nil

func init() {
	h, p, i := startcfg.GetGRedisHostPasswd()

	client = redis.NewClient(&redis.Options{
		Addr:     h,
		Password: p,
		DB:       i,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	//set max pid
}

////////////////////////////game redis begin

func SAddSetKey(key string, setValue interface{}) error {
	return client.SAdd(key, setValue).Err()
}

func HGet(key, field string) []byte {
	ret, err := client.HGet(key, field).Bytes()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		_ = seelog.Warnf("RedisError_HGet:key=%s,field=%s,err=%s", key, field, err.Error())
	}
	return ret
}

func HGetAll(key string) map[string]string {
	ret, err := client.HGetAll(key).Result()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		_ = seelog.Warnf("RedisError_HGetAll:key=%s,err=%s", key, err.Error())
		return nil
	}
	if len(ret) == 0 {
		return nil
	}
	return ret
}

func HGetAllErr(key string) (map[string]string, error) {
	ret, err := client.HGetAll(key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		_ = seelog.Warnf("RedisError_HGetAll:key=%s,err=%s", key, err.Error())
		return nil, err
	}
	if len(ret) == 0 {
		return nil, nil
	}
	return ret, nil
}

func HDEL(key string, field string) {
	if err := client.HDel(key, field).Err(); err != nil {
		_ = seelog.Warnf("RedisError_HDel:key=%s,field=%s,err=%s", key, field, err.Error())
	}
}
func HSetAllType(key, field string, value interface{}) {
	if err := client.HSet(key, field, value).Err(); err != nil {
		_ = seelog.Warnf("RedisError_HDel:key=%s,field=%s,value=%v,err=%s", key, field, value, err.Error())
	}
}
func HSet(key, field string, value []byte) {
	if err := client.HSet(key, field, value).Err(); err != nil {
		_ = seelog.Warnf("RedisError_HDel:key=%s,field=%s,value=%s,err=%s", key, field, string(value), err.Error())
	}
}

func SAdd(key string, value ...interface{}) {
	if err := client.SAdd(key, value...).Err(); err != nil {
		_ = seelog.Warnf("RedisError_SAdd:key=%s,err=%s", key, err.Error())
	}
}

func ZAdd(key string, score float64, member interface{}) {
	z := redis.Z{
		Score:  score,
		Member: member,
	}
	if err := client.ZAdd(key, z).Err(); err != nil {
		_ = seelog.Warnf("RedisError_ZAdd:key=%s,err=%s", key, err.Error())
	}
}

func SRandMemberN(key string, num int64) []string {
	v := client.SRandMemberN(key, num)
	strArr, err := v.Result()
	if err != nil {
		_ = seelog.Warnf("RedisError_SrandMember:key=%s,num=%d,err=%s", key, num, err.Error())
	}
	return strArr
}

func GetSetKeys(key string) []string {
	tmp := client.SMembers(key)
	ret, err := tmp.Result()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		_ = seelog.Warnf("RedisError_GetSetKeys:key=%s,err=%s", key, err.Error())
	}
	return ret
}

func GGet(key string) []byte {
	tmp := client.Get(key)
	b, err := tmp.Bytes()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		_ = seelog.Warnf("RedisError_GGet:key=%s,err=%s", key, err.Error())
	}
	return b
}

func GSet(key string, value []byte) {
	if err := client.Set(key, value, 0).Err(); err != nil {
		_ = seelog.Warnf("RedisError_GSet:key=%s,value=%s,err=%s", key, string(value), err.Error())
	}
}

func GSetAllType(key string, value interface{}) {
	if err := client.Set(key, value, 0).Err(); err != nil {
		_ = seelog.Warnf("RedisError_GSetAllType:key=%s,value=%v,err=%s", key, value, err.Error())
	}
}

func GExists(key string) bool {
	i64, err := client.Exists(key).Result()
	if err != nil {
		_ = seelog.Warnf("RedisError_GExists:key=%s,err=%s", key, err.Error())
	}
	return i64 == 1
}

func GDelete(key string) {
	if err := client.Del(key).Err(); err != nil {
		_ = seelog.Warnf("RedisError_GDelete:key=%s,err=%s", key, err.Error())
	}
}

func GGetKeys(pattern string) []string {
	ret, err := client.Keys(pattern).Result()
	if err != nil {
		_ = seelog.Warnf("RedisError_GGetKeys:pattern=%s,err=%s", pattern, err.Error())
	}
	return ret
}

func GIncr(counterKey string) int64 {
	ret, err := client.Incr(counterKey).Result()
	if err != nil {
		_ = seelog.Warnf("RedisError_GIncr:counterKey=%s,err=%s", counterKey, err.Error())
	}
	return ret
}

func HIncrBy(key, field string, incr int64) int64 {
	ret, err := client.HIncrBy(key, field, incr).Result()
	if err != nil {
		_ = seelog.Warnf("RedisError_HIncrBy:key=%s,field=%s,incr=%d,err=%s", key, field, incr, err.Error())
	}
	return ret
}

func CloseRedis() error {
	return client.Close()
}

////////////////////////////game redis end

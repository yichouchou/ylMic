package tool

import (
	red "github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

type Redis struct {
	pool *red.Pool
}

var redisPool *Redis

var one sync.Once

func InitRedisPool() {
	one.Do(func() {
		config := GetConfig().RedisConfig
		redisPool.pool = &red.Pool{
			MaxIdle:     256,
			MaxActive:   0,
			IdleTimeout: time.Duration(120),
			Dial: func() (red.Conn, error) {
				return red.Dial(
					config.NetWork,
					config.Addr+":"+config.Port,
					red.DialReadTimeout(time.Duration(config.DialReadTimeout)*time.Millisecond),
					red.DialWriteTimeout(time.Duration(config.DialWriteTimeout)*time.Millisecond),
					red.DialConnectTimeout(time.Duration(config.DialConnectTimeout)*time.Millisecond),
					red.DialDatabase(config.DialDatabase),
					//red.DialPassword(""),
				)
			},
		}
	})
}

func Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	con := redisPool.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	parmas := make([]interface{}, 0)
	parmas = append(parmas, key)

	if len(args) > 0 {
		for _, v := range args {
			parmas = append(parmas, v)
		}
	}
	return con.Do(cmd, parmas...)
}

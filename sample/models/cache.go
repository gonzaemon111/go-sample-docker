package models

// https://qiita.com/chan-p/items/5c5e7cc1e966f8a90422

import (
	"github.com/garyburd/redigo/redis"
)

func redis_connection() redis.Conn {
  const IP_PORT = "redis://go_sample_redis:6379"

  //redisに接続
  c, err := redis.Dial("tcp", IP_PORT)
  if err != nil {
    panic(err)
  }
  return c
}
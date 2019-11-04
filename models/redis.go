package models

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// Redis struct .
type Redis struct {
	DataBase
	c redis.Conn
}

// Init redis .
func (r *Redis) Init() {
	var err error
	r.c, err = redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println("connect to redis error:", err)
	}
}

// AuthenticateUser .
func (r *Redis) AuthenticateUser(userID, pwd string) bool {
	result, _ := r.c.Do("HGET", "user:"+userID, "pwd")

	if result == nil {
		return false
	}

	data := result.([]byte)

	return pwd == string(data[:])
}

package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect to redis error:", err)
	} else {
		fmt.Println("connect to redis success!")
	}

	defer c.Close()
}

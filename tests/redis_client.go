package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"os"
	"reflect"

	"encoding/hex"

	"github.com/garyburd/redigo/redis"
)

func welcome() {
	fmt.Println("##### welcome to redis client. #####")
	fmt.Println("##### 1. add user. #####")
	fmt.Println("##### 2. find user. #####")
	fmt.Println("##### 3. find message. #####")
	fmt.Println("##### 0. quit. #####")
}

func addUser(c *redis.Conn) {
	var name string
	var userID string
	var pwd string

	fmt.Printf("userId:")
	fmt.Scanf("%s", &userID)

	result, err := (*c).Do("HGET", "user:"+userID, "name")

	if err != nil || result != nil {
		fmt.Println("Add user", userID, "failed!")
		return
	}

	fmt.Printf("name:")
	fmt.Scanf("%s", &name)

	fmt.Printf("pwd:")
	fmt.Scanf("%s", &pwd)

	h := sha1.New()

	h.Write([]byte(pwd))
	encrypPwd := hex.EncodeToString(h.Sum(nil))
	h.Reset()

	_, err1 := (*c).Do("HSET", "user:"+userID, "name", name, "pwd", encrypPwd)

	if err1 != nil {
		fmt.Println("Add user", userID, "failed!")
	} else {
		fmt.Println("Add user", userID, "success!")
	}
}

func findUser(c *redis.Conn) {
	var userID string
	fmt.Printf("userId:")
	fmt.Scanf("%s", &userID)

	result, err := (*c).Do("HGET", "user:"+userID, "name")

	if err != nil {
		fmt.Println("Process find user", userID, "failed!")
	} else if result != nil {
		var name string
		data := result.([]byte)
		name = string(data[:])
		fmt.Println("Find user", userID, "success!")
		fmt.Printf("name: %s\n", name)
	} else {
		fmt.Println("Find user", userID, "failed!")
	}
}

func findMessage(c *redis.Conn) {
	var messageID string
	fmt.Printf("message_id:")
	fmt.Scanf("%s", &messageID)

	result, err := redis.StringMap((*c).Do("HGETALL", "message:"+messageID))

	if err != nil {
		fmt.Println("Process find user", messageID, "failed!")
	} else if result != nil {
		for key, value := range result {
			switch key {
			case "user_id":
				fmt.Println("用户名:", value, reflect.TypeOf(value))
			case "text":
				fmt.Println("文章内容:", value, reflect.TypeOf(value))
			case "time":
				fmt.Println("发布时间:", value, reflect.TypeOf(value))
			}
			fmt.Println(key, value)
		}
	} else {
		fmt.Println("Find user", messageID, "failed!")
	}
}

func process(c *redis.Conn) {
	w := 1

	for w != 0 {
		welcome()
		bufio.NewReader(os.Stdin)
		fmt.Scanf("%d", &w)

		switch w {
		case 1:
			addUser(c)
			break
		case 2:
			findUser(c)
			break
		case 3:
			findMessage(c)
		}
	}
}

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect to redis error:", err)
	} else {
		fmt.Println("connect to redis success!")
		process(&c)
	}

	defer c.Close()
}

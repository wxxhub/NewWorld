package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"

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

// AddUser .
func (r *Redis) AddUser(userID, name, pwd string) AddStatus {
	result, err := r.c.Do("HGET", "user:"+userID, "pwd")

	if err != nil || result != nil {
		return HaveExist
	}

	_, err2 := r.c.Do("HSET", "user:"+userID, "name", name, "pwd", pwd)

	if err2 != nil {
		return AddFaile
	}
	return AddSuccess
}

// AddMessage .
func (r *Redis) AddMessage(userID, text, image string) AddStatus {
	// var messageCounter uint64
	messageCounter, err := redis.Int64(r.c.Do("GET", "message_counter"))

	if err != nil {
		logs.Warn("Add message failed 1!")
		return AddFaile
	}

	messageIndex := fmt.Sprintf("message:%d", messageCounter)
	now := time.Now().Format("2006-01-02 15:04:05")
	r.c.Send("MULTI")
	r.c.Send("INCR", "message_counter")
	r.c.Send("LPUSH", "message_list:"+userID, messageCounter)
	r.c.Send("HSET", messageIndex,
		"user_id", userID,
		"text", text,
		"time", now,
		"praise", 0,
		"image", image)
	_, err1 := r.c.Do("EXEC")

	if err1 != nil {
		logs.Warn("Add message failed 3!")
		return AddFaile
	}

	logs.Info("Add message success.")
	return AddSuccess
}

// AddCommit .
func (r *Redis) AddCommit(messageID, userID, commit string) AddStatus {
	var info CommitInfo
	info.UserID = userID
	info.Commit = commit
	codeData, err := json.Marshal(info)
	if err != nil {
		return AddFaile
	}
	_, err1 := r.c.Do("LPUSH", "commits:"+messageID, string(codeData))
	if err1 != nil {
		return AddFaile
	}
	return AddSuccess
}

// AddConcern .
func (r *Redis) AddConcern(currentUserID, goalUserID string) AddStatus {
	_, err := r.c.Do("LPUSH", "user_concern:"+currentUserID, goalUserID)

	if err != nil {
		return AddFaile
	}
	return AddSuccess
}

// AddPraise .
func (r *Redis) AddPraise(messageID, userID string) AddStatus {
	result, err := redis.Int(r.c.Do("SADD", "praise_set:"+messageID, userID))

	if err != nil {
		logs.Warn("add praise AddFaile!")
		return AddFaile
	}

	if result == 0 {
		logs.Warn("add praise HaveExist!")
		return HaveExist
	}

	logs.Info("AddPraise Success!")
	return AddSuccess
}

// CancelPraise .
func (r *Redis) CancelPraise(messageID, userID string) AddStatus {
	result, err := redis.Int(r.c.Do("SISMEMBER", "praise_set:"+messageID, userID))

	if err != nil {
		logs.Warn("add praise AddFaile!")
		return AddFaile
	}

	if result == 0 {
		logs.Warn("add praise HaveExist!")
		return HaveExist
	}

	logs.Info("AddPraise Success!")
	return AddSuccess
}

// GetMessage .
func (r *Redis) GetMessage(messageID string) (Message, bool) {
	var message Message
	messageData, err := redis.StringMap(r.c.Do("HGETALL", messageID))

	if err != nil {
		return message, false
	}

	// 获取消息基本信息
	for key, value := range messageData {
		switch key {
		case "user_id":
			message.UserID = value
		case "text":
			message.Text = value
		case "time":
			message.Time = value
		}
	}

	// 获取消息评论
	commitSize, err1 := redis.Int(r.c.Do("LLEN", "commits:"+messageID))

	if err1 != nil {
		return message, false
	}

	commits, err2 := redis.Strings(r.c.Do("LRANGE", "commits:"+messageID, 0, commitSize))

	if err2 != nil {
		return message, false
	}

	// 获取消息点赞数
	praiseNum, err3 := redis.Int(r.c.Do("SCARD", "praise_set:"+messageID))

	if err3 != nil {
		return message, false
	}

	message.Praise = praiseNum

	if commitSize > 0 {
		message.Commit = make([]CommitInfo, commitSize)
		for i, commit := range commits {
			json.Unmarshal([]byte(commit), message.Commit[i])
		}
	}
	return message, true
}

// GetConcern .
func (r *Redis) GetConcern(userID string) []string {
	var concerns []string
	return concerns
}

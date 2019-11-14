package models

import (
	list "container/list"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/garyburd/redigo/redis"
)

// Redis struct .
type Redis struct {
	DataBase
	c redis.Conn
}

// Hot .
type Hot struct {
	HotCache1 []Message // 缓存
	HotCache2 []Message
	Switch    bool    // 1 true, 2 false
	Size      int     // 缓存容量
	MaxSize   int     // 最大容量, 避免数据过大
	Hour      float64 // 满足热度的时间限制(小时)
	Score     uint64  // 满足热度的分数
	rd        *Redis  // 数据库

	UpdateLock sync.Mutex // 更新锁
}

// HotManager .
var HotManager Hot

// Init .
func (h *Hot) Init() {
	h.HotCache1 = h.Process()
	h.Switch = true
	logs.Info("hot Init!")
}

// Update .
func (h *Hot) Update() {
	h.UpdateLock.Lock()
	defer h.UpdateLock.Unlock()
	if h.Switch { // 如果正在使用缓存1, 则更新缓存2, 并切换到缓存2
		h.HotCache2 = h.Process()
	} else {
		h.HotCache1 = h.Process()
	}

	logs.Info("hot update!")
	h.Switch = !h.Switch
}

// CalScore .
func (h *Hot) CalScore(commit, praise uint64) uint64 {
	return 2*commit + praise
}

// GetHotMessage .
func (h *Hot) GetHotMessage() []Message {
	if h.Switch {
		return h.HotCache1
	}

	return h.HotCache2
}

// Process .
func (h *Hot) Process() []Message {
	outTimeCache := make([]Message, 0)
	cache := make([]Message, 0)
	counter, err := redis.Int64(h.rd.c.Do("GET", "message_counter"))

	if err != nil {
		logs.Error("Find message_counter failed!")
	}

	now := time.Now()

	outTimeCacheIndex := 0
	cacheIndex := 0
	for index := counter - 1; index >= 0; index-- {

		message, ok := h.rd.GetMessage(strconv.FormatInt(index, 10))

		// 是否存在和计算分数是否满足
		if !ok || h.CalScore(uint64(len(message.Commit)), message.Praise) < h.Score {
			continue
		}

		messageTime, _ := time.Parse("2016-01-02 15:04:05", message.Time)
		du := now.Sub(messageTime)

		message.UserName, _ = redis.String(h.rd.c.Do("HGET", "user:"+message.UserID, "name"))
		message.MessageID = strconv.FormatInt(index, 10)
		// 判断是否超时
		if du.Hours() < h.Hour {
			if cacheIndex > h.MaxSize {
				break
			}
			cache = append(cache, message)
		} else if outTimeCacheIndex < h.Size { // 如果超时, 但超时缓存没满, 就放进超时缓存
			if outTimeCacheIndex < h.Size {
				outTimeCache = append(outTimeCache, message)
				outTimeCacheIndex++
			}
		} else if cacheIndex >= h.Size { // 如果超时且容量满了, 退出
			break
		}
	}

	// 如果没装满, 则将超时的加入
	if len(cache) < h.Size {
		cache = append(cache, outTimeCache...)
	}

	return cache
}

// Init redis .
func (r *Redis) Init() {
	var err error
	r.c, err = redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println("connect to redis error:", err)
	}

	HotManager = Hot{
		HotCache1: make([]Message, 0),
		HotCache2: make([]Message, 0),
		Switch:    true,
		Size:      20,
		MaxSize:   100,
		Hour:      2,
		Score:     3,
		rd:        r,
	}
}

// AuthenticateUser .
func (r *Redis) AuthenticateUser(userID, pwd string) (userName, head string, ok bool) {
	messageData, err := redis.StringMap(r.c.Do("HGETALL", "user:"+userID))

	if err != nil {
		return "", "", false
	}

	// 获取消息基本信息
	// var findPwd string
	// var findName string
	// var findHead string

	// for key, value := range messageData {
	// 	switch key {
	// 	case "pwd":
	// 		findPwd = value
	// 	case "name":
	// 		findName = value
	// 	case "image":
	// 		findHead = value
	// 	}
	// }

	// 获取消息基本信息
	findPwd := messageData["pwd"]
	findName := messageData["name"]
	findHead := messageData["image"]

	if findPwd != pwd {
		return "", "", false
	}

	return findName, findHead, true
}

// AddUser .
func (r *Redis) AddUser(userID, name, pwd, image string) ProcessStatus {
	result, err := r.c.Do("HGET", "user:"+userID, "pwd")

	if err != nil || result != nil {
		return HAVEEXIST
	}

	_, err2 := r.c.Do("HSET", "user:"+userID, "name", name, "pwd", pwd, "image", image)

	if err2 != nil {
		return FAILED
	}
	return SUCCESS
}

// AddMessage .
func (r *Redis) AddMessage(userID, text, image string) ProcessStatus {
	// var messageCounter uint64
	messageCounter, err := redis.Int64(r.c.Do("GET", "message_counter"))

	if err != nil {
		logs.Warn("Add message failed 1!")
		return FAILED
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
		return FAILED
	}

	logs.Info("Add message success.")
	return SUCCESS
}

// AddCommit .
func (r *Redis) AddCommit(messageID, userID, commit string) ProcessStatus {
	var info CommitInfo
	// fmt.Println(messageID, userID, commit)
	info.UserID = userID
	info.Commit = commit
	codeData, err := json.Marshal(info)
	if err != nil {
		logs.Warn("Marshal commit failed!")
		return FAILED
	}
	_, err1 := r.c.Do("LPUSH", "commits:"+messageID, string(codeData))
	if err1 != nil {
		logs.Warn("add commit failed!")
		return FAILED
	}

	logs.Info("add commit success!")
	return SUCCESS
}

// AddConcern .
func (r *Redis) AddConcern(currentUserID, goalUserID string) ProcessStatus {
	_, err := r.c.Do("SADD", "user_concern:"+currentUserID, goalUserID)

	if err != nil {
		return FAILED
	}
	return SUCCESS
}

// CancelConcern .
func (r *Redis) CancelConcern(currentUserID, goalUserID string) ProcessStatus {
	_, err := r.c.Do("SREM", "user_concern:"+currentUserID, goalUserID)

	if err != nil {
		return FAILED
	}
	return SUCCESS
}

// AddPraise .
func (r *Redis) AddPraise(messageID, userID string) ProcessStatus {
	result, err := redis.Int(r.c.Do("SADD", "praise_set:"+messageID, userID))

	if err != nil {
		logs.Warn("add praise FAILED!")
		return FAILED
	}

	if result == 0 {
		logs.Warn("add praise HAVEEXIST!")
		return HAVEEXIST
	}

	logs.Info("AddPraise Success!")
	return SUCCESS
}

// HavePraise .
func (r *Redis) HavePraise(messageID, userID string) bool {
	result, _ := redis.Int(r.c.Do("SISMEMBER", "praise_set:"+messageID, userID))
	return result == 1
}

// CancelPraise .
func (r *Redis) CancelPraise(messageID, userID string) ProcessStatus {
	_, err := redis.Int(r.c.Do("SREM", "praise_set:"+messageID, userID))

	if err != nil {
		logs.Warn("cancel praise faile!")
		return FAILED
	}

	logs.Info("CancelPraise Success!")
	return SUCCESS
}

// GetMessages .
func (r *Redis) GetMessages(userID string, start, end uint64) ([]string, error) {
	return redis.Strings(r.c.Do("LRANGE", "message_list:"+userID, start, end))
}

// GetMessage .
func (r *Redis) GetMessage(messageID string) (Message, bool) {
	var message Message
	messageData, err1 := redis.StringMap(r.c.Do("HGETALL", "message:"+messageID))
	if err1 != nil {
		return message, false
	}

	// 获取消息基本信息
	message.UserID = messageData["user_id"]
	message.Text = messageData["text"]
	message.Time = messageData["time"]

	// 获取消息基本信息
	// for key, value := range messageData {
	// 	switch key {
	// 	case "user_id":
	// 		message.UserID = value
	// 	case "text":
	// 		message.Text = value
	// 	case "time":
	// 		message.Time = value
	// 	}
	// }

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
	praiseNum, err3 := redis.Uint64(r.c.Do("SCARD", "praise_set:"+messageID))

	if err3 != nil {
		return message, false
	}

	message.Praise = praiseNum

	if commitSize > 0 {
		message.Commit = make([]CommitInfo, commitSize)
		for i, commit := range commits {
			json.Unmarshal([]byte(commit), &message.Commit[i])
			message.Commit[i].UserName, _ = redis.String(r.c.Do("HGET", "user:"+message.Commit[i].UserID, "name"))
			message.Commit[i].Image, _ = redis.String(r.c.Do("HGET", "user:"+message.Commit[i].UserID, "image"))
		}
	}

	logs.Info("find message success!")
	return message, true
}

// GetConcern .
func (r *Redis) GetConcern(userID string) ([]string, error) {
	return redis.Strings(r.c.Do("SMEMBERS", "user_concern:"+userID))
}

// GetConcernMessage .
func (r *Redis) GetConcernMessage(concerns []string, size uint64) []Message {
	var messages []Message
	var messageGroup []*list.List
	idMap := make(map[string]string)
	logs.Info("获取关注者的消息")
	// 获取关注者的消息
	for _, userID := range concerns {
		findMessages, ok := r.GetMessages(userID, 0, size)
		logs.Info("findMessages")
		if ok == nil {
			IDs := list.New()
			for _, messageID := range findMessages {
				id, _ := strconv.Atoi(messageID)
				idMap[messageID] = userID
				IDs.PushBack(id)
			}

			if IDs.Len() > 0 {
				messageGroup = append(messageGroup, IDs)
				// messageGroup.PushBack(IDs)
			}
		}
	}

	logs.Info("关注者的消息排序")
	// 关注者的消息排序
	if len(messageGroup) > 0 {
		logs.Info("MergeSorts")
		sortIDs := MergeSorts(messageGroup, size)
		index := uint64(0)
		e := sortIDs.Front()
		logs.Info("findMessage")
		for index < size && e != nil {
			findMessage, ok := r.GetMessage(strconv.Itoa(e.Value.(int)))

			if ok {
				findMessage.MessageID = strconv.Itoa(e.Value.(int))
				findMessage.UserID = idMap[strconv.Itoa(e.Value.(int))]
				findMessage.UserName, _ = redis.String(r.c.Do("HGET", "user:"+idMap[strconv.Itoa(e.Value.(int))], "name"))
				messages = append(messages, findMessage)
			}

			index++
			e = e.Next()
		}
	}

	return messages
}

// GetHotMessage .
func (r *Redis) GetHotMessage(userID string) []Message {
	return HotManager.GetHotMessage()
}

// MergeSorts .
func MergeSorts(messageGroup []*list.List, size uint64) *list.List {
	// for messageGroup.Len() > 1 {
	// 	list1 := messageGroup.Front()
	// 	messageGroup.Remove(list1)
	// 	list2 := messageGroup.Front()
	// 	messageGroup.Remove(list2)

	// 	messageGroup.PushBackList(MergeTwoSort(list1, list2))
	// }

	for len(messageGroup) > 1 {
		logs.Info(len(messageGroup))
		list1 := messageGroup[0]
		list2 := messageGroup[1]

		messageGroup = messageGroup[2:]

		messageGroup = append(messageGroup, MergeTwoSort(list1, list2))
	}

	return messageGroup[0]
}

// MergeTwoSort .
func MergeTwoSort(list1, list2 *list.List) *list.List {
	result := list.New()

	for {
		logs.Info(list1.Len(), list2.Len())
		if list1.Front().Value.(int) > list2.Front().Value.(int) {
			e := list1.Front()
			result.PushBack(e.Value.(int))
			list1.Remove(e)
			logs.Info("list1.Remove(e)")
		} else {
			e := list2.Front()
			result.PushBack(e.Value.(int))
			list2.Remove(e)
			logs.Info("list2.Remove(e)")
		}

		if list1.Len() == 0 {
			result.PushBackList(list2)
			break
		}

		if list2.Len() == 0 {
			result.PushBackList(list1)
			break
		}
	}

	return result
}

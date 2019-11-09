package models

// AddStatus .
type AddStatus int8

const (
	// AddSuccess 添加成功 .
	AddSuccess AddStatus = 0

	// AddFaile 添加失败 .
	AddFaile AddStatus = 1

	// HaveExist 已经存在 .
	HaveExist AddStatus = 2
)

// CommitInfo .
type CommitInfo struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Image    string `json:"image"`
	Commit   string `json:"commit"`
}

// Message .
type Message struct {
	UserID     string       `json:"user_id"`
	UserName   string       `json:"user_name"`
	MessageID  string       `json:"message_id"`
	Text       string       `json:"text"`
	Time       string       `json:"time"`
	Image      string       `json:"image"`
	Commit     []CommitInfo `json:"commits"`
	Praise     uint64       `json:"praise"`
	HavePraise bool         `json:"have_praise"`
}

// DataBase interface .
type DataBase interface {
	Init()                                                                // 初始化设置
	AuthenticateUser(userID, pwd string) (userName, head string, ok bool) // 用户验证
	AddUser(userID, name, pwd, image string) AddStatus                    // 添加用户
	AddMessage(userID, text, image string) AddStatus                      // 添加消息
	AddCommit(messageID, userID, commit string) AddStatus                 // 添加评论
	AddConcern(currentUserID, goalUserID string) AddStatus                // 添加关注
	AddPraise(messageID, userID string) AddStatus                         // 添加点赞
	CancelPraise(messageID, userID string) AddStatus                      // 取消点赞
	HavePraise(messageID, userID string) bool                             // 是否点赞
	GetMessages(userID string, start, end int) ([]string, error)          // 获取用户的消息列表
	GetMessage(messageID string) (Message, bool)                          // 获取消息
	GetConcern(userID string) []string                                    // 获取关注者
	GetHotMessage(userID string) []Message                               // 获取热点
}

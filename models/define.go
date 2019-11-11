package models

// ProcessStatus .
type ProcessStatus int8

const (
	// SUCCESS 添加成功 .
	SUCCESS ProcessStatus = 0

	// FAILED 添加失败 .
	FAILED ProcessStatus = 1

	// HAVEEXIST 已经存在 .
	HAVEEXIST ProcessStatus = 2
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
	AddUser(userID, name, pwd, image string) ProcessStatus                // 添加用户
	AddMessage(userID, text, image string) ProcessStatus                  // 添加消息
	AddCommit(messageID, userID, commit string) ProcessStatus             // 添加评论
	AddConcern(currentUserID, goalUserID string) ProcessStatus            // 添加关注
	CancelConcern(currentUserID, goalUserID string) ProcessStatus         // 取消关注
	AddPraise(messageID, userID string) ProcessStatus                     // 添加点赞
	CancelPraise(messageID, userID string) ProcessStatus                  // 取消点赞
	HavePraise(messageID, userID string) bool                             // 是否点赞
	GetMessages(userID string, start, end uint64) ([]string, error)       // 获取用户的消息列表
	GetMessage(messageID string) (Message, bool)                          // 获取消息
	GetConcern(userID string) ([]string, error)                                    // 获取关注者
	GetConcernMessage(concerns []string, size uint64) []Message           // 获取关注者的消息
	GetHotMessage(userID string) []Message                                // 获取热点
}

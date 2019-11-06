package models

import (
	"sync"
)

type model struct {
	db DataBase
}

// uniqueModel .
var uniqueModel *model = nil
var addPraiseLock sync.Mutex

// GetInstance .
func GetInstance() (modelInstance *model) {
	if uniqueModel == nil {
		db := Redis{}
		db.Init()
		// test
		// db.AddMessage("wxx", "Welcome to NewWorld.", "")
		uniqueModel = &model{&db}
	}
	return uniqueModel
}

// AuthenticateUser 验证用户 .
func (m *model) AuthenticateUser(name, pwd string) (ok bool) {
	return m.db.AuthenticateUser(name, pwd)
}

// AddUser .
func (m *model) AddUser(userID, name, pwd string) AddStatus {
	return m.db.AddUser(userID, name, pwd)
}

// AddMessage .
func (m *model) AddMessage(userID, text, image string) AddStatus {
	return m.db.AddMessage(userID, text, image)
}

// AddCommit .
func (m *model) AddCommit(messageID, userID, commit string) AddStatus {
	return m.db.AddCommit(messageID, userID, commit)
}

// AddConcern .
func (m *model) AddConcern(currentUserID, goalUserID string) AddStatus {
	return m.db.AddConcern(currentUserID, goalUserID)
}

// CancelPraise .
func (m *model) CancelPraise(messageID, userID string) AddStatus {
	return m.db.CancelPraise(messageID, userID)
}

// AddPraise .
func (m *model) AddPraise(messageID, userID string) AddStatus {
	return m.db.AddPraise(messageID, userID)
}

// GetMessage .
func (m *model) GetMessage(messageID string) (Message, bool) {
	return m.db.GetMessage(messageID)
}

// GetConcern .
func (m *model) GetConcern(userID string) []string {
	return m.db.GetConcern(userID)
}

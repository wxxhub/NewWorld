package models

type model struct {
	db DataBase
}

// uniqueModel .
var uniqueModel *model = nil

// GetInstance .
func GetInstance(flag int) (modelInstance *model) {
	if uniqueModel == nil {
		db := Redis{}
		db.Init()
		uniqueModel = &model{&db}
	}
	return uniqueModel
}

// AuthenticateUser 验证用户 .
func (m *model) AuthenticateUser(name, pwd string) (ok bool) {
	return m.db.AuthenticateUser(name, pwd)
}

//

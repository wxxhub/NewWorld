package models

type model struct {
	Flag int
}

// uniqueModel .
var uniqueModel *model = nil

// GetInstance .
func GetInstance(flag int) (modelInstance *model) {
	if uniqueModel == nil {
		uniqueModel = &model{flag}
	}
	return uniqueModel
}

// AuthenticateUser 验证用户 .
func (m *model) AuthenticateUser(name, pwd string) (ok bool) {
	return true
}

//

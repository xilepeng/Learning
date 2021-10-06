package Services

import "errors"

type IUserService interface {
	GetName(userid int) string
	DelUser(userid int) error
}

type UserService struct {}

func (this UserService) GetName(userid int) string {
	if userid == 522 {
		return "X"
	}
	return "guest"
}

func (this UserService) DelUser(userid int) error {
	if userid == 522 {
		return errors.New("无权限")
	}
	return nil
}

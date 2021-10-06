package Services

type IUserService interface {
	GetName(userid int) string
}

type UserService struct {}

func (this UserService) GetName(userid int) string {
	if userid == 522 {
		return "X"
	}
	return "guest"
}

package req

type CreateUserRequest struct {
	Username string `json:"username" example:"232323"` //账户
	Password string `json:"password" example:"232323"` //密码
}

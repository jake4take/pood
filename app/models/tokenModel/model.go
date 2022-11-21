package tokenModel

type Token struct {
	ID     uint   `json:"id"`
	Token  string `json:"token"`
	UserId uint   `json:"user_id"`
}

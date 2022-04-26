package account

type EncodedUser struct {
	Id       int    `json:"id"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`

	Token string `json:"token"`
}

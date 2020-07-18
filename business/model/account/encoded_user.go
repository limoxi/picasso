package account

type EncodedUser struct{
	Id int `json:"id"`
	Phone string `json:"phone"`
	CreatedAt string `json:"created_at"`

	Token string `json:"token"`
}
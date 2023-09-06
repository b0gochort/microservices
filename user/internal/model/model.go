package model

type UserCheckReq struct {
	UserID int `json:"user_id"`
}

type UserCheckRes struct {
	User User `json:"user"`
}

type User struct {
	UserID   int    `json:"id"`
	UserName string `json:"name"`
}

type HTTPResponse struct {
	Data any    `json:"data"`
	Err  string `json:"err"`
}

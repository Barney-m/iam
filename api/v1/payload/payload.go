package payload

import (
	"iam-server/token"
	"time"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	Payload *token.Payload `json:"payload"`
	Token   string         `json:"token"`
}

type RegisterReq struct {
	Email    string    `json:"email"`
	Password string    `json:"password"`
	FullName string    `json:"full_name"`
	MobileNo string    `json:"mobile_no"`
	Address  string    `json:"address"`
	Dob      time.Time `json:"dob"`
	Gender   string    `json:"gender"`
}

type RegisterRes struct {
	Payload *token.Payload `json:"payload"`
	Token   string         `json:"token"`
}

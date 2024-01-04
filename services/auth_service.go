package services

import (
	"iam-server/api/v1/payload"
	"iam-server/token"
)

type Authenticator interface {
	SignIn(email string, password string) (string, *token.Payload, error)
	SignUp(req *payload.LoginReq) (string, *token.Payload, error)
}

func SignIn() {
}

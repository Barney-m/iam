package models

type AuthUserAttr struct {
	UserId      string
	Salt        string
	Provider    string
	EncryptAlgo string
	HashAlgo    string
	BaseInfo
}

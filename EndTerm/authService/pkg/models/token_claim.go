package models

type TokenClaim struct {
	ID       int64
	Username string
	UserType string
	Exp      int64
	Iat      int64
}

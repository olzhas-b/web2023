package models

type Auth struct {
	ID     uint64
	Name   string
	Scopes []string
}

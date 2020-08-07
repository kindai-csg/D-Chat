package domain

type AuthType int

const (
	simple AuthType = iota
	ldap
)

type User struct {
	Id         string
	UserId     string
	Name       string
	Password   string
	Mail       string
	Bio        string
	Status     string
	StatusText string
	Auth       AuthType
}

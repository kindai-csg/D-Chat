package domain

type AuthType int

const (
	simple AuthType = iota
	ldap
)

type User struct {
	Id       string
	Name     string
	Password string
	Mail     string
	Profile  string
	Status   string
	Auth     AuthType
}

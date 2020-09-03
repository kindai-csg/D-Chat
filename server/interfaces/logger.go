package interfaces

type Logger interface {
	Write(interface{})
	Warnning(interface{})
	Error(interface{})
}

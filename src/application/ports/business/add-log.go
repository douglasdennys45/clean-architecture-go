package business

type AddLog interface {
	Add(logData []byte) error
}

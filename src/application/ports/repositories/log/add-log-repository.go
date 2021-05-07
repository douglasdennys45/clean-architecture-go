package log

type AddLogRepository interface {
	Add(logData []byte) error
}

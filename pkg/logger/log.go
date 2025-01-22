package logger

type ILogger interface {
	Info(msg string)
	InfoWithMeta(msg string, meta any)
}

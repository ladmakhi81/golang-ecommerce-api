package translations

type ITranslation interface {
	Message(key string) string
	MessageWithArgs(key string, data any) string
}

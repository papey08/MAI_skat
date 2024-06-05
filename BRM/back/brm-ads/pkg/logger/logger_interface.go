package logger

type Fields map[string]any

type Logger interface {
	Info(fields Fields, message string)
	Error(fields Fields, message string)
	Fatal(fields Fields, message string)
}

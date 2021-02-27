package logger

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type logger struct {
}

type Logger interface {
	logrus.FieldLogger
}

func SetContext(log Logger) context.Context {
	return context.WithValue(context.Background(), &logger{}, log.WithFields(map[string]interface{}{
		"time": time.Now(),
	}))
}

func NewLogger() Logger {
	log := logrus.StandardLogger()

	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	log = log.WithFields(map[string]interface{}{
		"time": time.Now(),
	}).Logger

	return log
}

func GetLogger(ctx context.Context) Logger {
	value := ctx.Value(&logger{})
	log, ok := value.(Logger)
	if !ok {
		log = NewLogger()
	}

	return log
}

func WithLogger(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, logger{}, log)
}
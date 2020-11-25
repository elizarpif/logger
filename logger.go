package logger

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type logger struct {
}

func SetContext(log *logrus.Logger) context.Context {
	return context.WithValue(context.Background(), &logger{}, log.WithFields(map[string]interface{}{
		"time": time.Now(),
	}))
}

func NewLogger() *logrus.Logger {
	log := logrus.StandardLogger()

	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	log = log.WithFields(map[string]interface{}{
		"time": time.Now(),
	}).Logger

	return log
}

func GetLogger(ctx context.Context) *logrus.Logger {
	value := ctx.Value(&logger{})
	log, ok := value.(*logrus.Logger)
	if !ok {
		log = NewLogger()
	}

	return log
}


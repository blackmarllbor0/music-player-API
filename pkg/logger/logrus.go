package logger

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	*logrus.Logger
}

func NewLogrusLogger() Logger {
	return &LogrusLogger{logrus.New()}
}

func (l LogrusLogger) SetupLogger(logDir string) (*os.File, error) {
	l.SetLevel(logrus.InfoLevel)

	l.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		FullTimestamp:    true,
		DisableColors:    false,
		DisableTimestamp: false,
	})

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0775); err != nil {
			return nil, err
		}
	}

	logFileName := filepath.Join(logDir, time.Now().Format("2006-01-02")+".log")

	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	multiWriter := io.MultiWriter(file, os.Stdout)

	l.SetOutput(multiWriter)

	return file, nil
}

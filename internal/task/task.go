package task

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Task struct {
	id     string
	logger *logrus.Logger
	fn     func(logger *logrus.Logger)
}

func NewTask(logger *logrus.Logger) *Task {
	id := uuid.New().String()
	task := &Task{
		id: id,
		logger: logger.WithFields(logrus.Fields{
			"taskId": id,
		}).Logger,
	}

	return task
}

func (t *Task) Run() {
	go t.fn(t.logger)
}

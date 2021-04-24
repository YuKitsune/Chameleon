package task

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Task struct {
	id string
	logger *logrus.Entry
	fn func(logger *logrus.Entry)
}

func NewTask(logger *logrus.Logger) *Task {
	id := uuid.New().String()
	task := &Task{
		id: id,
		logger: logger.WithFields(logrus.Fields{
			"taskId": id,
		}),
	}

	return task
}

func (t *Task) Run() {
	go t.fn(t.logger)
}

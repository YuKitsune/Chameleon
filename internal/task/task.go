package task

import (
	"github.com/google/uuid"
	"github.com/yukitsune/chameleon/internal/log"
)

type Task struct {
	id string
	logger log.ChameleonLogger
	fn func(logger log.ChameleonLogger)
}

func NewTask(logger log.ChameleonLogger) *Task {
	id := uuid.New().String()
	task := &Task{
		id: id,
		logger: logger.WithFields(log.Fields{
			"taskId": id,
		}),
	}

	return task
}

func (t *Task) Run() {
	go t.fn(t.logger)
}

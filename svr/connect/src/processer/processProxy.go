package processer

import (
	"connect/src/common/session"
	"connect/src/processer/internal"
)

type processerManager struct {
}

var instance *processerManager

func init() {
	if instance == nil {
		instance = new(processerManager)
	}
}

func Instance() *processerManager {
	return instance
}

func (receiver processerManager) Process(psession *session.Session) {
	internal.Do(psession)
}

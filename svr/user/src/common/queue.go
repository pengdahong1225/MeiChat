package common

import (
	"errors"
)

// 队列
type Queue[T any] struct {
	data   []T
	length int
}

func (receiver *Queue[T]) Push(val T) {
	receiver.data = append(receiver.data, val)
	receiver.length++
}

func (receiver *Queue[T]) Pop(val *T) error {
	if receiver.length == 0 {
		return errors.New("queue is empty")
	}
	*val = receiver.data[0]
	receiver.data = receiver.data[1:]
	receiver.length--
	return nil
}

func (receiver *Queue[T]) GetLength() int {
	return receiver.length
}

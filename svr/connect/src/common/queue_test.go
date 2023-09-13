package common

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	var que_ *Queue[int]
	for i := 0; i < 10; i++ {
		que_.Push(i)
	}
	fmt.Println(*que_)
	length := que_.length
	for i := 0; i < length; i++ {
		var val int
		err := que_.Pop(&val)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(val)
		}
	}
}

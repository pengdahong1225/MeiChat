package common

import (
	"github.com/panjf2000/ants/v2"
	"log"
	"runtime"
)

var AntsPoolInstance *ants.Pool

func init() {
	AntsPoolInstance, _ = ants.NewPool(ants.DefaultAntsPoolSize, ants.WithPanicHandler(AntsPanicHandler))
}

func AntsPanicHandler(i interface{}) {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	log.Printf("worker exits from panic: %s\n", string(buf[:n]))
}

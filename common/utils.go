package common

import (
	"fmt"
	"runtime"
	"time"
)

func TimeConsume(start time.Time) {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	fmt.Printf("func_time_consume:%s cost: %s\n", funcName, time.Since(start))
}

package common


import (
	"fmt"
	"runtime"

	log "github.com/kataras/golog"
)

const (
	maxStack  = 20
	separator = "---------------------------------------\n"
)

// HandlePanic handle panic
func HandlePanic() {
	if err := recover(); err != nil {
		errstr := fmt.Sprintf("%sruntime error: %v\ntraceback:\n", separator, err)

		i := 2
		for {
			pc, file, line, ok := runtime.Caller(i)
			if !ok || i > maxStack {
				break
			}
			errstr += fmt.Sprintf("    stack: %d %v [file: %s] [func: %s] [line: %d]\n", i-1, ok, file, runtime.FuncForPC(pc).Name(), line)
			i++
		}
		errstr += separator

		log.Error(errstr)
	}
}

// Safe safe func call auto handle panic
func Safe(cb func()) {
	defer HandlePanic()
	cb()
}

// Go safe go routine auto handle panic
func Go(cb func()) {
	go Safe(cb)
}
package recorder

import "sync"

type Recorder interface {
	Post(group string, key string, uuid string)
}

var (
	mu sync.Mutex
	r  Recorder
)

func Get() Recorder {
	mu.Lock()
	defer mu.Unlock()

	return r
}

func SetGlobal(recorder Recorder) {
	r = recorder
}

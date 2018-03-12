package stdout

import "fmt"

type StdoutRecorder struct{}

func (r *StdoutRecorder) Post(group string, key string, uuid string) {
	fmt.Printf(
		"Group: %s, Key: %s, UUID: %s\n",
		group,
		key,
		uuid,
	)
}

func NewStdoutRecorder() *StdoutRecorder {
	return &StdoutRecorder{}
}

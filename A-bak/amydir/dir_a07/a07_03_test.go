package dira07

import (
	"log"
	"testing"
)

/*
// 编译错误
type Logger struct {
	Level int
}

type MyJob struct {
	*Logger
	Name        string
	*log.Logger // duplicate field Logger
}
*/

type LoggerM struct {
	Level int
}

type MyJob struct {
	*LoggerM
	Name        string
	*log.Logger // duplicate field Logger
}

func TestZxx(t *testing.T) {
}

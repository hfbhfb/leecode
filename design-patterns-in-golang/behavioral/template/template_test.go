package template

import "testing"

//08 P8 模板模式
func TestWorker_Daily(t *testing.T) {
	worker :=NewWorker(&Coder{})

	worker.Daily()
}
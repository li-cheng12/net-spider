package schedule

import "GoDemo/crawler/engine"

type SimpleSchedule struct {
	WorkQueue chan engine.Request
}

func (s *SimpleSchedule)Submit(request engine.Request) {
	go func () {
		s.WorkQueue <- request
		return
	}()
}

func (s* SimpleSchedule)WorkReady(chan engine.Request) {

}
func (s* SimpleSchedule)WorkSchedule() chan engine.Request {
	return s.WorkQueue
}

func (s* SimpleSchedule)Run() {
	s.WorkQueue = make(chan engine.Request)
}

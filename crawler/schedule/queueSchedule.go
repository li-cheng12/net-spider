package schedule

import "GoDemo/crawler/engine"

type QueueSchedule struct {
	workQueue chan chan engine.Request
	requestQueue chan engine.Request
}
func (q* QueueSchedule) Submit(r engine.Request) {
	q.requestQueue <- r
}
func (q* QueueSchedule) WorkReady(rq chan engine.Request) {
	q.workQueue <- rq
}
func (q* QueueSchedule) WorkSchedule() chan engine.Request {
	return make(chan engine.Request)
}

func (q* QueueSchedule) Run() {
	q.workQueue = make(chan chan engine.Request)
	q.requestQueue = make(chan engine.Request)

	go func() {
		activeRequests := []engine.Request{}
		activeWorkers := []chan engine.Request{}
		for {
			var activeRequest engine.Request
			var activeWork chan engine.Request
			if len(activeRequests) > 0 && len(activeWorkers) > 0 {
				activeWork = activeWorkers[0]
				activeRequest = activeRequests[0]
			}
			select {
			case r := <- q.workQueue:
				activeWorkers = append(activeWorkers, r)
				case r := <- q.requestQueue:
					activeRequests = append(activeRequests, r)
					case activeWork <- activeRequest:
						activeRequests = activeRequests[1:]
						activeWorkers = activeWorkers[1:]
			}
		}
	}()
}

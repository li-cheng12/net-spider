package engine

import (
	"github.com/labstack/gommon/log"
	"time"
	"GoDemo/crawler/model"
)

type ConcurrentEngin struct {
	Schedule ISchedule
	WorkCount int
	Save ISave
}
type ISchedule interface {
	Submit(request Request)
	WorkReady(chan Request)
	WorkSchedule() chan Request
	Run()
}

type ISave interface {
	CreateSaver(item interface{})
	Run()
}
var count = 0
func (c *ConcurrentEngin) Run(seeds ...Request) {
	out := make(chan ParseResult)
	c.Schedule.Run()
	c.Save.Run()
	for i := 0;i < c.WorkCount;i++ {
		in := c.Schedule.WorkSchedule()
		createWorker(in, c.Schedule, out)
	}
	for _, r := range seeds {
		c.Schedule.Submit(r)
	}
	for v := range out {
		for _, v := range v.Items {
			profile,ok := v.(model.Profile)
			if ok {
				log.Print(profile)
			}
			log.Printf("Got item %d---%v", count, v)
			count++
		}
		for _, r := range v.Requests {
			c.Schedule.Submit(r)
		}
	}
}
var limiter = time.Tick(10*time.Millisecond)
func createWorker(in chan Request, s ISchedule, out chan ParseResult) {
	go func() {
		for {
			s.WorkReady(in)
			request := <- in
			log.Printf("Fetching %s", request.Url)
			<-limiter
			parserResult,err := worker(request)
			if err != nil {
				continue
			}
			out <- parserResult
		}
	}()

}

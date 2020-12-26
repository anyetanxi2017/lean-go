package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

var entryId cron.EntryID = 0
var c = cron.New()

type MyJob struct {
	Interval int
}

func (j MyJob) GetInterval() string {
	return fmt.Sprintf("@every %ds", j.Interval)
}

func (j MyJob) Run() {
	j.SetNextTime(5)
}

// 设置下次执行时间
func (j MyJob) SetNextTime(interval int) {
	c.Remove(entryId)
	myJob := MyJob{interval}
	entryId, _ = c.AddJob(myJob.GetInterval(), myJob)
}

func Start() {
	MyJob{}.SetNextTime(10)
	c.Start()
}

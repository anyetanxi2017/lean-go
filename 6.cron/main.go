package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

var entryId cron.EntryID = 0
var num = 0
var c = cron.New()

type MyJob struct {
	Interval int
}

func (j MyJob) GetInterval() string {
	return fmt.Sprintf("@every %ds", j.Interval)
}

func (j MyJob) Run() {
	num++
	fmt.Println(time.Now())
	if num%2 == 0 {
		c.Remove(entryId)
		fmt.Println("快 修改时间")
		myJob := MyJob{5}
		entryId, _ = c.AddJob(myJob.GetInterval(), myJob)
	}
}

func Start() {
	myJob := MyJob{10}
	entryId, _ = c.AddJob(myJob.GetInterval(), myJob)
	c.Start()
}

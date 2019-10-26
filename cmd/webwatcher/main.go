package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webwatcher/internal/jobs"

	"github.com/robfig/cron/v3"
)

func main() {
	fmt.Println("Program Start.")

	//schedule jobs
	schedule := cron.New(cron.WithLocation(time.FixedZone("UTC+8", 8*60*60)))
	schedule.AddFunc("0 9 * * *", jobs.CheckNaebaBusBookingDates) //run at 10:00 daily
	schedule.Start()

	waitForInterrupt()

	//wait for jobs completed
	<-schedule.Stop().Done()

	fmt.Println("Program Exit.")
}

func waitForInterrupt() {
	systemSignalChan := make(chan os.Signal)

	signal.Notify(systemSignalChan)

	for {
		if s := <-systemSignalChan; s == syscall.SIGINT || s == syscall.SIGTERM {
			break
		}
	}
}

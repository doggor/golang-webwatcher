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
	//log program life status
	fmt.Println("Chick wakeup")
	defer fmt.Println("Chick sleep")

	//schedule jobs
	schedule := cron.New(cron.WithLocation(time.FixedZone("UTC+8", 8*60*60)))
	entryID, err := schedule.AddFunc("0 10 * * *", jobs.CheckNaebaBusBookingDates) //run at 10:00 daily

	//start schedule
	schedule.Start()
	defer func() {
		//wait for jobs completed
		<-schedule.Stop().Done()
	}()

	//log job status
	if err != nil {
		fmt.Println("Fail to schedule Naeba bus booking checker:", err)
	} else {
		fmt.Println("The next time checking Naeba bus booking is", schedule.Entry(entryID).Next)
	}

	//wait for system interrupt
	systemSignalChan := make(chan os.Signal)
	signal.Notify(systemSignalChan)
	for {
		s := <-systemSignalChan
		if s == syscall.SIGINT || s == syscall.SIGTERM {
			break
		}
	}
}

package main

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/xian1367/layout-go-zero/config"
	"github.com/xian1367/layout-go-zero/pkg/console"
	"github.com/xian1367/layout-go-zero/pkg/orm"
	"github.com/zeromicro/go-zero/core/logc"
	"time"
)

func main() {
	config.Init("cron")
	logc.MustSetup(config.Get().Cron.Log)
	orm.Init()

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		// handle error
		console.ExitIf(err)
	}

	// add a job to the scheduler
	_, err = scheduler.NewJob(
		gocron.DurationJob(
			time.Second,
		),
		gocron.NewTask(
			func() {

			},
		),
	)
	if err != nil {
		// handle error
		console.ExitIf(err)
	}

	// start the scheduler
	scheduler.Start()

	// block until you are ready to shut down
	select {
	case <-time.After(time.Minute):
	}

	// when you're done, shut it down
	err = scheduler.Shutdown()
	if err != nil {
		// handle error
	}
}

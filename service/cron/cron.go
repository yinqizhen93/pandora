package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	parser := cron.NewParser(
		cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)
	c := cron.New(cron.WithParser(parser))

	c.AddFunc("*/7 * * * * *", func() {
		fmt.Println(time.Now())
	})

	c.Start()
	select {}
	//time.Sleep(time.Second * 5)
}

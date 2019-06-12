package main

import (
	"agent/service"
	"agent/utill"
	"github.com/robfig/cron"
	"log"
	"os"
)

func main() {

	schedule, commandType, command := utill.ParseLaunchArgs()

	cronJob := cron.New()
	cronJob.AddFunc(schedule, func() {

		defer func() {
			if err := recover(); err != nil {
				log.Print(err)
				os.Exit(-1)
			}
		}()

		service.ExecCommand(commandType, command)

	})

	cronJob.Start()

	select {}
}

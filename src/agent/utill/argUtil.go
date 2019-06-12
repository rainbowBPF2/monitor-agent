package utill

import (
	"agent/conf"
	"fmt"
	"log"
	"os"
)

var defaultSchedule = "*/5 * * * *"
var defaultType = "FILE"
var defaultCommand = "startAgent.sh"

func Help() {
	fmt.Println("************ Ops Monitor Agent ************\n")
	fmt.Printf("\t%s\t%s\n", "-s", " cronjob schedule. Such as -s=*/5 * * * * , execute for each 5 seconds.")
	fmt.Printf("\t%s\t%s\n", "-t", " cronjob command type. Such as FILE, BASH.")
	fmt.Printf("\t%s\t%s\n", "-c", " command to be execute, such as date, ls, mv, test.sh.")
	fmt.Printf("\t%s\t%s\n", "-h", " ops agent, help tips.")
}

func ParseLaunchArgs() (schedule string, commandType string, command string) {

	args := os.Args

	if len(args) > 1 && (args[1] == "-h" || args[1] == "--help") {
		Help()
		return
	}

	log.Print("****** Launch ops-monitor agent ******")

	configMap := ParseAppConf()

	if configMap[conf.CronSchedule] != "" {
		defaultSchedule = configMap[conf.CronSchedule]
	}

	if configMap[conf.CommandType] != "" {
		defaultType = configMap[conf.CommandType]
	}

	if configMap[conf.Command] != "" {
		defaultCommand = configMap[conf.Command]
	}

	return defaultSchedule, defaultType, defaultCommand
}

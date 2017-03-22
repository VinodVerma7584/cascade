package command

import (
	"fmt"
	"github.com/jwaldrip/odin/cli"
	"log"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/watch"
	"errors"
)

/**
Some basic wait functionality built on top of consul.

Lots of potential improvements could be made here:

  - pass in a specific tag
  - implement a timeout (and add as optional arg)
  - implement a quorum, eg wait until n nodes are healthy, along with status output to support it
  -
 */
var Wait = cli.NewSubCommand("wait", "Wait operations", waitRun)

func init() {

	Wait.SetLongDescription(`
Wait for cascade services

Usage:
	cascade wait <service> [services...]

  `)
}

func waitRun(c cli.Command) {
	services := c.Args().Strings()
	if len(services) == 0 {
		cli.ShowUsage(c)
		log.Fatalln("err: missing <service> argument")
	}
	doWait(services)
}

func doWait(services [] string) {
	for _, service := range services {
		fmt.Println("Waiting for " + service + " to report as healthy")
		err := waitFor(service)
		if (err != nil) {
			log.Fatalln("failed waiting for " + service, err)
		}
	}
}

func waitFor(serviceName string) error {

	// Setup watch
	watchParams := make(map[string]interface{})

	watchParams["type"] = "checks"
	watchParams["service"] = serviceName

	theWatch, err := watch.Parse(watchParams)
	if err != nil {
		return err
	}

	theWatch.Handler = func(idx uint64, data interface{}) {
		checks := data.([]*api.HealthCheck)

		passedCount := 0
		for _, check := range checks {
			if check.Status == api.HealthPassing {
				fmt.Println("passing: " + check.CheckID)
				passedCount ++
			} else {
				fmt.Println("not passing: " + check.CheckID)
			}
		}

		totalCount := len(checks)
		if totalCount > 0 && passedCount == totalCount {
			fmt.Println("all passing")
			theWatch.Stop()
		}
	}

	// Execute Watch
	if err := theWatch.Run("127.0.0.1:8500"); err != nil {
		return errors.New(fmt.Sprintf("err: querying Consul agent: %s", err))
	}
	return nil
}
package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/spark"
)

func main() {
	gbot := gobot.NewGobot()

	sparkCore := spark.NewSparkCoreAdaptor("spark", "DEVICE_ID", "ACCESS_TOKEN")

	work := func() {
		temp, err := sparkCore.Variable("temperature")

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("temp from variable is: %v", temp)
		}
	}

	robot := gobot.NewRobot("spark",
		[]gobot.Connection{sparkCore},
		[]gobot.Device{},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}

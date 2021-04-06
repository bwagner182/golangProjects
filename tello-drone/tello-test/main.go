package main

import (
    "time"

    "gobot.io/x/gobot"
    "gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
    drone := tello.NewDriver("8888")

    work := func() {
        drone.TakeOff()

        gobot.After(3*time.Second, func() {
        	drone.Down(20)
        })

        gobot.After(5*time.Second, func() {
        	drone.Land()
        })

        gobot.After(1*time.Second, func() {
        	drone.Halt()
        })
    }

    robot := gobot.NewRobot("tello",
        []gobot.Connection{},
        []gobot.Device{drone},
        work,
    )

    robot.Start()
}

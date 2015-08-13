package main

import (
	"fmt"
//	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/ChimeraCoder/anaconda"
	"strings"
	"strconv"
)

func main() {
	gbot := gobot.NewGobot()

	firmataAdaptor := firmata.NewFirmataAdaptor("firmata", "COM3")
	servo := gpio.NewServoDriver(firmataAdaptor, "servo", "9")
	fmt.Println("Hello Twitter, 世界")
	anaconda.SetConsumerKey("bsYrbbtqg9l2u2hYhgz76h6ux")
	anaconda.SetConsumerSecret("urcJNRZ9qZeCUQfqEoZ0e4Q5JMUmbDgn5bRtbWN0lTyhkgvkPx")
	api := anaconda.NewTwitterApi("3145150396-22DQi9gphmz3XX30Y6jCAmKrWvzjaldIQFybqry", "oEOcLxm8uo8V7jSm3IKNq44jcgS66zhikIbgEg4r0l4UT")

	stream := api.UserStream(nil)

	work := func() {
		for {
			channel := <-stream.C
			t, ok := channel.(anaconda.Tweet)
			if !ok {
				fmt.Println("Recieved non tweet message")
			}
			if(strings.Contains(t.Text,"#turn")){
				degree := strings.Split(t.Text,"#turn")[1]
				x,_ := strconv.ParseUint(degree,10,8)
				fmt.Println("Turning ",x)
				servo.Move(uint8(x))
			}
		}
	}

	robot := gobot.NewRobot("servoBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{servo},
		work,
	)

	gbot.AddRobot(robot)
	gbot.Start()
}
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	redHeadDuck := NewReadHeadDuck(nil, nil)
	redHeadDuck.Display()

	flyWithNothing := NewFlyWithNothing()
	redHeadDuck.SetFlyBehavior(flyWithNothing)
	redHeadDuck.Display()

	rubbleDuck := NewRubbleDuck(nil, nil)
	rubbleDuck.Display()

	flyWithWings := NewFlyWithWings()
	rubbleDuck.SetFlyBehavior(flyWithWings)
	rubbleDuck.Display()

}

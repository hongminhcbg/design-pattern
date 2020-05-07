package main
import (
	"fmt"
)

type command interface {
	Execute()
}

func main() {
	fmt.Println("hello world")
	remote := remote{}
	light := NewLight()

	lightOnCommand := NewLightOnCommand(light)
	lightOffCommand := NewLightOffCommand(light)

	remote.SetCommand(lightOnCommand)
	remote.ButtonWasPressed()

	remote.SetCommand(lightOffCommand)
	remote.ButtonWasPressed()

	remote.SetCommand(lightOnCommand)
	remote.ButtonWasPressed()
}

package main

import "fmt"

func NewLight() Light {
	return &lightImpl{}
}

type lightOnCommandImpl struct {
	light Light
}
func (cmd *lightOnCommandImpl) Execute() {
	cmd.light.On()
}
func NewLightOnCommand(light Light) command {
	return &lightOnCommandImpl{light:light}
}


func NewLightOffCommand(light Light) command {
	return &lightOffCommandImpl{light:light}
}
type lightOffCommandImpl struct {
	light Light
}
func (cmd *lightOffCommandImpl) Execute() {
	cmd.light.Off()
}

type Light interface {
	On()
	Off()
	Break()
}

type lightImpl struct {
}

func (l *lightImpl) On()  {
	fmt.Println("turn on light")
}
func (l *lightImpl) Off()  {
	fmt.Println("turn off light")
}
func (l *lightImpl) Break()  {
	fmt.Println("light is broken")
}

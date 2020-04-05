package main

import "fmt"

type NoFly struct {

}

func (client *NoFly)Fly()  {
	fmt.Println("I can't fly, ahihihi")
}

func NewFlyWithNothing() FlyBehavior {
	return &NoFly{}
}
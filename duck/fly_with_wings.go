package main

import "fmt"

type FlyWithWings struct {

}

func (client *FlyWithWings)Fly()  {
	fmt.Println("I'm fly with wings")
}

func NewFlyWithWings() FlyBehavior {
	return &FlyWithWings{}
}
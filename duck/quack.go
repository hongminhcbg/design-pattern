package main

import "fmt"

type quack struct {

}

func (client *quack) Quack() {
	fmt.Println("quack quack !!!!")
}

func NewQuackSound() QuackBehavior {
	return &quack{}
}
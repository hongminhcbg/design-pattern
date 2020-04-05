package main

import "fmt"

type quackNothing struct {

}

func (quack *quackNothing)Quack()  {
	fmt.Println("I can't make quack quack sound, sorry")
}

func NewQuackNothing() QuackBehavior {
	return &quackNothing{}
}
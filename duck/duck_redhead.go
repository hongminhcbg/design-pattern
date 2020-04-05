package main

type duckRedHeadImpl struct {
	flyBehavior FlyBehavior
	quackBehavior QuackBehavior
}

func (duck *duckRedHeadImpl)PerformFly()  {
	duck.flyBehavior.Fly()
}

func (duck *duckRedHeadImpl)PerformQuack()  {
	duck.quackBehavior.Quack()
}

func (duck *duckRedHeadImpl)SetFlyBehavior(fly FlyBehavior)  {
	duck.flyBehavior = fly
}

func (duck *duckRedHeadImpl)SetQuackBehavior(quack QuackBehavior)  {
	duck.quackBehavior = quack
}

func (duck *duckRedHeadImpl) Display() {
	duck.quackBehavior.Quack()
	duck.flyBehavior.Fly()
}

func NewReadHeadDuck(fly FlyBehavior, quack QuackBehavior) Duck {
	if fly == nil {
		fly = NewFlyWithWings()
	}

	if quack == nil {
		quack = NewQuackSound()
	}

	return &duckRedHeadImpl{
		flyBehavior:   fly,
		quackBehavior: quack,
	}
}
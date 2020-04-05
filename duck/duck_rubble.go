package main

type duckRubbleImpl struct {
	flyBehavior FlyBehavior
	quackBehavior QuackBehavior
}

func (duck *duckRubbleImpl)PerformFly()  {
	duck.flyBehavior.Fly()
}

func (duck *duckRubbleImpl)PerformQuack()  {
	duck.quackBehavior.Quack()
}

func (duck *duckRubbleImpl)SetFlyBehavior(fly FlyBehavior)  {
	duck.flyBehavior = fly
}

func (duck *duckRubbleImpl)SetQuackBehavior(quack QuackBehavior)  {
	duck.quackBehavior = quack
}

func (duck *duckRubbleImpl) Display() {
	duck.quackBehavior.Quack()
	duck.flyBehavior.Fly()
}

func NewRubbleDuck(fly FlyBehavior, quack QuackBehavior) Duck {
	if fly == nil {
		fly = NewFlyWithNothing()
	}

	if quack == nil {
		quack = NewQuackNothing()
	}

	return &duckRubbleImpl{
		flyBehavior:   fly,
		quackBehavior: quack,
	}
}

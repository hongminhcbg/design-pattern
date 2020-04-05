package main
type Duck interface {
	PerformFly()
	PerformQuack()
	SetFlyBehavior(fly FlyBehavior)
	SetQuackBehavior(quack QuackBehavior)
	Display()
}

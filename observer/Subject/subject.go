package Subject

import "github.com/hongminhcbg/design-pattern/observer/observer"

type Subject interface {
	Subscribe(observer observer.Observer)
	Remove(observer observer.Observer)
	Notification()
}

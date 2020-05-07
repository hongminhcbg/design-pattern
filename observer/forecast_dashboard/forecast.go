package forecast_dashboard

import (
	"fmt"
	"github.com/hongminhcbg/design-pattern/observer/Subject"
	"github.com/hongminhcbg/design-pattern/observer/observer"
)

type ForecastDashBoard struct {
	subject Subject.Subject
	light   float64
	temp    float64
	hum     float64
}

func NewForecastDashBoard(subject Subject.Subject) observer.Observer {
	result := ForecastDashBoard{
		subject: subject,
		light:   0,
		temp:    0,
		hum:     0,
	}

	subject.Subscribe(&result)
	return &result
}

func (f *ForecastDashBoard) Update(light, temp, hum float64) {
	f.light = light
	f.temp = temp
	f.hum = hum
	f.Display()
}

func (f *ForecastDashBoard) Display() {
	fmt.Printf("light = %f, temp = %f, hum = %f\n", f.light, f.temp, f.hum)
}

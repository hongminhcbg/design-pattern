package weather_station

import "github.com/hongminhcbg/design-pattern/observer/observer"

type WeatherStation struct {
	observers []observer.Observer
	light     float64
	temp      float64
	hum       float64
}

func NewWeatherStation() WeatherStation {
	result := WeatherStation{}
	result.observers = make([]observer.Observer, 0)
	result.light = 0
	result.temp = 0
	result.hum = 0
	return result
}

func (w *WeatherStation) Subscribe(observer observer.Observer) {
	w.observers = append(w.observers, observer)
}

func (w *WeatherStation) Remove(observer observer.Observer) {
	for i := 0; i < len(w.observers); i++ {
		if w.observers[i] == observer {
			w.observers[i] = nil
			return
		}
	}
}

func (w *WeatherStation) Notification() {
	for _, val := range w.observers {
		if val == nil {
			continue
		}
		val.Update(w.light, w.temp, w.hum)
	}
}

func (w *WeatherStation) SetData(light, temp, hum float64) {
	w.light = light
	w.temp = temp
	w.hum = hum
	w.Notification()
	return
}

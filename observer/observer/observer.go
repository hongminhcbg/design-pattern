package observer

type Observer interface {
	Update(light, temp, hum float64)
	Display()
}

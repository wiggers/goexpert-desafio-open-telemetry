package entity

type Weather struct {
	Temp_C float32
}

func (w *Weather) GetFahrenheit() float32 {
	return w.Temp_C*1.8 + 32
}

func (w *Weather) GetKelvin() float32 {
	return w.Temp_C + 273
}

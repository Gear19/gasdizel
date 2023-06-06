package gasdizel

type Liters float64
type Gallons float64

func (m Liters) ToGallons(s Gallons) Liters {
	return Liters(s)
}

func (m Gallons) ToLiters(s Liters) Gallons {
	return Gallons(s)
}

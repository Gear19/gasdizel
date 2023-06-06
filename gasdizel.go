package gasdizel

type Liters float64
type Gallons float64

func (m Liters) ToGallons(s Liters) Gallons {
	return Gallons(s * 0.264)
}

func (m Gallons) ToLiters(s Gallons) Liters {
	return Liters(s * 3.785)
}

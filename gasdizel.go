package gasdizel

type Liters float64
type Gallons float64

func (m Liters) ToGallons(s Gallons) Liters {
	return Liters(s * 3.785)
}

func (m Gallons) ToLiters(s Liters) Gallons {
	return Gallons(s * 0.264)
}

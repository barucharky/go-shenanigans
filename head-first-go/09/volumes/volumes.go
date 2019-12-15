// B"H

package volumes

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (l Liters) DoubleWrong() {
	l *= 2
}

func (l *Liters) Double() {
	*l *= 2
}

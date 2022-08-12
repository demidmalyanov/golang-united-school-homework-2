package square

import (
	"math"
)
// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type sideQuantity int

const (

	SidesTriangle sideQuantity = 3
	SidesSquare sideQuantity = 4
	SidesCircle sideQuantity = 0
)



func CalcSquare(sideLen float64, sidesNum sideQuantity) float64 {
	switch sidesNum {
	case SidesTriangle:
		square := math.Sqrt(3)/4*math.Pow(sideLen, 2)
		return square

	case SidesSquare:
		square := math.Pow(sideLen,2)
		return square

	case SidesCircle:
		square := math.Pi * math.Pow(sideLen,2)
		return square
	default:
		return 0
	}
}


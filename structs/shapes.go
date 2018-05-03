package shapes

type Rectangle struct {
	Wight float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Wight + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Wight * rectangle.Height
}

func main() {
	return
}
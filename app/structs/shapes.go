package main

// Rectangle has the dimentions of rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimerter returns the perimeter of a rectangle.
func Perimerter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area returns the area of a rectangle.
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

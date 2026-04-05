package perimeter

// function for count perimeter of rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// function for count perimeter of rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

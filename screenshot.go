package screenshot

import (
	"image"
)

// CaptureDisplay captures whole region of displayIndex'th display.
func CaptureDisplay(displayIndex int) (*image.RGBA, error) {
	rect := GetDisplayBounds(displayIndex)
	return CaptureRect(rect)
}

// CaptureRect captures specified region of desktop.
func CaptureRect(rect image.Rectangle) (*image.RGBA, error) {
	return Capture(rect.Min.X, rect.Min.Y, rect.Dx(), rect.Dy())
}
// CaptureDisplay captures whole region of displayIndex'th display.
func CaptureDisplayX2(displayIndex int) (*image.RGBA, error) {
	rect := GetDisplayBounds(displayIndex)
	return CaptureRectX2(rect)
}

// CaptureRect captures specified region of desktop.
func CaptureRectX2(rect image.Rectangle) (*image.RGBA, error) {
	return Capture(rect.Min.X, rect.Min.Y, rect.Dx()*2, rect.Dy()*2)
}

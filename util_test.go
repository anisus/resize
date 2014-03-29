package resize

import (
	"image"
	"testing"
)

func Test_FromWidth(t *testing.T) {
	rect := image.Rect(0, 0, 100, 200)
	w, h := FromWidth(rect, 100)

	if w != 100 || h != 200 {
		t.Fail()
	}
}

func Test_FromHeight(t *testing.T) {
	rect := image.Rect(0, 0, 100, 200)
	w, h := FromHeight(rect, 200)

	if w != 100 || h != 200 {
		t.Fail()
	}
}

func Test_FromWidthLargeInt(t *testing.T) {
	rect := image.Rect(0, 0, 500000, 500000)
	w, h := FromWidth(rect, 10000)

	if w != 10000 || h != 10000 {
		t.Fail()
	}
}

func Test_Thumbnail(t *testing.T) {
	var thumbnailTests = []struct {
		origWidth      int
		origHeight     int
		maxWidth       uint
		maxHeight      uint
		expectedWidth  uint
		expectedHeight uint
	}{
		{5, 5, 10, 10, 5, 5},
		{10, 10, 5, 5, 5, 5},
		{10, 50, 10, 10, 2, 10},
		{50, 10, 10, 10, 10, 2},
		{50, 100, 60, 90, 45, 90},
		{120, 100, 60, 90, 60, 50},
		{200, 250, 200, 150, 120, 150},
	}

	for _, tt := range thumbnailTests {
		rect := image.Rect(0, 0, tt.origWidth, tt.origHeight)
		actualWidth, actualHeight := Thumbnail(rect, tt.maxWidth, tt.maxHeight)
		if actualWidth != tt.expectedWidth || actualHeight != tt.expectedHeight {
			t.Fail()
		}
	}
}

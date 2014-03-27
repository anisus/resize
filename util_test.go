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

/*
Copyright (c) 2014, Jan Schlicht <jan.schlicht@gmail.com>

Permission to use, copy, modify, and/or distribute this software for any purpose
with or without fee is hereby granted, provided that the above copyright notice
and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS
OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER
TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF
THIS SOFTWARE.
*/

package resize

import (
	"image"
)

// Utility function that calculates the height needed to resize an image
// to a certain width, preserving its aspect ratio.
// Returns width and height to be easily interchangeable with the other
// utility functions.
func FromWidth(rect image.Rectangle, width uint) (uint, uint) {
	aspectInv := float32(rect.Dy()) / float32(rect.Dx())
	return width, uint(float32(width) * aspectInv)
}

// Utility function that calculates the width needed to resize an image
// to a certain height, preserving its aspect ratio.
// Returns width and height to be easily interchangeable with the other
// utility functions.
func FromHeight(rect image.Rectangle, height uint) (uint, uint) {
	aspect := float32(rect.Dx()) / float32(rect.Dy())
	return uint(float32(height) * aspect), height
}

// Utility function that calculates the width and height needed to resize an
// image to a maximum width or height, preserving its aspect ratio.
// It will return the original image width and height if these are already
// smaller than the provided maximum width and height.
func Thumbnail(rect image.Rectangle, maxWidth, maxHeight uint) (uint, uint) {
	aspect := float32(rect.Dx()) / float32(rect.Dy())
	if maxWidth >= uint(rect.Dx()) && maxHeight >= uint(rect.Dy()) {
		return uint(rect.Dx()), uint(rect.Dy())
	}

	if aspect >= 1 {
		return FromWidth(rect, maxWidth)
	}
	return FromHeight(rect, maxHeight)
}

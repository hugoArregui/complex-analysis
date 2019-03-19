package julia

import (
	"testing"
	"image"
)

func BenchmarkFullDrawJulia(b *testing.B) {
	im := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	p := Parameters{
		C:       complex(0.25, 0),
		MaxIter: 1000,
		XMin:    -5.0,
		XMax:    5.0,
		YMin:    -5.0,
		YMax:    5.0,
	}
	for i := 0; i < b.N; i++ {
		DrawJuliaSet(im, p)
	}
}

func BenchmarkFullExpensiveDrawJulia(b *testing.B) {
	im := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	p := Parameters{
		C:       complex(0.25, 0),
		MaxIter: 1000,
		XMin:    -1.0,
		XMax:    1.0,
		YMin:    -1.0,
		YMax:    1.0,
	}
	for i := 0; i < b.N; i++ {
		DrawJuliaSet(im, p)
	}
}

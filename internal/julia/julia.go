package julia

import (
	"math"
	"math/cmplx"

	"github.com/fogleman/gg"
)

type Parameters struct {
	C          complex128
	MaxIter    int
	XMin, XMax float64
	YMin, YMax float64
}

func DrawJuliaSet(dc *gg.Context, p Parameters) {
	r := 1 + math.Sqrt(1+4*cmplx.Abs(p.C))

	kx := (p.XMax - p.XMin) / float64(dc.Width())
	ky := (p.YMax - p.YMin) / float64(dc.Height())

	dc.SetRGB(0, 0, 0)
	dc.Clear()

	dc.SetRGB(1, 1, 1)
	for i := 0; i < dc.Width(); i++ {
		for j := 0; j < dc.Height(); j++ {
			x := p.XMin + float64(i)*kx
			y := p.YMin + float64(j)*ky
			z := complex(x, y)

			for n := 0; n < p.MaxIter; n++ {
				z0 := cmplx.Pow(z, 2) + p.C

				if cmplx.Abs(z0) > r {
					dc.SetPixel(i, j)
					break
				}

				z = z0
			}

		}
	}

	dc.Fill()
}

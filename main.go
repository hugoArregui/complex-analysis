package main

import (
	"math"
	"math/cmplx"

	"github.com/fogleman/gg"
)

type Parameters struct {
	c       complex128
	maxIter int
	a, b    float64
}

func graphJuliaSet(dc *gg.Context, p Parameters) {
	r := 1 + math.Sqrt(1+4*cmplx.Abs(p.c))

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			u := p.a + float64(float64(x)*(p.b-p.a))/1000.0
			v := p.a + float64(float64(y)*(p.b-p.a))/1000.0
			z := complex(u, v)

			toInf := false
			for n := 0; n < p.maxIter; n++ {
				z0 := cmplx.Pow(z, 2) + p.c

				if cmplx.Abs(z0) > r {
					toInf = true
					break
				}

				z = z0
			}

			if toInf {
				dc.SetRGB(1, 1, 1)
			} else {
				dc.SetRGB(0, 0, 0)
			}

			dc.SetPixel(x, y)
		}
	}

	dc.Fill()
}

func main() {
	dc := gg.NewContext(1000, 1000)
	p := Parameters{
		c:       complex(1.0/4.0, 0),
		maxIter: 1000,
		a:       -5.0,
		b:       5.0,
	}

	graphJuliaSet(dc, p)
	dc.SavePNG("out.png")
}

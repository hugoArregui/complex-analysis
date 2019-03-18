package julia

import (
	"math"
	"math/cmplx"
	"sync"

	"github.com/fogleman/gg"
)

type Parameters struct {
	C          complex128
	MaxIter    int
	XMin, XMax float64
	YMin, YMax float64
	r          float64
}

type region struct {
	XMin, XMax int
	YMin, YMax int
}

func drawRegion(dc *gg.Context, p Parameters, region region) {
	kx := (p.XMax - p.XMin) / float64(dc.Width())
	ky := (p.YMax - p.YMin) / float64(dc.Height())

	for j := region.YMin; j < region.YMax; j++ {
		y := p.YMin + float64(j)*ky
		for i := region.XMin; i < region.XMax; i++ {
			x := p.XMin + float64(i)*kx
			z := complex(x, y)

			for n := 0; n < p.MaxIter; n++ {
				z0 := cmplx.Pow(z, 2) + p.C

				if cmplx.Abs(z0) > p.r {
					dc.SetPixel(i, j)
					break
				}

				z = z0
			}

		}
	}
}

func DrawJuliaSet(dc *gg.Context, p Parameters) {
	p.r = 1 + math.Sqrt(1+4*cmplx.Abs(p.C))

	dc.SetRGB(0, 0, 0)
	dc.Clear()

	dc.SetRGB(1, 1, 1)

	splitX := 4
	splitY := 4
	var wg sync.WaitGroup
	wg.Add(splitX * splitY)

	// TODO this won't work if with or height are not divisible by their split
	regionWidth := dc.Width() / splitX
	regionHeight := dc.Height() / splitY
	for regionX := 0; regionX < splitX; regionX++ {
		for regionY := 0; regionY < splitY; regionY++ {
			r := region{}
			r.XMin = regionX * regionWidth
			r.XMax = r.XMin + regionWidth
			r.YMin = regionY * regionHeight
			r.YMax = r.YMin + regionHeight

			go func() {
				drawRegion(dc, p, r)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	dc.Fill()
}

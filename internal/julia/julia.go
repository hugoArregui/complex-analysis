package julia

import (
	"image"
	"image/color"
	"math"
	"math/cmplx"
	"math/rand"
	"sync"
)

type Parameters struct {
	C          complex128
	MaxIter    int
	XMin, XMax float64
	YMin, YMax float64
	r          float64
	colors     []color.Color
}

type region struct {
	XMin, XMax int
	YMin, YMax int
}

func drawRegion(im *image.RGBA, p Parameters, region region) {
	size := im.Bounds().Size()
	kx := (p.XMax - p.XMin) / float64(size.X)
	ky := (p.YMax - p.YMin) / float64(size.Y)

	for j := region.YMin; j < region.YMax; j++ {
		y := p.YMin + float64(j)*ky
		for i := region.XMin; i < region.XMax; i++ {
			x := p.XMin + float64(i)*kx
			z := complex(x, y)

			toInf := false
			for n := 0; n < p.MaxIter; n++ {
				z0 := cmplx.Pow(z, 2) + p.C

				if cmplx.Abs(z0) > p.r {
					toInf = true
					im.Set(i, j, p.colors[n])
					break
				}

				z = z0
			}

			if !toInf {
				im.Set(i, j, color.Black)
			}
		}
	}
}

func DrawJuliaSet(im *image.RGBA, p Parameters) {
	p.r = 1 + math.Sqrt(1+4*cmplx.Abs(p.C))
	splitX := 2
	splitY := 2

	p.colors = make([]color.Color, p.MaxIter)

	for colorIndex := 0; colorIndex < p.MaxIter; colorIndex++ {
		c := color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 255,
		}

		p.colors[colorIndex] = c
	}

	var wg sync.WaitGroup
	wg.Add(splitX * splitY)

	// TODO this won't work if with or height are not divisible by their split
	size := im.Bounds().Size()
	regionWidth := size.X / splitX
	regionHeight := size.Y / splitY
	for regionX := 0; regionX < splitX; regionX++ {
		for regionY := 0; regionY < splitY; regionY++ {
			r := region{}
			r.XMin = regionX * regionWidth
			r.XMax = r.XMin + regionWidth
			r.YMin = regionY * regionHeight
			r.YMax = r.YMin + regionHeight

			go func() {
				drawRegion(im, p, r)
				wg.Done()
			}()
		}
	}

	wg.Wait()
}

package julia

import (
	"testing"

	"github.com/fogleman/gg"
)

// BenchmarkFullDrawJulia-8   	       1	3782568577 ns/op
// ok  	github.com/hugoArregui/julia-sets/internal/julia	3.785s

func BenchmarkFullDrawJulia(b *testing.B) {
	dc := gg.NewContext(1000, 1000)
	p := Parameters{
		C:       complex(1.0/4.0, 0),
		MaxIter: 1000,
		XMin:    -5.0,
		XMax:    5.0,
		YMin:    -5.0,
		YMax:    5.0,
	}
	for i := 0; i < b.N; i++ {
		DrawJuliaSet(dc, p)
	}
}

package transform

import (
	"image"
	"math"

	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/parallel"
)

type RotationMatrix [6]float64

func GetRotationMatrix(centerX, centerY, angle, scale float64) RotationMatrix {
	r := RotationMatrix{}
	angle *= math.Pi / 180
	alpha := math.Cos(angle) * scale
	beta := math.Sin(angle) * scale
	r[0] = alpha
	r[1] = beta
	r[2] = (1-alpha)*centerX - beta*centerY
	r[3] = -beta
	r[4] = alpha
	r[5] = beta*centerX + (1-alpha)*centerY
	return r
}

func (r RotationMatrix) Reverse() RotationMatrix {
	inv := RotationMatrix{}
	scale := r[0]*r[4] - r[1]*r[3]
	if scale != 0 {
		scale = 1.0 / scale
	}
	inv[0] = r[4] * scale
	inv[1] = -r[1] * scale
	inv[2] = -inv[0]*r[2] - inv[1]*r[5]
	inv[3] = -r[3] * scale
	inv[4] = r[0] * scale
	inv[5] = -inv[3]*r[2] - inv[4]*r[5]
	return inv
}

func (r RotationMatrix) Rotate(srcX, srcY float64) (dstX, dstY float64) {
	dstX = r[0]*srcX + r[1]*srcY + r[2]
	dstY = r[3]*srcX + r[4]*srcY + r[5]
	return
}

func WrapAffine(img image.Image, r RotationMatrix) *image.RGBA {
	src := clone.AsRGBA(img)
	srcW, srcH := src.Bounds().Dx(), src.Bounds().Dy()
	srcStride := src.Stride
	dst := image.NewRGBA(image.Rect(0, 0, srcW, srcH))
	inv := r.Reverse()
	parallel.Line(srcH, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < srcW; x++ {
				srcX, srcY := inv.Rotate(float64(x), float64(y))
				if srcX < 0 || srcX > float64(srcW-1) || srcY < 0 || srcY > float64(srcH-1) {
					continue
				}
				srcPos := int(srcY+0.5)*srcStride + int(srcX+0.5)*4
				dstPos := y*srcStride + x*4
				copy(dst.Pix[dstPos:dstPos+4], src.Pix[srcPos:srcPos+4])
			}
		}
	})
	return dst
}

package gorgb

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
)

func ensureRGBA(im image.Image) *image.RGBA {
	switch im := im.(type) {
	case *image.RGBA:
		return im
	default:
		dst := image.NewRGBA(im.Bounds())
		draw.Draw(dst, im.Bounds(), im, image.ZP, draw.Src)
		return dst
	}
}

func Convert(im image.Image) image.Image {
	rgba := ensureRGBA(im)
	tree := NewOctree()
	indexes := rand.Perm(numColors)
	w := rgba.Bounds().Size().X
	for _, i := range indexes {
		x := i % w
		y := i / w
		c := rgba.RGBAAt(x, y)
		r := int(c.R)
		g := int(c.G)
		b := int(c.B)
		r, g, b = tree.Pop(r, g, b)
		c.R = uint8(r)
		c.G = uint8(g)
		c.B = uint8(b)
		rgba.SetRGBA(x, y, c)
	}
	return rgba
}

func Verify(im image.Image) error {
	rgba := ensureRGBA(im)
	w := rgba.Bounds().Size().X
	h := rgba.Bounds().Size().Y
	seen := make(map[color.RGBA]bool)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := rgba.RGBAAt(x, y)
			if _, ok := seen[c]; ok {
				return fmt.Errorf("image contains duplicate colors")
			}
			seen[c] = true
		}
	}
	if w != 4096 || h != 4096 {
		return fmt.Errorf("colors are distinct but image is not 4096x4096")
	}
	return nil
}

package texture

import (
	"image"
	"image/draw"
	"log"
	"os"

	"github.com/go-gl/gl/v2.1/gl"
)

type Name string

const (
	GrassSide   Name = "grass_side"
	GrassTop    Name = "grass_top"
	GrassBottom Name = "grass_bottom"
)

type Texture struct {
	Bucket map[Name]uint32
}

func NewTexture() *Texture {
	return &Texture{
		Bucket: map[Name]uint32{},
	}
}

func (t *Texture) Load() {
	t.Bucket[GrassSide] = newTexture("./assets/grass_side.png")
	t.Bucket[GrassTop] = newTexture("./assets/grass_top.png")
	t.Bucket[GrassBottom] = newTexture("./assets/gold_block.png")
}

func (t Texture) Delete() {
	for _, value := range t.Bucket {
		gl.DeleteTextures(1, &value)
	}
}

func newTexture(file string) uint32 {
	imgFile, err := os.Open(file)
	if err != nil {
		log.Fatalf("texture %q not found on disk: %v\n", file, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture
}

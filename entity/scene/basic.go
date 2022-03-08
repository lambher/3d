package scene

import (
	"encoding/json"
	"io"
	"os"

	"github.com/lambher/3d/models"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/lambher/3d/context"
)

type Basic struct {
	*Scene
}

func NewBasic(ctx *context.Context, path string) (*Basic, error) {
	basic := &Basic{
		Scene: &Scene{
			ctx: ctx,
		},
	}

	if path == "" {
		path = "./assets/maps/first_world/world.json"
	}

	scene, err := LoadSceneFromFile(path)
	if err != nil {
		return nil, err
	}

	basic.LoadScene(ctx, scene)
	return basic, nil
}

func LoadSceneFromFile(path string) (models.Scene, error) {
	file, err := os.Open(path)
	if err != nil {
		return models.Scene{}, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return models.Scene{}, err
	}
	var scene models.Scene
	err = json.Unmarshal(data, &scene)
	if err != nil {
		return models.Scene{}, err
	}
	return scene, nil
}

func (b Basic) Setup() {
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.LIGHTING)

	gl.ClearColor(0.5, 0.5, 0.5, 0.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)

	ambient := []float32{0.5, 0.5, 0.5, 1}
	diffuse := []float32{1, 1, 1, 1}
	lightPosition := []float32{b.Scene.Light.Position.X, b.Scene.Light.Position.Y, b.Scene.Light.Position.Z, 0}
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[0])
	gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPosition[0])
	gl.Enable(gl.LIGHT0)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	f := float64(b.Camera.Width)/float64(b.Camera.Height) - 1
	gl.Frustum(-1-f, 1+f, -1, 1, 1.0, 10.0)
	gl.Translatef(b.Camera.Position.X(), b.Camera.Position.Y(), b.Camera.Position.Z())
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func (b Basic) Draw() {
	b.Scene.Draw()
}

func (b *Basic) Update() {
	b.Scene.Update()
}

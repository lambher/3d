package maps

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/lambher/3d/models"
)

func Seed() error {
	scene := GenerateScene()
	data, err := json.MarshalIndent(scene, "", "\t")
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("./assets/maps/seed/%d.json", time.Now().Unix())
	err = os.WriteFile(fileName, data, os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Println("new map!", fileName)
	return nil
}

func GenerateScene() models.Scene {
	var scene models.Scene

	scene.Camera = models.Camera{
		View: models.Vector2{
			X: 1600,
			Y: 900,
		},
	}
	scene.Light = models.Light{}

	for x := 0; x < 100; x++ {
		for z := 0; z < 100; z++ {
			scene.Cubes = append(scene.Cubes, models.Cube{
				Position: models.Vector3{
					X: float32(x),
					Y: 0,
					Z: float32(z),
				},
				Type: models.Grass,
			})
		}
	}
	return scene
}

package world

import (
	"math/rand"

	"github.com/aquilax/go-perlin"
)

type Asset int

const (
	AssetGrass Asset = iota
	AssetGround
	AssetPath
	AssetSand
	AssetWater
)

type Generator struct {

}

// GenerateRandomWorld generates a random world
func (w *Generator) GenerateRandomWorld() [1000][1000]Asset{
	const (
		alpha       = 2.
		beta        = 2.
		n           = 3
		seed  int64 = 100
	)
	p := perlin.NewPerlinRandSource(alpha, beta, n, rand.NewSource(seed))
	var r float64
	var data Asset
	var world [1000][1000]Asset

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			r = p.Noise2D(float64(i)/10, float64(j)/10)
			switch {
			case r <= 0.0:
				data = AssetGrass
			case r <= 0.25:
				data = AssetGround
			case r <= 0.5:
				data = AssetPath
			case r <= 0.6:
				data = AssetSand
			case r <= 1:
				data = AssetWater
			}
			world[i][j] = data
		}
	}

	return world
}

package assets

import (
	"softcraft/pkg/common"
	"softcraft/pkg/world"

	"github.com/veandco/go-sdl2/sdl"
)

type AssetLoader struct {

}

func (l *AssetLoader) LoadWorldAssets(renderer *sdl.Renderer) map[world.Asset]*sdl.Texture {
	var textures map[world.Asset]*sdl.Texture
	textures = make(map[world.Asset]*sdl.Texture,5)

	textures[world.AssetGrass] =  l.loadAsset("assets/world/grass.bmp", renderer)
	textures[world.AssetGround] =  l.loadAsset("assets/world/ground.bmp", renderer)
	textures[world.AssetPath] =  l.loadAsset("assets/world/path.bmp", renderer)
	textures[world.AssetSand] =  l.loadAsset("assets/world/sand.bmp", renderer)
	textures[world.AssetWater] =  l.loadAsset("assets/world/water.bmp", renderer)

	return textures
}

func (l *AssetLoader) loadAsset(fileName string, renderer *sdl.Renderer) *sdl.Texture {
	t, err := common.LoadTextureFromBMP(fileName, renderer)
	if err != nil {
		// No point in continue if we can't load the assets
		panic(err)
	}
	return t
}

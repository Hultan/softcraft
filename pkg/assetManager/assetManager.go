package assetManager

import (
	"softcraft/pkg/common"

	"github.com/veandco/go-sdl2/sdl"
)

type AssetManager struct {
	assets map[AssetString]*sdl.Texture
}

// New creates a new AssetManager
func New() *AssetManager {
	a := &AssetManager{}
	a.assets = make(map[AssetString]*sdl.Texture,100)
	return a
}

var assetLoadingErr error

// Load loads all assets from file
func (a *AssetManager) Load(renderer *sdl.Renderer) error {
	assetLoadingErr = nil

	a.assets[AssetStringWorldGrass] = a.loadAsset(AssetPathWorldGrass, renderer)
	a.assets[AssetStringWorldGround] = a.loadAsset(AssetPathWorldGround, renderer)
	a.assets[AssetStringWorldPath] = a.loadAsset(AssetPathWorldPath, renderer)
	a.assets[AssetStringWorldSand] = a.loadAsset(AssetPathWorldSand, renderer)
	a.assets[AssetStringWorldWater] = a.loadAsset(AssetPathWorldWater, renderer)
	a.assets[AssetStringPlayer] = a.loadAsset(AssetPathPlayer, renderer)

	return assetLoadingErr
}

// GetAsset returns a specific asset
func (a *AssetManager) GetAsset(name AssetString) *sdl.Texture {
	return a.assets[name]
}

// GetWorldAsset returns a specific world asset
func (a *AssetManager) GetWorldAsset(wa AssetMap) *sdl.Texture {
	switch wa {
	case AssetMapGrass: return a.assets[AssetStringWorldGrass]
	case AssetMapGround: return a.assets[AssetStringWorldGround]
	case AssetMapPath: return a.assets[AssetStringWorldPath]
	case AssetMapSand: return a.assets[AssetStringWorldSand]
	case AssetMapWater: return a.assets[AssetStringWorldWater]
	default: return nil
	}
}

// loadAsset loads a single asset
func (a *AssetManager) loadAsset(fileName string, renderer *sdl.Renderer) *sdl.Texture {
	t, err := common.LoadTextureFromBMP(fileName, renderer)
	assetLoadingErr = err
	return t
}

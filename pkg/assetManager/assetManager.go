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

	a.assets[AssetStringGrass] = a.loadAsset(AssetPathWorldGrass, renderer)
	a.assets[AssetStringGround] = a.loadAsset(AssetPathWorldGround, renderer)
	a.assets[AssetStringPath] = a.loadAsset(AssetPathWorldPath, renderer)
	a.assets[AssetStringSand] = a.loadAsset(AssetPathWorldSand, renderer)
	a.assets[AssetStringWater] = a.loadAsset(AssetPathWorldWater, renderer)

	return assetLoadingErr
}

// GetAsset returns a specific asset
func (a *AssetManager) GetAsset(name AssetString) *sdl.Texture {
	return a.assets[name]
}

// GetWorldAsset returns a specific world asset
func (a *AssetManager) GetWorldAsset(wa AssetNumeric) *sdl.Texture {
	switch wa {
	case AssetNumericGrass: return a.assets[AssetStringGrass]
	case AssetNumericGround: return a.assets[AssetStringGround]
	case AssetNumericPath: return a.assets[AssetStringPath]
	case AssetNumericSand: return a.assets[AssetStringSand]
	case AssetNumericWater: return a.assets[AssetStringWater]
	default: return nil
	}
}

// loadAsset loads a single asset
func (a *AssetManager) loadAsset(fileName string, renderer *sdl.Renderer) *sdl.Texture {
	t, err := common.LoadTextureFromBMP(fileName, renderer)
	assetLoadingErr = err
	return t
}

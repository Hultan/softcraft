package assetManager

type AssetString string

const (
	AssetStringGrass  AssetString = "world.grass"
	AssetStringGround AssetString = "world.ground"
	AssetStringPath   AssetString = "world.path"
	AssetStringSand   AssetString = "world.sand"
	AssetStringWater  AssetString = "world.water"
)

const (
	AssetPathWorldGrass  = "assets/world/grass.bmp"
	AssetPathWorldGround = "assets/world/ground.bmp"
	AssetPathWorldPath   = "assets/world/path.bmp"
	AssetPathWorldSand   = "assets/world/sand.bmp"
	AssetPathWorldWater  = "assets/world/water.bmp"
)

type AssetNumeric int

const (
	AssetNumericGrass AssetNumeric = iota
	AssetNumericGround
	AssetNumericPath
	AssetNumericSand
	AssetNumericWater
)

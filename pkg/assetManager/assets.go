package assetManager

type AssetString string

const (
	AssetStringWorldGrass  AssetString = "world.grass"
	AssetStringWorldGround AssetString = "world.ground"
	AssetStringWorldPath   AssetString = "world.path"
	AssetStringWorldSand   AssetString = "world.sand"
	AssetStringWorldWater  AssetString = "world.water"
	AssetStringPlayer      AssetString = "player"
)

const (
	AssetPathWorldGrass  = "assets/world/grass.bmp"
	AssetPathWorldGround = "assets/world/ground.bmp"
	AssetPathWorldPath   = "assets/world/path.bmp"
	AssetPathWorldSand   = "assets/world/sand.bmp"
	AssetPathWorldWater  = "assets/world/water.bmp"
	AssetPathPlayer      = "assets/player.bmp"
)

type AssetMap int

const (
	AssetMapGrass AssetMap = iota
	AssetMapGround
	AssetMapPath
	AssetMapSand
	AssetMapWater
)

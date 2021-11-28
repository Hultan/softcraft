package world

import (
	"io/ioutil"
	"strconv"
	"strings"

	"softcraft/pkg/assetManager"
)

type Loader struct {}

var assetLoadingErr error

// LoadWorld loads the world from the file assets/softcraft.world
func (l *Loader) LoadWorld() [][]assetManager.AssetMap {
	var w [][]assetManager.AssetMap

	data, err := ioutil.ReadFile("assets/softcraft.world")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data),"\n")
	for _, line := range lines {
		if strings.Trim(line, " ") == "" {
			continue
		}
		var row []assetManager.AssetMap
		blocks := strings.Split(line, " ")
		for _, block := range blocks {
			item, err := strconv.Atoi(block)
			if err != nil {
				panic(err)
			}
			row = append(row, assetManager.AssetMap(item))
		}
		w = append(w,row)
	}
	return w
}

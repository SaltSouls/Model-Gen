package modelTypes

import (
	"BlockGenerator/main/helpers"
)

// define all the structures used to build the couch model variants
type couchGeneric struct {
	Particle  string `json:"particle" binding:"required"`
	Top       string `json:"top" binding:"required"`
	Front     string `json:"front" binding:"required"`
	SideOuter string `json:"side_outer" binding:"required"`
	SideInner string `json:"side_inner" binding:"required"`
	Back      string `json:"back" binding:"required"`
	Bottom    string `json:"bottom" binding:"required"`
}

type couchMiddle struct {
	Particle string `json:"particle" binding:"required"`
	Top      string `json:"top" binding:"required"`
	Front    string `json:"front" binding:"required"`
	Back     string `json:"back" binding:"required"`
	Bottom   string `json:"bottom" binding:"required"`
}

type couchCorner struct {
	Particle string `json:"particle" binding:"required"`
	Top      string `json:"top" binding:"required"`
	Side     string `json:"side" binding:"required"`
	Back     string `json:"back" binding:"required"`
	Bottom   string `json:"bottom" binding:"required"`
}

// define all structures used to build the couch couchState variants
type couchState struct {
	Model string `json:"model" binding:"required"`
	Y     int    `json:"y,omitempty"`
}

// BuildCouches creates couch model variants based on color/wood/both.
func BuildCouches(nameBase string, texturePath string, fileName string, mainDir string) {
	space, couchType := "_", [3]string{"single", "left", "right"}
	// loop through all couch types
	for t := 0; t < 3; t++ {
		if t == 0 {
			for i := 0; i < len(couchType); i++ {
				parentName, textureType := nameBase+space+couchType[i], "single"
				if couchType[i] == "single" {
					textureType = "single"
				} else {
					textureType = "side"
				}
				// construct template for single/left/right couch model variants
				generic := couchGeneric{
					Particle:  "#front",
					Top:       texturePath + space + textureType + "_top",
					Front:     texturePath + space + textureType + "_front",
					SideOuter: texturePath + "_side_outer",
					SideInner: texturePath + "_side_inner",
					Back:      texturePath + space + textureType + "_back",
					Bottom:    nameBase + space + textureType + "_bottom"}
				// convert template info into json format
				msg := helpers.GenericParent[couchGeneric]{Parent: parentName, Textures: generic}
				helpers.WriteModelFile(msg, mainDir, fileName+space+couchType[i]+".json")
			}
		} else if t == 1 {
			// construct template for middle couch model variants
			parentName, middle := nameBase+"_middle", couchMiddle{
				Particle: "#front",
				Top:      texturePath + "_middle_top",
				Front:    texturePath + "_middle_front",
				Back:     texturePath + "_middle_back",
				Bottom:   nameBase + "_middle_bottom"}
			// convert template info into json format
			message := helpers.GenericParent[couchMiddle]{Parent: parentName, Textures: middle}
			helpers.WriteModelFile(message, mainDir, fileName+"_middle.json")
		} else {
			// construct template for corner couch model variants
			parentName, corner := nameBase+"_corner", couchCorner{
				Particle: "#side",
				Top:      texturePath + "_corner_front",
				Side:     texturePath + "_corner",
				Back:     texturePath + "_corner_back",
				Bottom:   nameBase + "_corner_bottom"}
			// convert template info into json format
			msg := helpers.GenericParent[couchCorner]{Parent: parentName, Textures: corner}
			helpers.WriteModelFile(msg, mainDir, fileName+"_corner.json")
		}
	}
}

func buildTypes(keyBase string, y int, filePath string, blockstates *map[string]couchState) {
	for t := 0; t < 6; t++ {
		switch t {
		case 0:
			key, modelPath := keyBase+"single", filePath+"_single"
			(*blockstates)[key] = couchState{Model: modelPath, Y: y}
			break
		case 1:
			key, modelPath := keyBase+"left", filePath+"_left"
			(*blockstates)[key] = couchState{Model: modelPath, Y: y}
			break
		case 2:
			key, modelPath := keyBase+"right", filePath+"_right"
			(*blockstates)[key] = couchState{Model: modelPath, Y: y}
			break
		case 3:
			key, modelPath := keyBase+"middle", filePath+"_middle"
			(*blockstates)[key] = couchState{Model: modelPath, Y: y}
			break
		case 4:
			key, modelPath := keyBase+"corner_left", filePath+"_corner"
			(*blockstates)[key] = couchState{Model: modelPath, Y: y}
			break
		case 5:
			key, modelPath := keyBase+"corner_right", filePath+"_corner"
			y = y + 90
			if y == 360 {
				y = 0
			}
			(*blockstates)[key] = couchState{Model: modelPath, Y: y}
			break
		default:
			return
		}
	}
}

// BuildCouchStates creates couch blockstate variants based on color/wood/both.
func BuildCouchStates(filePath string, fileName string, mainDir string) {
	blockstates := make(map[string]couchState)
	keyBase := "facing="
	// construct the blockstates template for the different variants
	for d := 0; d < 4; d++ {
		switch d {
		case 0:
			y, key := 0, keyBase+"north,type="
			buildTypes(key, y, filePath, &blockstates)
			break
		case 1:
			key, y := keyBase+"south,type=", 180
			buildTypes(key, y, filePath, &blockstates)
			break
		case 2:
			key, y := keyBase+"east,type=", 90
			buildTypes(key, y, filePath, &blockstates)
			break
		case 3:
			key, y := keyBase+"west,type=", 270
			buildTypes(key, y, filePath, &blockstates)
			break
		default:
			return
		}
	}
	msg := helpers.GenericVariant[couchState]{Variants: blockstates}
	helpers.WriteStatesFile(msg, mainDir, fileName+".json")
}

package main

import (
	"ModelGen/main/helpers"
	"ModelGen/main/modelTypes"
	"ModelGen/main/parser"
	"path/filepath"
	"strings"
)

func createModels(model string, texturePath string, fileName string) {
	mainDir := helpers.GetMainDir()
	parentName := helpers.GetParentName()
	parentName = strings.ReplaceAll(parentName, "\\", "/")
	switch model {
	case "couch":
		helpers.CreateDirectory(mainDir)
		modelTypes.BuildCouches(parentName, texturePath, fileName, mainDir)
		break
	default:
		return
	}
}

func createStates(model string, fileName string, filePath string, mainDirBase string) {
	mainDir := filepath.Join(mainDirBase + "/blockstates/")
	helpers.CreateDirectory(mainDir)
	switch model {
	case "couch":
		modelTypes.BuildCouchStates(filePath, fileName, mainDir)
		break
	case "chair":
		helpers.CreateDirectory(mainDir)
		break
	default:
		return
	}
}

func main() {
	parser := parser.NewParser()
	modID, model, name, mode, path, genericPath, space :=
		strings.ToLower(parser.GetID()),
		strings.ToLower(parser.GetModel()),
		strings.ToLower(parser.GetName()),
		strings.ToTitle(parser.GetMode()),
		parser.GetPath(),
		helpers.GetGenericPath(),
		"_"
	genericPath = strings.ReplaceAll(genericPath, "\\", "/")
	if name == " " {
		name = strings.ReplaceAll(name, " ", "")
	} else {
		name = strings.ReplaceAll(name, " ", "_")
	}
	colors, woodTypes :=
		[16]string{"white", "light_gray", "gray", "black", "pink", "red", "magenta", "purple", "blue", "light_blue", "cyan", "green", "lime", "yellow", "orange", "brown"},
		[9]string{"oak", "spruce", "birch", "jungle", "acacia", "dark_oak", "mangrove", "crimson", "warped"}
	baseDir := path + "/assets/" + modID
	var fileName, texturePath string
	switch mode {
	case "WOOD":
		for i := 0; i < len(woodTypes); i++ {
			if name != "" {
				fileName = woodTypes[i] + space + name
				texturePath = genericPath + woodTypes[i] + space + name
			} else {
				fileName = woodTypes[i] + space + model
				texturePath = genericPath + woodTypes[i] + space + model
			}
			filePath := genericPath + fileName
			createModels(model, texturePath, fileName)
			createStates(model, fileName, filePath, baseDir)
		}
		break
	case "BOTH":
		for i := 0; i < len(woodTypes); i++ {
			for j := 0; j < len(colors); j++ {
				if name != "" {
					fileName = colors[j] + space + woodTypes[i] + space + name
					texturePath = genericPath + colors[j] + space + woodTypes[i] + space + name
				} else {
					fileName = colors[j] + space + woodTypes[i] + space + model
					texturePath = genericPath + colors[j] + space + woodTypes[i] + space + model
				}
				filePath := genericPath + fileName
				createModels(model, texturePath, fileName)
				createStates(model, fileName, filePath, baseDir)
			}
		}
		break
	default:
		for i := 0; i < len(colors); i++ {
			if name != "" {
				fileName = colors[i] + space + name
				texturePath = genericPath + colors[i] + space + name
			} else {
				fileName = colors[i] + space + model
				texturePath = genericPath + colors[i] + space + model
			}
			filePath := genericPath + fileName
			createModels(model, texturePath, fileName)
			createStates(model, fileName, filePath, baseDir)
		}
	}
}

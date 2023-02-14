package helpers

import (
	"BlockGenerator/main/parser"
	"path/filepath"
	"strings"
)

// params used to build and format the paths
type params struct {
	mainDir     string
	parentName  string
	genericPath string
}

// formats the path based on user inputs
func pathFormat(baseDir string, baseName string, genericPath string, model string, name string, resLoc string) params {
	var mainDir, parentName string
	// test if resource location is blank, if so, build paths accordingly
	switch resLoc {
	case "":
		if name != "" {
			mainDir = filepath.Join(baseDir, name)
			parentName = baseName + name + "/" + name
			genericPath = baseName + name + "/"
		} else {
			mainDir = filepath.Join(baseDir + "/" + model + "/")
			parentName = baseName + model + "/" + model
			genericPath = baseName + model + "/"
		}
		break
	default:
		mainDir = baseDir
		genericPath = baseName
		if name != "" {
			parentName = baseName + name

		} else {
			parentName = baseName + model
		}
	}
	// return params to be called in buildPaths
	return params{
		mainDir:     mainDir,
		parentName:  parentName,
		genericPath: genericPath,
	}
}

// builds the path used to look for/store files in
func buildPaths(modID string, model string, name string, baseDir string) params {
	baseDir = filepath.Join(baseDir + "/models/block/")
	baseName := modID + ":block/"
	if name == " " {
		name = strings.ReplaceAll(name, " ", "")
	} else {
		name = strings.ReplaceAll(name, " ", "_")
	}
	var (
		mainDir, parentName, genericPath string
		// explicit declaration
		parser = parser.NewParser()
		resLoc = parser.GetResourceLocation()
		path   = pathFormat(baseDir, baseName, genericPath, model, name, resLoc)
	)
	// test if resource location is empty, if not, build paths accordingly
	switch resLoc {
	case "":
		mainDir = path.mainDir
		parentName = path.parentName
		genericPath = path.genericPath
		break
	default:
		if resLoc == " " {
			resLoc = ""
			mainDir = path.mainDir
			parentName = path.parentName
			genericPath = path.genericPath
		} else {
			resLoc = strings.ReplaceAll(resLoc, " ", "_")
			mainDir = filepath.Dir(baseDir + "/" + resLoc + "/")
			parentName = baseName + resLoc
			genericPath = path.genericPath + resLoc + "/"
		}
	}
	// return params to be called in Get functions
	return params{
		mainDir:     mainDir,
		parentName:  parentName,
		genericPath: genericPath,
	}
}

func GetMainDir() string {
	// get/set required variables and return mainDir
	parser := parser.NewParser()
	modID, model, name, path :=
		strings.ToLower(parser.GetID()),
		strings.ToLower(parser.GetModel()),
		strings.ToLower(parser.GetName()),
		parser.GetPath()
	baseDir := path + "/assets/" + modID
	return buildPaths(modID, model, name, baseDir).mainDir
}

func GetParentName() string {
	// get/set required variables and return parentName
	parser := parser.NewParser()
	modID, model, name, path :=
		strings.ToLower(parser.GetID()),
		strings.ToLower(parser.GetModel()),
		strings.ToLower(parser.GetName()),
		parser.GetPath()
	baseDir := path + "/assets/" + modID
	return buildPaths(modID, model, name, baseDir).parentName
}

func GetGenericPath() string {
	// get/set required variables and return texturePath
	parser := parser.NewParser()
	modID, model, name, path :=
		strings.ToLower(parser.GetID()),
		strings.ToLower(parser.GetModel()),
		strings.ToLower(parser.GetName()),
		parser.GetPath()
	baseDir := path + "/assets/" + modID
	return buildPaths(modID, model, name, baseDir).genericPath
}

package helpers

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

// Everything here is used in the creation of the json model/states files. These
// should not be touched under any circumstance as they are designed to be universal.
// Changing them could cause just about everything to break, so don't do it.

// handle errors
func check(err error) {
	if err != nil {
		log.Fatal("Something went wrong! ", err)
	}
}

// GenericParent is used to define base model inheritance.
type GenericParent[T any] struct {
	Parent   string `json:"parent" binding:"required"`
	Textures T      `json:"textures" binding:"required"`
}

// GenericVariant is used to define base variant used in blockstates.
type GenericVariant[T any] struct {
	Variants map[string]T `json:"variants" binding:"required"`
}

// CreateDirectory creates directory structure used in modding environments.
func CreateDirectory(mainDir string) {
	// TODO maybe check necessary file perms required?
	err := os.MkdirAll(mainDir, 0666)
	check(err)
}

// WriteModelFile creates and write model info into the json format.
func WriteModelFile[T any](message GenericParent[T], mainDir string, extension string) {
	// convert template info into json format
	jsonMessage, err := json.MarshalIndent(message, "", "\t")
	check(err)
	// create json file with corresponding model name then close the file
	jsonFile, err := os.Create(filepath.Join(mainDir, extension))
	check(err)
	defer jsonFile.Close()
	// write json to the newly created file
	_, err = jsonFile.Write(jsonMessage)
	check(err)
}

// WriteStatesFile creates and writes blockstate info into json format.
func WriteStatesFile[T any](message GenericVariant[T], mainDir string, extension string) {
	// convert template info into json format
	jsonMessage, err := json.MarshalIndent(message, "", "\t")
	check(err)
	// create json file with corresponding model name then close the file
	jsonFile, err := os.Create(filepath.Join(mainDir, extension))
	check(err)
	defer jsonFile.Close()
	// write json to the newly created file
	_, err = jsonFile.Write(jsonMessage)
	check(err)
}

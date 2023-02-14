package parser

import (
	"errors"
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"strings"
)

type CommandOptions struct {
	modID string
	model string
	// Not required
	condition string
	// Not required
	name string
	mode string
	path string
	// Not required
	resLoc string
}

func NewParser() CommandOptions {
	parser := argparse.NewParser("generate", "Helps with the creation of variant block models and states.")
	modID, model, condition, name, mode, path, resLoc, err :=
		parser.String("i", "mod-id", &argparse.Options{
			Required: true,
			Help:     "Mod ID used by the models",
			Default:  "minecraft"}),
		parser.String("t", "model-type", &argparse.Options{Required: true,
			Help:    "The model variants will be made for",
			Default: "couch"}),
		parser.String("s", "special-option", &argparse.Options{
			Required: false,
			Validate: func(args []string) error {
				special := strings.ToTitle(args[0])
				if special == "NONE" || special == "NOBACK" {
					return nil
				}
				return errors.New("valid option(s) are [NONE|NOBACK]")
			},
			Help:    "Specifies special model conditions\n\t\t\t   Accepted values: NONE|NOBACK",
			Default: "NONE"}),
		parser.String("n", "name", &argparse.Options{
			Required: false,
			Help:     "The optional name you want to give the model",
			Default:  ""}),
		parser.String("m", "mode", &argparse.Options{
			Required: true,
			Validate: func(args []string) error {
				mode := strings.ToTitle(args[0])
				if mode == "COLOR" || mode == "WOOD" || mode == "BOTH" {
					return nil
				}
				return errors.New("valid mode(s) are [COLOR|WOOD|BOTH]")
			},
			Help:    "Creates variants based on desired mode\n\t\t\t   Accepted values: COLOR|WOOD|BOTH",
			Default: "COLOR"}),
		parser.String("p", "path", &argparse.Options{
			Required: true,
			Help:     "The directory to store the saved files in",
			Default:  ""}),
		parser.String("d", "resource-location", &argparse.Options{
			Required: false,
			Help:     "The location the resources are stored/located in. Otherwise they will be stored/looked for in block/<model name>",
			Default:  ""}),
		parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	return CommandOptions{
		modID:     *modID,
		model:     *model,
		condition: *condition,
		name:      *name,
		mode:      *mode,
		path:      *path,
		resLoc:    *resLoc,
	}
}

func (o CommandOptions) GetID() string {
	return o.modID
}
func (o CommandOptions) GetModel() string {
	return o.model
}
func (o CommandOptions) GetCondition() string {
	return o.condition
}
func (o CommandOptions) GetName() string {
	return o.name
}
func (o CommandOptions) GetMode() string {
	return o.mode
}
func (o CommandOptions) GetPath() string {
	return o.path
}
func (o CommandOptions) GetResourceLocation() string {
	return o.resLoc
}

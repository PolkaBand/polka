package command

import (
	"fmt"
	"github.com/PolkaBand/polka/config"
	"github.com/PolkaBand/polka/utils"
	"github.com/mitchellh/cli"
	//"log"
	//"os"
	"strings"
	//"text/template"
)

type ConfigCommand struct {
	Name   string
	Ui     cli.Ui
}

func (c *ConfigCommand) Help() string {
	helpText := `
Usage: polka config [item]

Configure a global polka item that applies to the entire app.

Items:

	s3	the root bucket for this app 

`
	return strings.TrimSpace(helpText)
}

func (c *ConfigCommand) Run(args []string) int {
	if _, err := utils.AreWeInProjectDir(); err != nil {
		fmt.Printf("%v", NotInAppDirectoryMessage())
		return 1
	} else {
		// config.CreateProjectConfigAsNeeded(appDir)  //TODO needed?
	}

	if len(args) < 2 {
		fmt.Printf("%v", c.Help())
		return 1
	}

	subcommand := args[0]
	//assume that we're itn the base polka directory
	//confirm that by checking to see if ./app/ exists

	switch subcommand {
		case "s3":
		//e.g. polka config s3 s3://SomeUrlHere/Etc
		c.configureS3(args[1])
		return 0
	}
	return 1
}

func (c *ConfigCommand) Synopsis() string {
	return "configure a global polka app item"
}

func (c *ConfigCommand) configureS3(s3url string) {
	if config, err := config.LoadProjectConfig(); err == nil {
		config.S3 = s3url
		config.Save()
	}
}

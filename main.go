package main

import (
	"github.com/urfave/cli"
	"os"
	"fmt"
  "log"
  "github.com/Flipkart/artcli/utils"
  "github.com/Flipkart/artcli/commands/search"
	"github.com/Flipkart/artcli/commands/create"
	"github.com/Flipkart/artcli/prehook"
)

//go:generate go-bindata -pkg bindata -o util/bindata/bindata.go config/

func main() {
  var (
    fileDescriptor *os.File
  )
	app := cli.NewApp()
	app.Name = "ART Repo Service"
	app.Usage = "To use ART service from cmd"
	app.Version = "1.0.0"
	app.EnableBashCompletion = true
	app.Commands = getCommands()
	app.Before = func(c *cli.Context) error{
		 prehook.Prehook()
	   utils.GetConfigValues()
	   fileDescriptor = utils.SetUpLog()
	 	 return nil
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Printf("%s: '%s' is not a valid command. See '%s --help'.", c.App.Name, command, os.Args[0])
		//prints the stack trace to log file
		utils.PrintStackTraceToLogFile()
		log.Fatalf("%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, os.Args[0])
	}
	app.After = func(c *cli.Context) error {
		 defer fileDescriptor.Close()
		 return nil
	}
	app.Run(os.Args)
}

func getCommands() []cli.Command {
	commands := []cli.Command{
		{
			Name:    "search",
			Aliases: []string{"S"},
			Usage:   "Used to search a single artifact in the ART repo",
			Flags:   search.GetSearchFlags(),
			Action:  search.SearchArtifactory,
		},
		{
			Name:    "create",
			Aliases: []string{"C"},
			Usage:   "Add an artifact to ART repository",
			Flags:   create.GetCreateFlags(),
			Action:  create.AddArtifact,
		},
	}
	return commands
}

package main

import (
	"os"

	"github.com/shnellpavel/json-stream/jsonstream-cli/cmd"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("jsonstream", "Utils to process and analyze stream of json")

	filterCommand := cmd.NewFilter()
	filterCmd := app.Command("filter", "Filters json stream by conditions").Action(filterCommand.Run)
	filterCommand.InitArgs(filterCmd)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}

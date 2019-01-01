package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/shnellpavel/json-stream/jsonstream/filter"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// FilterCommand represents command to filter stream
type FilterCommand struct {
	condition string
}

// NewFilter constructs FilterCommand
func NewFilter() *FilterCommand {
	return &FilterCommand{}
}

// InitArgs initialize arguments and flags to run command
func (c *FilterCommand) InitArgs(cmd *kingpin.CmdClause) {
	cmd.Flag("condition", "expression with condition").
		Required().
		StringVar(&c.condition)
}

// Run handles command execution
func (c *FilterCommand) Run(_ *kingpin.ParseContext) error {
	info, err := os.Stdin.Stat()
	if err != nil {
		return errors.Wrap(err, "check stdin stat error")
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		return errors.New("The command is intended to work with pipes")
	}

	reader := bufio.NewReader(os.Stdin)

	filterExpr, err := filter.NewConditionFromStr(c.condition)
	if err != nil {
		return errors.Wrap(err, "parse filter error")
	}

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		if err != nil {
			return errors.Wrap(err, "read line error")
		}

		resLine, isOk, err := filter.ProcessElem(*filterExpr, line)
		if err != nil {
			return errors.Wrap(err, "process line error")
		}

		if isOk {
			fmt.Println(resLine)
		}
	}

	return nil
}

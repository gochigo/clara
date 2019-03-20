package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/gobuffalo/plush"
)

const (
	SUCCESS = "\u2713"
	ERROR   = "\u2718"
	WARNING = "\u26A0"
)

func Warning(help plush.HelperContext) (string, error) {
	return colorize(color.YellowString, help, WARNING)
}

func Error(help plush.HelperContext) (string, error) {
	return colorize(color.RedString, help, ERROR)
}

func Success(help plush.HelperContext) (string, error) {
	return colorize(color.GreenString, help, SUCCESS)
}

func colorize(fn func(s string, i ...interface{}) string, help plush.HelperContext, mark string) (string, error) {
	if !help.HasBlock() {
		return "", errors.New("no block given")
	}
	x, err := help.Block()
	if err != nil {
		return "", err
	}
	x = strings.TrimSpace(x)
	if len(mark) > 0 {
		x = fmt.Sprintf("%s %s", mark, x)
	}
	x = strings.TrimSpace(fn(x))
	return x, nil
}

package display

import (
	"fmt"
	"github.com/fatih/color"
)

func Warn(tpl string, args ...interface{}) {
	k := color.New()
	k.Add(color.FgHiYellow).Print(fmt.Sprintf(tpl, args))
}

func Error(tpl string, args ...interface{}) {
	k := color.New()
	k.Add(color.FgHiRed).Print(fmt.Sprintf(tpl, args))
}

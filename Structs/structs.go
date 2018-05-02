// Package Structs contains the structures needed in code.
package Structs

import "os/exec"
import (
	"github.com/ProtonMail/ui"
)

// Result contains the result of a command executable.
type Result struct {
	Command     *exec.Cmd
	CommandFunc func(cmd *exec.Cmd)
}

// UIComp contains the struct for the necessary GUI components.
type UIComp struct {
	Label *ui.Label
	Output *ui.MultilineEntry
}

package Structs

import "os/exec"
import "github.com/ProtonMail/ui"

type Result struct {
	Command     *exec.Cmd
	CommandFunc func(cmd *exec.Cmd)
}

type UIComp struct {
	Label *ui.Label
	Output *ui.MultilineEntry
}
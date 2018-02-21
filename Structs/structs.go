package Structs

import "os/exec"

type Result struct {
	Command     *exec.Cmd
	CommandFunc func(cmd *exec.Cmd)
}
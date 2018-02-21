package Structs

import "os/exec"

type Result struct {
	Field1 *exec.Cmd
	Field2 func(cmd *exec.Cmd)
}
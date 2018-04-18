package Osutil

import (
	"Capstone/Configuration"
)

// every operating system will implement this to run tools and build UI its own way
// right now OSX and GNU/Linux aren't implementing GUI because ProtonMail/ui doesn't actually work cross platform despite claims
type ToolRunner interface {
	MakeGUI(config Configuration.Config)
}

// Package Osutil is the package that contains the various code for different operating systems.
package Osutil

import (
	"Capstone/Configuration"
)


// ToolRunner - every operating system will implement this to run tools and build UI its own way.
// Right now OSX and GNU/Linux aren't implementing GUI because ProtonMail/ui doesn't actually work cross platform despite claims.
type ToolRunner interface {
	MakeGUI(config Configuration.Config)
}

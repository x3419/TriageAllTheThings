package Osutil

import (
	"Capstone/Windows"
	"Capstone/Linux"
	"github.com/ProtonMail/ui"
	"Capstone/Configuration"
	"Capstone/Structs"
)

type Util struct {}

func (u Util) MakeGUI(config Configuration.Config) {
	// Implement GNU/Linux

	nix := config.NixTools

	// GNU/Linux tools

	if nix.Ps.Enabled {
		Linux.Ps(nix.Ps.Args)
	}

}
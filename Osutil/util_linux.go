package Osutil

import (
	"Capstone/Windows"
	"Capstone/Linux"
	"github.com/ProtonMail/ui"
	"Capstone/Configuration"
	"Capstone/Structs"
)

type Util struct {}

func (u Util) BuildUi(myBox *ui.Box, uiCompMap map[string]Structs.UIComp, config Configuration.Config, tsks chan <- Structs.Result) {

	//win := config.WinTools
	nix := config.NixTools

	// GNU/Linux tools

	if nix.Ps.Enabled {
		Linux.Ps(nix.Ps.Args)
	}

}

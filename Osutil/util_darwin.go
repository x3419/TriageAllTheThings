package Osutil

import (
	//"Capstone/Windows"
	//"Capstone/Linux"
	"Capstone/Configuration"
	"Capstone/Structs"
	"github.com/ProtonMail/ui"
)

type Util struct{}

func (u Util) BuildUi(myBox *ui.Box, uiCompMap map[string]Structs.UIComp, config Configuration.Config, tsks chan<- Structs.Result) {

	// demo code just to test if GUI can work on OSX

}

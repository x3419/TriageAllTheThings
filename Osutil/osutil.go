package Osutil

import (
	"Capstone/Structs"
	"Capstone/Configuration"
	"github.com/ProtonMail/ui"
)

type ToolRunner interface{
	BuildUi(myBox *ui.Box, uiCompMap map[string]Structs.UIComp, toolStatuses *ui.MultilineEntry, config Configuration.Config, tsks chan <- Structs.Result)
}


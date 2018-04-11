package Osutil

import (
	"Capstone/Configuration"
)

type ToolRunner interface{
	//BuildUi(myBox *ui.Box, uiCompMap map[string]Structs.UIComp, toolStatuses *ui.MultilineEntry, config Configuration.Config)
	MakeGUI(config Configuration.DynamicConfig)
}


package Osutil

import (
	"github.com/ProtonMail/ui"
	"Capstone/Configuration"
	"Capstone/Structs"

	"strings"
)

type Util struct {}

var RelativePath bool

// Defining the custom functions for tools you would like to do output parsing with (or other stuff like hide windows and whatnot)
// If function not defined for a given tool, default output will be used
var WinFunctions = map[string] func(tool Configuration.Tool, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {
	"bulkextractor":BulkExtractor,
	"tcpflow":Tcpflow,
	"winprefetch":WinPrefetch,
	"fiwalk":Fiwalk,
	"mftdump":MftDump,
}


func BuildUi(myBox *ui.Box, uiCompMap map[string]Structs.UIComp, toolStatuses *ui.MultilineEntry, config Configuration.Config) {

	for _,t := range(config.Tool){
		RelativePath = config.RelativePath
		if(t.Enabled){
			uiComp := uiCompMap[strings.ToLower(t.Name)]
			if(WinFunctions[strings.ToLower(t.Name)] != nil){
				WinFunctions[strings.ToLower(t.Name)](t, uiComp, toolStatuses)
			} else {
				Default(t, uiComp, toolStatuses)
			}
		}
	}
}


func (u Util) MakeGUI(config Configuration.Config) {

	//----------- GUI
	err := ui.Main(func() {

		myBox := ui.NewVerticalBox()
		window := ui.NewWindow("Forensic Triager", 600, 350, false)
		window.SetMargined(true)



		dropDown := ui.NewCombobox()
		myBox.Append(dropDown, false)
		myBox.Append(ui.NewLabel(""), false) // some padding


		tabs := ui.NewTab()
		statusBox := ui.NewVerticalBox()
		toolStatuses := ui.NewMultilineEntry()
		statusBox.Append(toolStatuses, true)
		tabs.Append("Tools", myBox)
		tabs.Append("Status", statusBox)
		window.SetChild(tabs)



		componentMap := make(map[string]Structs.UIComp)
		var groupList []*ui.Box

		index := 0
		tools := config.Tool
		for _, t := range tools {
			enabled := t.Enabled
			if(enabled) {

				toolName := strings.ToLower(t.Name)
				compStruct := Structs.UIComp{
					ui.NewLabel("Output:	\nStatus: Processing"),
					ui.NewMultilineNonWrappingEntry()}

				index++
				componentMap[toolName] = compStruct

				dropDown.Append(strings.Title(toolName))

				// copied
				group := ui.NewHorizontalBox()

				groupList = append(groupList, group)
				group.Hide()

				newBox := ui.NewHorizontalBox()

				labelBox := ui.NewVerticalBox()
				labelBox.SetPadded(true)
				labelBox.Append(compStruct.Label, true)

				textBox := ui.NewVerticalBox()
				textBox.Append(compStruct.Output, true)
				textBox.SetPadded(true)

				newBox.Append(labelBox, false)
				newBox.Append(textBox, true)

				group.Append(newBox, true)


				myBox.Append(group, true)

				toolStatuses.Append(strings.Title(toolName) + " - Processing\n")

			}
		}

		groupList[0].Show()
		dropDown.SetSelected(0)

		dropDown.OnSelected(func(c *ui.Combobox){

			for i:=0; i < index; i++ {
				if(c.Selected() == i){
					groupList[i].Show()
					for j:=0; j < index; j++ {
						if(i != j){
							groupList[j].Hide()
						}
					}
				}
			}

		})

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()

		go BuildUi(myBox, componentMap, toolStatuses, config)
	})

	if err != nil {
		panic(err)
	}
	//---------- GUI


}


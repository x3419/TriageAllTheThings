package Osutil

import (
	//"Capstone/Windows"
	//"Capstone/Linux"
	"github.com/ProtonMail/ui"
	"Capstone/Configuration"
	"Capstone/Structs"

	"github.com/fatih/structs"
	"strings"
)

type Util struct {}

func BuildUi(myBox *ui.Box, uiCompMap map[string]Structs.UIComp, toolStatuses *ui.MultilineEntry, config Configuration.Config) {

	win := config.WinTools
	//nix := config.NixTools

	// Windows tools
	if win.BulkExtractor.Enabled {
		//uiComp := uiCompMap["bulkextractor"]
		//BulkExtractor(win.BulkExtractor.Args, uiComp, toolStatuses)
	}
	if win.Fiwalk.Enabled {
		uiComp := uiCompMap["fiwalk"]
		Fiwalk(win.Fiwalk.Args, uiComp, toolStatuses)
	}
	if win.Blkcalc.Enabled {
		Blkcalc(win.Blkcalc.Args)
	}
	if win.Blkcat.Enabled {
		Blkcat(win.Blkcat.Args)
	}
	if win.Blkls.Enabled {
		Blkls(win.Blkls.Args)
	}
	if win.Blkstat.Enabled {
		Blkstat(win.Blkstat.Args)
	}
	if win.Fcat.Enabled {
		Fcat(win.Fcat.Args)
	}
	if win.Ffind.Enabled {
		Ffind(win.Ffind.Args)
	}
	if win.Fls.Enabled {
		Fls(win.Fls.Args)
	}
	if win.Fsstat.Enabled {
		Fsstat(win.Fsstat.Args)
	}
	if win.Hfind.Enabled {
		Hfind(win.Hfind.Args)
	}
	if win.Icat.Enabled {
		Icat(win.Icat.Args)
	}
	if win.Ifind.Enabled {
		Ifind(win.Ifind.Args)
	}
	if win.Ils.Enabled {
		Ils(win.Ils.Args)
	}
	if win.Imgcat.Enabled {
		Img_cat(win.Imgcat.Args)
	}
	if win.Imgstat.Enabled {
		Img_stat(win.Imgstat.Args)
	}
	if win.Istat.Enabled {
		Istat(win.Istat.Args)
	}
	if win.Jcat.Enabled {
		Jcat(win.Jcat.Args)
	}
	if win.Jls.Enabled {
		Img_cat(win.Jls.Args)
	}
	if win.Mmcat.Enabled {
		Mmcat(win.Mmcat.Args)
	}
	if win.Mmls.Enabled {
		Mmls(win.Mmls.Args)
	}
	if win.Mmstat.Enabled {
		Mmstat(win.Mmstat.Args)
	}
	if win.TskCompareDir.Enabled {
		Tsk_comparedir(win.TskCompareDir.Args)
	}
	if win.TskGetTimes.Enabled {
		Tsk_gettimes(win.TskCompareDir.Args)
	}
	if win.TskLoaddb.Enabled {
		Tsk_loaddb(win.TskCompareDir.Args)
	}
	if win.TskRecover.Enabled {
		Tsk_recover(win.TskRecover.Args)
	}
	if win.Tcpflow.Enabled {
		uiComp := uiCompMap["tcpflow"]
		Tcpflow(win.Tcpflow.Args, uiComp, toolStatuses)
	}
	if win.WinPrefetch.Enabled {

		uiComp := uiCompMap["winprefetch"]
		WinPrefetch(win.WinPrefetch.Args, uiComp, toolStatuses)

	}
	if win.MFTDump.Enabled {

		uiComp := uiCompMap["mftdump"]
		MftDump(win.MFTDump.Args, uiComp, toolStatuses)
	}

	// Copying the windows file

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
		t := structs.New(config.WinTools)
		tools := t.Fields()
		for _, t := range tools {
			enabled := t.Value().(Configuration.Tool).Enabled
			if(enabled) {

				toolName := strings.ToLower(t.Name())
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

		//var os Osutil.ToolRunner = Osutil.Util{}
		go BuildUi(myBox, componentMap, toolStatuses, config)
	})
	if err != nil {
		panic(err)
	}
	//---------- GUI


}


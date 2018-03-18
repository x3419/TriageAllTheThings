package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"encoding/json"
	"runtime"
	Configuration "Capstone/Configuration"
	"Capstone/Structs"
	"strings"
	"sync"
	"github.com/ProtonMail/ui"
	"github.com/fatih/structs"

	"Capstone/Osutil"
)



func main() {

	configPathPtr := flag.String("config", "Configuration/config.txt", "Location of the configuration file")
	flag.Parse()
	var config Configuration.Config = ParseConfig(*configPathPtr)

	fmt.Println(strings.Title(runtime.GOOS) + " OS detected\nEnabled tools will begin to run in parallel. This may take some time and will slow the system down, so please be patient.")

	tasks := make(chan Structs.Result, 64)

	// spawn ten worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < 10 ; i++ {
		wg.Add(1)
		go func() {
			for results := range tasks {
				results.CommandFunc(results.Command)
			}
			wg.Done()
		}()
	}

	makeGUI(config, tasks)
	close(tasks)

	// wait for the workers to finish
	wg.Wait()

}


func ParseConfig(configFile string) Configuration.Config {

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("Unable to open config file: ", err)
	}

	myConfig := Configuration.Config{}

	if err := json.Unmarshal(b, &myConfig); err != nil {
		fmt.Println("Error! Problem parsing the configuration file - please ensure that it reflects the example on Github.\n", err)
		panic(err)
		return Configuration.Config{}
	}

	return myConfig

}

func addToolToUI(myBox *ui.Box, tool string, label *ui.Label, output *ui.MultilineEntry) {



	ui.QueueMain(func() {

		group := ui.NewGroup(tool)
		newBox := ui.NewHorizontalBox()

		labelBox := ui.NewVerticalBox()
		labelBox.SetPadded(true)
		labelBox.Append(label, true)

		textBox := ui.NewVerticalBox()
		textBox.Append(output, true)
		textBox.SetPadded(true)

		newBox.Append(labelBox, false)
		newBox.Append(textBox, true)

		group.SetChild(newBox)
		group.SetMargined(true)

		myBox.Append(group, true)
		myBox.Append(ui.NewLabel(""), false)


	})

}

func makeGUI(config Configuration.Config, tsks chan <- Structs.Result) {

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
				//group.SetChild(newBox)
				//group.SetMargined(true)

				myBox.Append(group, true)

				toolStatuses.Append(strings.Title(toolName) + " - Processing\n")


				//addToolToUI(myBox, strings.Title(strings.ToLower(t.Name())), compStruct.Label, compStruct.Output)
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

		var os Osutil.ToolRunner = Osutil.Util{}
		go os.BuildUi(myBox, componentMap, toolStatuses, config, tsks)
	})
	if err != nil {
		panic(err)
	}
	//---------- GUI


}


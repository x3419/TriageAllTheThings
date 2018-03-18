package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"encoding/json"
	"runtime"
	Configuration "Capstone/Configuration"
	//Windows "Capstone/Windows"
	//Linux "Capstone/Linux"
	"Capstone/Structs"
	"strings"
	"sync"
	//"Capstone/Osutil"
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

	// spawn four worker goroutines
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

	windowsTools(config, tasks)
	//
	//if runtime.GOOS == "windows" {
	//	windowsTools(config, tasks)
	//} else if runtime.GOOS == "linux" {
	//	fmt.Println("GNU/Linux compatibility coming soon!")
	//} else if runtime.GOOS == "darwin" {
	//	fmt.Println("OSX compatibility coming soon!")
	//} else {
	//	fmt.Println(strings.Title(runtime.GOOS) + " OS is not supported in this project.")
	//}




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



		//newLable := ui.NewLabel("Last output:			\nStatus: Processing")
		//newText := ui.NewMultilineNonWrappingEntry()

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

func windowsTools(config Configuration.Config, tsks chan <- Structs.Result) {

	//----------- GUI
	err := ui.Main(func() {


		myBox := ui.NewVerticalBox()
		window := ui.NewWindow("Forensic Triager", 600, 350, false)
		window.SetMargined(true)

		window.SetChild(myBox)


		componentMap := make(map[string]Structs.UIComp)

		t := structs.New(config.WinTools)
		tools := t.Fields()
		for _, t := range tools {
			enabled := t.Value().(Configuration.Tool).Enabled
			if(enabled) {
				//myTool := t.Value().(Configuration.Tool)
				toolName := strings.ToLower(t.Name())
				compStruct := Structs.UIComp{
					ui.NewLabel("Last output:	\nStatus: Processing"),
					ui.NewMultilineNonWrappingEntry()}

				componentMap[toolName] = compStruct
				
				addToolToUI(myBox, strings.Title(strings.ToLower(t.Name())), compStruct.Label, compStruct.Output)
			}
		}

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()

		var os Osutil.ToolRunner = Osutil.Util{}
		go os.BuildUi(myBox, componentMap, config, tsks)
	})
	if err != nil {
		panic(err)
	}
	//---------- GUI


}


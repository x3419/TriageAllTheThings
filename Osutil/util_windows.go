package Osutil

import (
	"Capstone/Configuration"
	"Capstone/Structs"
	"bufio"
	"fmt"
	"github.com/ProtonMail/ui"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
	"time"
)

type Util struct{}

var RelativePath bool

// Defining the custom functions for tools you would like to do output parsing with (or other stuff like hide windows and whatnot)
// If function not defined for a given tool, default output will be used
var WinFunctions = map[string]func(tool Configuration.Tool, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry){
	"bulkextractor": BulkExtractor,
	"tcpflow":       Tcpflow,
	"winprefetch":   WinPrefetch,
	"fiwalk":        Fiwalk,
	"mftdump":       MftDump,
}

func BuildUi(myBox *ui.Box, uiCompMap map[string]Structs.UIComp, toolStatuses *ui.MultilineEntry, config Configuration.Config) {

	for _, t := range config.Tool {
		RelativePath = config.RelativePath
		if t.Enabled {
			uiComp := uiCompMap[strings.ToLower(t.Name)]
			if WinFunctions[strings.ToLower(t.Name)] != nil {
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
			if enabled {

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

		dropDown.OnSelected(func(c *ui.Combobox) {

			for i := 0; i < index; i++ {
				if c.Selected() == i {
					groupList[i].Show()
					for j := 0; j < index; j++ {
						if i != j {
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

// -- Below is all the code that was in WinHelp.go. They must now be in this file to ensure
// -- cross platform compatibility works corectly
const TCPFLOW_TIME_LIMIT = time.Hour * 2

func BulkExtractorParse(cmd *exec.Cmd, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	go func() {
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()

		first := time.Now()
		uiComp.Output.Append("BulkExtractor: Processing...\n")
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()
			uiComp.Output.Append(m + "\r\n")
			if strings.Contains(m, "(") && strings.Contains(m, "%)") {

				now := time.Now()
				if now.Sub(first) > time.Second*5 {
					uiComp.Output.Append("BulkExtractor: " + m[strings.Index(m, "(")+1:strings.Index(m, ")")] + " Complete\r\n")
					first = time.Now()
				}

			}
		}

		uiComp.Label.SetText(strings.Replace(uiComp.Label.Text(), "Processing", "Complete", 1))
		toolStatuses.SetText(strings.Replace(toolStatuses.Text(), "Bulkextractor - Processing", "Bulkextractor - Complete", 1))
		toolStatuses.SetText(strings.Replace(toolStatuses.Text(), "\n", "\r\n", -1))

		cmd.Wait()

	}()

}

func MftDumpParse(cmd *exec.Cmd, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	go func() {
		uiComp.Output.Append("MftDump: Beggining dumping and parsing the Master File Table")
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()

			uiComp.Output.Append(m + "\n")
			if strings.Contains(m, "Executable name") {

				uiComp.Output.Append("Processing executable " + m[strings.Index(m, "Executable:")+16:len(m)])
				//time.Sleep(time.Second * 5)
			}
		}

		uiComp.Label.SetText(strings.Replace(uiComp.Label.Text(), "Processing", "Complete", 1))
		toolStatuses.SetText(strings.Replace(toolStatuses.Text(), "Mftdump - Processing", "Mftdump - Complete", 1))
		toolStatuses.SetText(strings.Replace(toolStatuses.Text(), "\n", "\r\n", -1))

		cmd.Wait()
	}()

}

func WinPrefetchParse(cmd *exec.Cmd, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	go func() {

		stdout, _ := cmd.StdoutPipe()
		cmd.Start()

		uiComp.Output.Append("WinPrefetch: Processing...\n")
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()

			if strings.Contains(m, "Executable name") {

				uiComp.Output.Append("Processing executable " + m[strings.Index(m, "Executable:")+16:len(m)] + "\n")
				//time.Sleep(time.Second * 5)
			}
		}

		uiComp.Label.SetText(strings.Replace(uiComp.Label.Text(), "Processing", "Complete", 1))
		toolStatuses.SetText(strings.Replace(toolStatuses.Text(), "Winprefetch - Processing", "Winprefetch - Complete", 1))
		toolStatuses.SetText(strings.Replace(toolStatuses.Text(), "\n", "\r\n", -1))

		cmd.Wait()
	}()

}

func TcpFlowParse(cmd *exec.Cmd, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	go func() {
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()

		totalTime := time.Now()
		uiComp.Output.Append("TcpFlow: Processing...\n")
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()
			if strings.Contains(m, "Packets     :") {

				currPacketTotal := m[strings.Index(m, "Packets     :")+14 : len(m)]
				uiComp.Output.Append("Tcpflow: logged " + currPacketTotal + " packets\n")

			} else {
				uiComp.Output.Append(m + "\n")
			}

			if time.Now().Sub(totalTime) > TCPFLOW_TIME_LIMIT {
				uiComp.Output.Append("TcpFlow limit exceeded. Pcap file written.\n")
			}
		}
		uiComp.Label.SetText(strings.Replace(uiComp.Label.Text(), "Processing", "Complete", -1))
		toolStatuses.SetText(strings.Replace(toolStatuses.Text(), "Tcpflow - Processing", "Tcpflow - Complete", 1))
		toolStatuses.SetText(strings.Replace(toolStatuses.Text(), "\n", "\r\n", -1))

		cmd.Wait()
	}()

}

func FiwalkParse(cmd *exec.Cmd, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	go func() {
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()

		uiComp.Output.Append("Fiwalk: Processing...\n")
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()
			uiComp.Output.Append(m + "\n")
		}
		uiComp.Label.SetText(strings.Replace(uiComp.Label.Text(), "Processing", "Complete", -1))
		toolStatuses.SetText(strings.Replace(toolStatuses.Text(), "Fiwalk - Processing", "Fiwalk - Complete", 1))
		toolStatuses.SetText(strings.Replace(toolStatuses.Text(), "\n", "\r\n", -1))

		cmd.Wait()
	}()

}

func DefaultParse(name string, cmd *exec.Cmd, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	go func() {
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()
		uiComp.Output.Append(name + ": Processing...\n")
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()
			uiComp.Output.Append(m + "\n")
		}
		uiComp.Label.SetText(strings.Replace(uiComp.Label.Text(), "Processing", "Complete", -1))
		toolStatuses.SetText(strings.Replace(toolStatuses.Text(), strings.Title(strings.ToLower(name))+" - Processing", strings.Title(strings.ToLower(name))+" - Complete", 1))
		toolStatuses.SetText(strings.Replace(toolStatuses.Text(), "\n", "\r\n", -1))

		cmd.Wait()
	}()

}

func WriteCmdResultToDisk(filename string) func(cmd *exec.Cmd) {
	return func(cmd *exec.Cmd) {

		stdout, _ := cmd.StdoutPipe()
		cmd.Start()

		fullOutput := ""

		first := time.Now()

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()
			now := time.Now()
			if now.Sub(first) > time.Second*5 {
				fmt.Println(filename[0:strings.Index(filename, ".")] + " looks like it's making progress...")
				first = time.Now()
			}

			fullOutput += m + "\n"
		}

		cmd.Wait()

		fmt.Println(strings.Title(filename[0:strings.Index(filename, ".")]) + ": Complete")

		d1 := []byte(fullOutput)
		err := ioutil.WriteFile("./Output/"+filename, d1, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func BulkExtractor(tool Configuration.Tool, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	//Commenting out for dev speed related reasons
	//cmd :=  cmdTool(args, tool.Path)
	//BulkExtractorParse(cmd, uiComp, toolStatuses)
	fmt.Println("Running BulkExtractor!")
}

func Default(tool Configuration.Tool, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	cmd := cmdTool(tool.Args, tool.Path)
	DefaultParse(tool.Name, cmd, uiComp, toolStatuses)
}

func Fiwalk(tool Configuration.Tool, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {
	cmd := cmdTool(tool.Args, tool.Path)
	FiwalkParse(cmd, uiComp, toolStatuses)
}

func Tcpflow(tool Configuration.Tool, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	cmd := cmdTool(tool.Args, tool.Path)
	cmd = makeCmdQuiet(cmd)
	TcpFlowParse(cmd, uiComp, toolStatuses)
}

func WinPrefetch(tool Configuration.Tool, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {
	cmd := cmdTool(tool.Args, tool.Path)
	WinPrefetchParse(cmd, uiComp, toolStatuses)
}

func MftDump(tool Configuration.Tool, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	cmd := cmdTool(tool.Args, tool.Path)
	MftDumpParse(cmd, uiComp, toolStatuses)
}

func makeCmdQuiet(cmd *exec.Cmd) *exec.Cmd {

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}

func cmdTool(args string, tool string) *exec.Cmd {

	var myArgs []string
	if RelativePath {
		myArgs = []string{"/C", "Tools\\" + tool}
	} else {
		myArgs = []string{"/C", tool}
	}

	r := regexp.MustCompile("[^\\s]+")
	myArgs2 := r.FindAllString(args, -1)
	myArgs = append(myArgs, myArgs2...)

	cmd := exec.Command("cmd", myArgs...)

	return cmd
}

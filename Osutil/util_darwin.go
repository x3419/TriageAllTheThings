package Osutil

import (
	"Capstone/Configuration"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"sync"
	"time"
)

// -- This is the "debug" version of the UI for Mac OSX. The UI code from util_windows.go will
// -- work on certain versions of OSX (10.11 supposedly) but seems to act buggy in others.
// -- For this reason, I have decided to disable the UI for OSX. Each tool will have an output file
// -- showing the output of each tool. In the future I will implement separate UI frameworks for
// -- each operating system to ensure stable compatibility.
// -- NOTE: If you are interested in trying to get the ProtonMail/ui framework working on OSX, just copy
// -- the code from util_windows.go. The instability made me feel iffy about including it in my final Capstone

type Util struct{}

var RelativePath bool
var toolsProcessing = make([]string, 100, 100)
var toolsComplete = make([]string, 100, 100)
var timeStart time.Time
var comm bool = false

func (u Util) MakeGUI(config Configuration.Config) {

	fmt.Println()
	timeStart = time.Now()
	var wg sync.WaitGroup

	for _, t := range config.Tool {
		RelativePath = config.RelativePath
		if t.Enabled {
			defaultFunc(t, &wg)
		}
	}

	wg.Wait()

	fmt.Println("Triage Tool Done!")

}

func defaultFunc(tool Configuration.Tool, wg *sync.WaitGroup) error {

	cmd := cmdTool(tool.Args, tool.Path)
	err := defaultParse(tool.Name, cmd, wg)
	return err
}

func cmdTool(args string, tool string) *exec.Cmd {

	var myArgs string
	if RelativePath {
		myArgs = "Tools\\" + tool
	} else {
		myArgs = tool
	}

	r := regexp.MustCompile("[^\\s]+")
	myArgs2 := r.FindAllString(args, -1)

	cmd := exec.Command(myArgs, myArgs2...)

	return cmd
}

func defaultParse(name string, cmd *exec.Cmd, wg *sync.WaitGroup) error {

	toolsProcessing = append([]string{name}, toolsProcessing...)
	var err error
	wg.Add(1)
	go func() {

		stdout, _ := cmd.StdoutPipe()
		cmd.Start()

		scanner := bufio.NewScanner(stdout)
		fullOutput := ""

		for scanner.Scan() {
			m := scanner.Text()

			fullOutput += m + "\n"
			timeNext := time.Now()
			toolsCurrentProgress := ""
			if !comm {
				comm = true
				if timeNext.Sub(timeStart) > time.Second*5 {

					toolsCurrentProgress += "Tools still processing:"
					for _, element := range toolsProcessing {
						if element != "" {
							toolsCurrentProgress += "\n\t" + element
						}
					}

					toolsCurrentProgress += "\n\nTools complete:"
					for _, element := range toolsComplete {
						if element != "" {
							toolsCurrentProgress += "\n\t" + element
						}
					}

					fmt.Println(toolsCurrentProgress + "\n")
					timeStart = time.Now()
				}
				comm = false
			}

		}

		cmd.Wait()

		d1 := []byte(fullOutput)
		CreateDirIfNotExist("./Output/")
		err = ioutil.WriteFile("./Output/"+name+".txt", d1, 0777)

		for index, element := range toolsProcessing {
			if element == name {
				toolsProcessing = append(toolsProcessing[:index], toolsProcessing[index+1:]...)
			}
		}

		toolsComplete = append([]string{name}, toolsComplete...)

		wg.Done()

	}()
	return err

}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			panic(err)
		}
	}
}

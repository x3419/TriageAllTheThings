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

// -- This is the "debug" version of the UI for GNU/Linux. The UI code from util_windows.go will
// -- not for for Linux for some reason.
// -- For this reason, I have decided to disable the UI for GNU/Linux. Each tool will have an output file
// -- showing the output of each tool. In the future I will implement separate UI frameworks for
// -- each operating system to ensure stable compatibility.
// -- NOTE: If you are interested in trying to get the ProtonMail/ui framework working on GNu/Linux, just copy
// -- the code from util_windows.go. The instability made me feel iffy about including it in my final Capstone

// Util specifies that this is an implementation of the interface ToolRunner.
type Util struct{}

// RelativePath determines whether the relative path is set to true within the configuration file.
var RelativePath bool
var toolsProcessing = make([]string, 100, 100)
var toolsComplete = make([]string, 100, 100)
var timeStart time.Time

// comm is a goroutine currently outputing progress data.
var comm bool = false

// MakeGUI doesn't really make GUI. Instead it runs the tools and eventually outputs tool statuses.
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


// defaultFunc writes the executable output to disk without any parsing.
func defaultFunc(tool Configuration.Tool, wg *sync.WaitGroup) error {

	cmd := cmdTool(tool.Args, tool.Path)
	err := defaultParse(tool.Name, cmd, wg)
	return err
}

// cmdTool takes arguments and an executable string and returns the executable Cmd.
func cmdTool(args string, tool string) *exec.Cmd {

	// check where the executable is located
	var myArgs string
	if RelativePath {
		myArgs = "Tools\\" + tool
	} else {
		myArgs = tool
	}

	// conv from "a b c" to ["a", "b", "c"]
	r := regexp.MustCompile("[^\\s]+")
	myArgs2 := r.FindAllString(args, -1)

	cmd := exec.Command(myArgs, myArgs2...)

	return cmd
}

// defaultParse outputs tool statuses and writes output to file.
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
			
			// if another goroutine isn't printing statuses
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

		// write output
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

// CreateDirIfNotExists checks whether a directory exists. If it does not, it creates a folder with the name specified by the "dir" parameter.
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			panic(err)
		}
	}
}

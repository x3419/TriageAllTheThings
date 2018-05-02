package Osutil

import (
	"Capstone/Configuration"
	"bufio"
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"sync"
	"time"
)

// -- This is the "debug" version of the UI for GNU/Linux. The UI code from util_windows.go can
// -- hypothetically work on GNU/Linux but seems to show some type of compilation error.
// -- For this reason, I have decided to disable the UI for GNU/Linux. Each tool will have an output file
// -- showing the output of each tool. In the future I will implement separate UI frameworks for
// -- each operating system to ensure stable compatibility.
// -- NOTE: If you are interested in trying to get the ProtonMail/ui framework working on GNU/Linux, just copy
// -- the code from util_windows.go. The instability made me feel iffy about including it in my final Capstone

type Util struct{}

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
func defaultFunc(tool Configuration.Tool, wg *sync.WaitGroup) {

	cmd := cmdTool(tool.Args, tool.Path)
	defaultParse(tool.Name, cmd, wg)
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
func defaultParse(name string, cmd *exec.Cmd, wg *sync.WaitGroup) {

	toolsProcessing = append([]string{name}, toolsProcessing...)

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

		d1 := []byte(fullOutput)
		err := ioutil.WriteFile("./Output/"+name+".txt", d1, 0644)
		if err != nil {
			panic(err)
		}

		for index, element := range toolsProcessing {
			if element == name {
				toolsProcessing = append(toolsProcessing[:index], toolsProcessing[index+1:]...)
			}
		}

		toolsComplete = append([]string{name}, toolsComplete...)

		wg.Done()
	}()

}

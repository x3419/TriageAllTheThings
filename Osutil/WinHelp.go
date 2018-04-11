package Osutil

import (
	"os/exec"
	"fmt"
	"regexp"
	"bufio"
	"strings"
	"Capstone/Structs"
	"io/ioutil"
	"time"
	"syscall"
	"github.com/ProtonMail/ui"
	"Capstone/Configuration"
)

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
			if ( strings.Contains(m, "(") && strings.Contains(m, "%)") ) {


				now := time.Now()
				if(now.Sub(first) > time.Second * 5) {
					uiComp.Output.Append("BulkExtractor: " + m[strings.Index(m, "(") + 1:strings.Index(m,")")] + " Complete\r\n")
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
			if (strings.Contains(m, "Executable name")) {

				uiComp.Output.Append("Processing executable " + m[strings.Index(m, "Executable:") + 16:len(m)])
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

			if (strings.Contains(m, "Executable name")) {

				uiComp.Output.Append("Processing executable " + m[strings.Index(m, "Executable:") + 16:len(m)] + "\n")
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

				currPacketTotal := m[strings.Index(m, "Packets     :")+14: len(m)]
				uiComp.Output.Append("Tcpflow: logged " + currPacketTotal + " packets\n")

			} else {
				uiComp.Output.Append(m + "\n")
			}

			if (time.Now().Sub(totalTime) > TCPFLOW_TIME_LIMIT) {
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
		toolStatuses.SetText(strings.Replace(toolStatuses.Text(), strings.Title(strings.ToLower(name)) + " - Processing", strings.Title(strings.ToLower(name)) + " - Complete", 1))
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
			if(now.Sub(first) > time.Second * 5) {
				fmt.Println(filename[0:strings.Index(filename, ".")] + " looks like it's making progress...")
				first = time.Now()
			}

			fullOutput += m + "\n"
		}

		cmd.Wait()

		fmt.Println(strings.Title(filename[0:strings.Index(filename, ".")]) + ": Complete")


		d1 := []byte(fullOutput)
		err := ioutil.WriteFile("./Output/" + filename, d1, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func BulkExtractor(tool Configuration.DynamicTool, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	//Commenting out for dev speed related reasons
	//cmd :=  cmdTool(args, "bulk_extractor32.exe")
	//BulkExtractorParse(cmd, uiComp, toolStatuses)
	fmt.Println("Running BulkExtractor!")
}

func Default(tool Configuration.DynamicTool, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	//Commenting out for dev speed related reasons
	cmd :=  cmdTool(tool.Args, tool.Location)
	DefaultParse(tool.Name, cmd, uiComp, toolStatuses)
}

func Fiwalk(tool Configuration.DynamicTool, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {
	cmd :=  cmdTool(tool.Args, "fiwalk-0.6.3.exe")
	FiwalkParse(cmd, uiComp, toolStatuses)
}

func Tcpflow(tool Configuration.DynamicTool, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	//Commenting out for dev speed related reasons

	cmd :=  cmdTool(tool.Args,"RawCap.exe")
	cmd = makeCmdQuiet(cmd)
	TcpFlowParse(cmd, uiComp, toolStatuses)
}


func WinPrefetch(tool Configuration.DynamicTool, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {
	cmd :=  cmdTool(tool.Args, "PECmd.exe")
	WinPrefetchParse(cmd, uiComp, toolStatuses)
}

func MftDump(tool Configuration.DynamicTool, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	//Commenting out for dev speed related reasons

	cmd :=  cmdTool(tool.Args, "mftdump.exe")
	MftDumpParse(cmd, uiComp, toolStatuses)
}

func makeCmdQuiet(cmd *exec.Cmd) *exec.Cmd{

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}


func cmdTool(args string, tool string) *exec.Cmd {
	myArgs := []string{"/C", "Tools\\" + tool}
	r := regexp.MustCompile("[^\\s]+")
	myArgs2 := r.FindAllString(args, -1)
	myArgs = append(myArgs, myArgs2...)

	cmd := exec.Command("cmd", myArgs...)

	return cmd
}


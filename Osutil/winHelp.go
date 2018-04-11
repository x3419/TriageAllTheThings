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
	//"Capstone/Configuration"
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

func BulkExtractor(args string, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	//Commenting out for dev speed related reasons
	//cmd :=  cmdTool(args, "bulk_extractor32.exe")
	//BulkExtractorParse(cmd, uiComp, toolStatuses)
}

func Fiwalk(args string, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {
	cmd :=  cmdTool(args, "fiwalk-0.6.3.exe")
	FiwalkParse(cmd, uiComp, toolStatuses)
}

func Blkcalc(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\blkcalc.exe")
	runDefault(cmd)
}

func Blkcat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\blkcat.exe")
	runDefault(cmd)
}

func Blkls(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\blkls.exe")
	runDefault(cmd)
}

func Blkstat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\blkstat.exe")
	runDefault(cmd)
}

func Fcat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\fcat.exe")
	runDefault(cmd)
}

func Ffind(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\ffind.exe")
	runDefault(cmd)
}

func Fls(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\fls.exe")
	runDefault(cmd)
}

func Fsstat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\fsstat.exe")
	runDefault(cmd)
}

func Hfind(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\hfind.exe")
	runDefault(cmd)
}

func Icat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\icat.exe")
	runDefault(cmd)
}

func Ifind(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\ifind.exe")
	runDefault(cmd)
}

func Ils(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\ils.exe")
	runDefault(cmd)
}

func Img_cat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\img_cat.exe")
	runDefault(cmd)
}

func Img_stat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\img_stat.exe")
	runDefault(cmd)
}

func Istat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\istat.exe")
	runDefault(cmd)
}

func Jcat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\jcat.exe")
	runDefault(cmd)
}

func Jls(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\jls.exe")
	runDefault(cmd)
}

func Mmcat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\mmcat.exe")
	runDefault(cmd)
}

func Mmls(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\mmls.exe")
	runDefault(cmd)
}

func Mmstat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\mmstat.exe")
	runDefault(cmd)
}


func Tsk_comparedir(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\tsk_comparedir.exe")
	runDefault(cmd)
}

func Tsk_gettimes(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\tsk_gettimes.exe")
	runDefault(cmd)
}


func Tsk_loaddb(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\tsk_loaddb.exe")
	runDefault(cmd)
}


func Tsk_recover(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\tsk_recover.exe")
	runDefault(cmd)
}

func Tcpflow(args string, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	//Commenting out for dev speed related reasons
	//
	//cmd :=  cmdTool(args, "RawCap.exe")
	//cmd = makeCmdQuiet(cmd)
	//TcpFlowParse(cmd, uiComp, toolStatuses)
}


func WinPrefetch( args string, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {
	cmd :=  cmdTool(args, "PECmd.exe")
	WinPrefetchParse(cmd, uiComp, toolStatuses)
}

func MftDump(args string, uiComp Structs.UIComp, toolStatuses *ui.MultilineEntry) {

	//Commenting out for dev speed related reasons
	//
	//cmd :=  cmdTool(args, "mftdump.exe")
	//MftDumpParse(cmd, uiComp, toolStatuses)
}


func mrutools(args string) {
	cmd :=  cmdTool(args, "mrutools.exe")
	runDefault(cmd)
}


func makeCmdQuiet(cmd *exec.Cmd) *exec.Cmd{

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}

func runDefault(cmd *exec.Cmd) {
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	cmd.Wait()
}


func cmdTool(args string, tool string) *exec.Cmd {
	myArgs := []string{"/C", "Tools\\" + tool}
	r := regexp.MustCompile("[^\\s]+")
	myArgs2 := r.FindAllString(args, -1)
	myArgs = append(myArgs, myArgs2...)

	cmd := exec.Command("cmd", myArgs...)

	return cmd
}


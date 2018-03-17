package Windows

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
)

const TCPFLOW_TIME_LIMIT = time.Hour * 2


func BulkExtractorParse(cmd *exec.Cmd) {
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	first := time.Now()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()

		if ( strings.Contains(m, "(") && strings.Contains(m, "%)") ) {


			now := time.Now()
			if(now.Sub(first) > time.Second * 5) {
				fmt.Println("BulkExtractor: " + m[strings.Index(m, "(") + 1:strings.Index(m,")")] + " Complete")
				first = time.Now()
			}

		}
	}

	fmt.Println("BulkExtractor: Complete")

	cmd.Wait()
}

func MftDumpParse(cmd *exec.Cmd, label *ui.Label, output *ui.MultilineEntry) {
	ui.QueueMain(func(){

		output.Append("MftDump: Beggining dumping and parsing the Master File Table")
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()

			if (strings.Contains(m, "Executable name")) {

				output.Append("Processing executable " + m[strings.Index(m, "Executable:") + 16:len(m)])
				time.Sleep(time.Second * 5)
			}
		}

		label.SetText(strings.Replace(label.Text(), "Processing", "Complete", -1))
		cmd.Wait()

	})

}

func WinPrefetchParse(cmd *exec.Cmd) {
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()

		if (strings.Contains(m, "Executable name")) {

			fmt.Println("WinPrefetch: processing executable " + m[strings.Index(m, "Executable:") + 16:len(m)])
			time.Sleep(time.Second * 5)
		}
	}

	fmt.Println("WinPrefetch: Complete")

	cmd.Wait()
}

func TcpFlowParse(cmd *exec.Cmd) {
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	first := time.Now()
	totalTime := time.Now()

	scanner := bufio.NewScanner(stdout)
	lastPacketTotal := "0"
	for scanner.Scan() {
		m := scanner.Text()
		if strings.Contains(m,"Packets     :") {
			currPacketTotal := m[strings.Index(m, "Packets     :")+14: len(m)]

			if ( currPacketTotal != lastPacketTotal ) {

				now := time.Now()
				if(now.Sub(first) > time.Second * 10) {
					fmt.Println("Tcpflow: logged " + currPacketTotal + " packets")
					first = time.Now()
				}
			}
		}

		if (time.Now().Sub(totalTime) > TCPFLOW_TIME_LIMIT) {
			fmt.Println("TcpFlow limit exceeded. Pcap file written.")
		}
	}

	cmd.Wait()
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

func BulkExtractor(args string, tsks chan <- Structs.Result) {

	cmd :=  cmdTool(args, "bulk_extractor32.exe")
	tsks <- Structs.Result{cmd, BulkExtractorParse}
}

func Fiwalk(args string, tsks chan <- Structs.Result) {
	cmd :=  cmdTool(args, "fiwalk-0.6.3.exe")
	tsks <- Structs.Result{cmd, WriteCmdResultToDisk("fiwalk.txt")}
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

func Tcpflow(args string) {
	cmd :=  cmdTool(args, "RawCap.exe")
	//runDefault(cmd)

	// Ideally we would want to use this:
	cmd = makeCmdQuiet(cmd)
	TcpFlowParse(cmd)
}


func WinPrefetch(args string) {
	cmd :=  cmdTool(args, "PECmd.exe")
	WinPrefetchParse(cmd)
}

func MftDump(args string, label *ui.Label, output *ui.MultilineEntry) {
	cmd :=  cmdTool(args, "mftdump.exe")
	MftDumpParse(cmd, label, output)
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


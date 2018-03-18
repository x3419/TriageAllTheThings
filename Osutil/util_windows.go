package Osutil

import (
	"Capstone/Windows"
	//"Capstone/Linux"
	"github.com/ProtonMail/ui"
	"Capstone/Configuration"
	"Capstone/Structs"
)

type Util struct {}

func (u Util) BuildUi(myBox *ui.Box, uiCompMap map[string]Structs.UIComp, config Configuration.Config, tsks chan <- Structs.Result) {

	win := config.WinTools
	//nix := config.NixTools

	// Windows tools
	if win.BulkExtractor.Enabled {
		BulkExtractor(win.BulkExtractor.Args, tsks)
	}
	if win.Fiwalk.Enabled {
		Fiwalk(win.Fiwalk.Args, tsks)
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
		uiComp := uiCompMap["tcpFlow"]
		//addToolToUI(myBox, "tcpflow", uiComp.Label, uiComp.Output)
		Tcpflow(win.Tcpflow.Args, uiComp.Label, uiComp.Output)
	}
	if win.WinPrefetch.Enabled {

		uiComp := uiCompMap["winprefetch"]
		//addToolToUI(myBox, "WinPrefetch", uiComp.Label, uiComp.Output)
		WinPrefetch(win.WinPrefetch.Args, uiComp.Label, uiComp.Output)

	}
	if win.MFTDump.Enabled {

		uiComp := uiCompMap["mftdump"]
		//addToolToUI(myBox, "MFTDump", uiComp.Label, uiComp.Output)
		MftDump(win.MFTDump.Args, uiComp.Label, uiComp.Output)
	}

	// Copying the windows file


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

		go func() {
			output.Append("MftDump: Beggining dumping and parsing the Master File Table")
			stdout, _ := cmd.StdoutPipe()
			cmd.Start()

			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				m := scanner.Text()

				output.Append(m + "\n")
				if (strings.Contains(m, "Executable name")) {

					output.Append("Processing executable " + m[strings.Index(m, "Executable:") + 16:len(m)])
					//time.Sleep(time.Second * 5)
				}
			}

			label.SetText(strings.Replace(label.Text(), "Processing", "Complete", -1))
			cmd.Wait()
		}()



	})

}

func WinPrefetchParse(cmd *exec.Cmd, label *ui.Label, output *ui.MultilineEntry) {

	go func() {

		stdout, _ := cmd.StdoutPipe()
		cmd.Start()

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()

			if (strings.Contains(m, "Executable name")) {

				output.Append("WinPrefetch: processing executable " + m[strings.Index(m, "Executable:") + 16:len(m)] + "\n")
				//time.Sleep(time.Second * 5)
			}
		}

		label.SetText(strings.Replace(label.Text(), "Processing", "Complete", -1))

		cmd.Wait()
	}()

}

func TcpFlowParse(cmd *exec.Cmd, label *ui.Label, output *ui.MultilineEntry) {

	go func() {
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
						output.Append("Tcpflow: logged " + currPacketTotal + " packets\n")
						first = time.Now()
					}
				}
			}

			if (time.Now().Sub(totalTime) > TCPFLOW_TIME_LIMIT) {
				output.Append("TcpFlow limit exceeded. Pcap file written.\n")
				label.SetText(strings.Replace(label.Text(), "Processing", "Complete", -1))
			}
		}

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

func Tcpflow(args string, label *ui.Label, output *ui.MultilineEntry) {
	cmd :=  cmdTool(args, "RawCap.exe")
	//runDefault(cmd)

	// Ideally we would want to use this:
	cmd = makeCmdQuiet(cmd)
	TcpFlowParse(cmd, label, output)
}


func WinPrefetch( args string, label *ui.Label, output *ui.MultilineEntry) {
	cmd :=  cmdTool(args, "PECmd.exe")
	WinPrefetchParse(cmd, label, output)
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



}

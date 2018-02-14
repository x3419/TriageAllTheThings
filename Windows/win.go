package Windows

import (
	"os/exec"
	"fmt"
	"regexp"
	"bufio"
	"strings"
)




func BulkExtractor(args string) {

	cmd :=  cmdTool(args, "bulk_extractor32.exe")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		if strings.Contains(m, "(") {
			fmt.Println("BulkExtractor: " + m[strings.Index(m, "(") + 1:strings.Index(m,")")] + " Complete")
		}
	}

	cmd.Wait()


}

func Fiwalk(args string) {
	cmd :=  cmdTool(args, "fiwalk-0.6.3.exe")
	runDefault(cmd)
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

func tcpflow(args string) {
	cmd :=  cmdTool(args, "tcpflow.exe")
	runDefault(cmd)
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
	myArgs := []string{"/C", "Windows\\" + tool}
	r := regexp.MustCompile("[^\\s]+")
	myArgs2 := r.FindAllString(args, -1)
	myArgs = append(myArgs, myArgs2...)

	cmd := exec.Command("cmd", myArgs...)

	return cmd
}


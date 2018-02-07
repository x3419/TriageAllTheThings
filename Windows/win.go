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

func blkcalc(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\blkcalc.exe")
	runDefault(cmd)
}

func blkcat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\blkcat.exe")
	runDefault(cmd)
}

func blkstat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\blkstat.exe")
	runDefault(cmd)
}

func fcat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\fcat.exe")
	runDefault(cmd)
}

func ffind(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\ffind.exe")
	runDefault(cmd)
}

func fls(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\fls.exe")
	runDefault(cmd)
}

func fsstat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\fsstat.exe")
	runDefault(cmd)
}

func hfind(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\hfind.exe")
	runDefault(cmd)
}

func icat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\icat.exe")
	runDefault(cmd)
}

func ifind(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\ifind.exe")
	runDefault(cmd)
}

func ils(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\ils.exe")
	runDefault(cmd)
}

func img_cat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\img_cat.exe")
	runDefault(cmd)
}

func img_stat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\img_stat.exe")
	runDefault(cmd)
}

func istat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\istat.exe")
	runDefault(cmd)
}

func jcat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\jcat.exe")
	runDefault(cmd)
}

func jls(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\jls.exe")
	runDefault(cmd)
}

func mmcat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\mmcat.exe")
	runDefault(cmd)
}

func mmls(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\mmls.exe")
	runDefault(cmd)
}

func mmstat(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\mmstat.exe")
	runDefault(cmd)
}


func tsk_comparedir(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\tsk_comparedir.exe")
	runDefault(cmd)
}

func tsk_gettimes(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\tsk_gettimes.exe")
	runDefault(cmd)
}


func tsk_loaddb(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\tsk_loaddb.exe")
	runDefault(cmd)
}


func tsk_recover(args string) {
	cmd :=  cmdTool(args, "sleuthkit\\bin\\tsk_recover.exe")
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


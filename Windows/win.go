package Windows

import (
	"os/exec"
	"fmt"
	"regexp"
	"bufio"
	"strings"
)




func BulkExtractor(args string) {

	cmd :=  runTool(args, "bulk_extractor32.exe")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		if strings.Contains(m, "(") {
			fmt.Println("BulkExtractor: " + m[strings.Index(m, "(") + 1:strings.Index(m,")")] + "% Complete")
		}
	}

	cmd.Wait()


}

func Fiwalk(args string) {
	cmd :=  runTool(args, "fiwalk-0.6.3.exe")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	cmd.Wait()
}



func runTool(args string, tool string) *exec.Cmd {
	myArgs := []string{"/C", "Windows\\" + tool}
	r := regexp.MustCompile("[^\\s]+")
	myArgs2 := r.FindAllString(args, -1)
	myArgs = append(myArgs, myArgs2...)

	cmd := exec.Command("cmd", myArgs...)

	return cmd
}


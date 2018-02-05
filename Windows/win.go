package Windows

import (
	"os/exec"
	"fmt"
	"regexp"
	"bufio"
	"strings"
)
func BulkExtractor(args string) {

	myArgs := []string{"/C", "Windows\\bulk_extractor32.exe"}
	r := regexp.MustCompile("[^\\s]+")
	myArgs2 := r.FindAllString(args, -1)
	myArgs = append(myArgs, myArgs2...)

	cmd := exec.Command("cmd", myArgs...)

	//var out bytes.Buffer
	//var stderr bytes.Buffer
	//cmd.Stdout = &out
	//cmd.Stderr = &stderr
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	//	return
	//}
	//fmt.Println("Result: " + out.String())

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

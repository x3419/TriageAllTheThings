package Windows

import (
	"os/exec"
	"fmt"
	"bytes"
	"regexp"
)
func BulkExtractor(args string) {
	fmt.Println(args)
	myArgs := []string{"/C", "Windows\\bulk_extractor32.exe"}

	r := regexp.MustCompile("[^\\s]+")
	myArgs2 := r.FindAllString(args, -1)

	myArgs = append(myArgs, myArgs2...)
	//cmd := exec.Command("cmd", "/C", "Windows\\bulk_extractor32.exe", strings.Fields(args))
	cmd := exec.Command("cmd", myArgs...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())

}

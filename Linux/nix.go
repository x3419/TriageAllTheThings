package Linux

import (
	"os/exec"
	"bufio"
	"fmt"
	"regexp"
	"os"
)

func Ps(args string) {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "ps"
	if cmdOut, err = exec.Command(cmdName).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running ps ", err)
		os.Exit(1)
	}
	sha := string(cmdOut)
	fmt.Println(sha)
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
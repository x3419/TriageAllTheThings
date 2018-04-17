package Osutil

import (
	"Capstone/Configuration"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"testing"
)

// NOTE: This test file requires administrative priviledges because of READ/WRITE capabilities. Please use "sudo go test"

func TestCmdTool(t *testing.T) {
	arg := "arg1 arg2 arg3 arg4"
	tool := "toolName"
	var cmd *exec.Cmd = cmdTool(tool, arg)

	if cmd.Path != arg {
		t.Error("cmdTool argument serialization not working")
	}

	if cmd.Args[1] != tool {
		t.Error("cmdTool executable name serialization not working")
	}
}

func TestDefaultParse(t *testing.T) {
	name := "helloworld"
	command := "hello world"
	prog := "echo"
	var cmd *exec.Cmd = cmdTool(command, prog)
	var wg sync.WaitGroup

	err := defaultParse(name, cmd, &wg)
	if err != nil {
		t.Fail()
	}

	b, err := ioutil.ReadFile("./Output/" + name + ".txt")

	if err != nil {
		t.Error("Error reading output file. Run this test with sudo!")
	}

	if string(b)[0:len(command)] != command {
		t.Error("Actual output of default parse does not match the expected output")
	}

	// This was giving me some very odd permission issues so I'm commenting this for now
	//RemoveContents("./Output/")

	wg.Wait()

}

func TestDefaultFunc(t *testing.T) {

	testTool := Configuration.Tool{
		Name:    "test",
		Enabled: true,
		Args:    "test args",
		Path:    "v^2<.!",
	}

	var wg sync.WaitGroup

	err := defaultFunc(testTool, &wg)

	if err != nil {
		t.Fail()
	}

	testTool = Configuration.Tool{
		Name:    "test",
		Enabled: true,
		Args:    "test args",
		Path:    "__12jdA!",
	}

	err = defaultFunc(testTool, &wg)

	if err != nil {
		t.Fail()
	}

	testTool = Configuration.Tool{
		Name:    "test",
		Enabled: true,
		Args:    "test args",
		Path:    "&^@8jcma0",
	}

	err = defaultFunc(testTool, &wg)

	if err != nil {
		t.Fail()
	}

	testTool = Configuration.Tool{
		Name:    "test",
		Enabled: true,
		Args:    "test args",
		Path:    "_*@#){",
	}

	err = defaultFunc(testTool, &wg)

	if err != nil {
		t.Fail()
	}

}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

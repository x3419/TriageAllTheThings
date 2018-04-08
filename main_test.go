package main

import (
	"testing"
	"Capstone/Configuration"
	"Capstone/Structs"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {

	testConfig := Configuration.Config{}

	var config Configuration.Config = ParseConfig("foo")

	if(config != testConfig){
		t.Failed()
	}

	testConfig2 := ParseConfig("TestConfig1.txt")

	if(testConfig2.WinTools.Blkstat.Enabled){
		t.Failed()
	}
	if(testConfig2.NixTools.Ps.Enabled){
		t.Failed()
	}
	if(testConfig2.MacTools.Pac4Mac.Enabled){
		t.Failed()
	}

	testConfig3 := ParseConfig("TestConfig2.txt")

	if(testConfig3.WinTools.BulkExtractor.Enabled != true){
		t.Failed()
	}
}

func TestGUI(t *testing.T) {
	configFile := ParseConfig("TestConfig2.txt")
	tasks := make(chan Structs.Result, 64)

	assert.Panics(t, func() { makeGUI(configFile, tasks) }, "The code did not panic")

	//makeGUI(configFile, tasks)


}
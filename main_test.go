package main

import (
	"testing"
	"Capstone/Configuration"
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
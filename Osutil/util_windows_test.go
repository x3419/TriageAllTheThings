// +build_windows
package Osutil

import (
	"testing"
	"Capstone/Configuration"
)

func TestOsutil(t *testing.T) {
	badCmd := cmdTool("", "")
	if(badCmd != nil){
		t.Failed()
	}
}

func TestBuildUi(t *testing.T){

	testTool := Configuration.Tool{
		Name:    "test",
		Enabled: true,
		Args:    "test args",
		Path:    "/folder/directory/executable.exe",
	}

	var testConfig Configuration.Config = Configuration.Config{
		Tool: []Configuration.Tool{
			testTool,
		},
		RelativePath : false,
	}

	buildGood := BuildUi(nil, nil, nil, testConfig)
	if(buildGood){
		t.Fail()
	}
}
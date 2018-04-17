package main

import (
	"testing"
	"Capstone/Configuration"
	"os"
)

func TestConfig(t *testing.T) {

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

	testString := `
		[[Tool]]
		Name = "test"
		Enabled = true
		Args = "test args"
		Path = "/folder/directory/executable.exe"
	`

	var testConfig2 Configuration.Config = TomlParseConfig(testString)

	compareConfigs(testConfig, testConfig2, t)

	if(checkConfig(testConfig)){
		t.Fail()
	}

	if(checkConfig(testConfig2)) {
		t.Fail()
	}

}

func TestConfig2(t *testing.T) {
	var testConfig []Configuration.Config = []Configuration.Config{
		Configuration.Config{
			Tool: []Configuration.Tool{
				Configuration.Tool{
					Name:    "test",
					Enabled: true,
					Args:    "test args",
					Path:    "/folder/directory/executable.exe",
				},
			},
			RelativePath: false,
		},
		Configuration.Config{
			Tool: []Configuration.Tool{
				Configuration.Tool{
					Name:    "test2",
					Enabled: true,
					Args:    "?ja@592L",
					Path:    "C:/NotAFolder",
				},
			},
			RelativePath: false,
		},
		Configuration.Config{
			Tool: []Configuration.Tool{
				Configuration.Tool{
					Name:    "test3",
					Enabled: false,
					Args:    "||/*-426.",
					Path:    "??!!~`/\\",
				},
			},
			RelativePath: false,
		},
		Configuration.Config{
			Tool: []Configuration.Tool{
				Configuration.Tool{
					Name:    "test4",
					Enabled: true,
					Args:    "test args",
					Path:    "__292m_+\"]",
				},
			},
			RelativePath: false,
		},
	}

	for _, c := range(testConfig){
		if(checkConfig(c)){
			t.Fail()
		}
	}
}

func compareConfigs(config1 Configuration.Config, config2 Configuration.Config, t *testing.T) {
	for _, t1 := range config2.Tool {
		for _, t2 := range config2.Tool {
			if( (t1.Name != t2.Name) ||
				(t1.Enabled != t2.Enabled) ||
				(t1.Path != t2.Path) ||
				(t1.Args != t2.Args)){
					t.Fail()
			}
		}
	}
	if(config1.RelativePath != config2.RelativePath){
		t.Fail()
	}
}

func TestDumpTools(t *testing.T){
	err := dumpTools("/fake/path", os.FileInfo(nil), nil)
	if(err != nil){
		t.Fail()
	}
}
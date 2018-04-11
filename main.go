package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"encoding/json"
	"runtime"
	Configuration "Capstone/Configuration"
	"strings"
	"Capstone/Osutil"
)

func main() {

	configPathPtr := flag.String("config", "Configuration/config.txt", "Location of the configuration file")
	flag.Parse()
	//var config Configuration.Config = ParseConfig(*configPathPtr)

	var config Configuration.DynamicConfig = DynamicParseConfig(*configPathPtr)

	fmt.Println(config)

	fmt.Println(strings.Title(runtime.GOOS) + " OS detected\nEnabled tools will begin to run in parallel. This may take some time and will slow the system down, so please be patient.")

	var os Osutil.ToolRunner = Osutil.Util{}
	os.MakeGUI(config)

}

func ParseConfig(configFile string) Configuration.Config {

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("Unable to open config file: ", err)
	}

	myConfig := Configuration.Config{}

	if err := json.Unmarshal(b, &myConfig); err != nil {
		fmt.Println("Error! Problem parsing the configuration file - please ensure that it reflects the example on Github.\n", err)
		panic(err)
		return Configuration.Config{}
	}

	return myConfig

}


func DynamicParseConfig(configFile string) Configuration.DynamicConfig {

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("Unable to open config file: ", err)
	}

	myConfig := Configuration.DynamicConfig{}

	if err := json.Unmarshal(b, &myConfig); err != nil {
		fmt.Println("Error! Problem parsing the configuration file - please ensure that it reflects the example on Github.\n", err)
		panic(err)
		return Configuration.DynamicConfig{}
	}

	return myConfig

}


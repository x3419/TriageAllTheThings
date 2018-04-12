package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	Configuration "Capstone/Configuration"
	"github.com/BurntSushi/toml"
	"log"
	"strings"
	"runtime"
	"Capstone/Osutil"
)

func main() {

	configPathPtr := flag.String("config", "Configuration/config.txt", "Location of the configuration file")
	flag.Parse()

	var config Configuration.Config = TomlParseConfig(*configPathPtr)

	fmt.Println(strings.Title(runtime.GOOS) + " OS detected\nEnabled tools will begin to run in parallel. This may take some time and will slow the system down, so please be patient.")

	var os Osutil.ToolRunner = Osutil.Util{}
	os.MakeGUI(config)

}

func TomlParseConfig(configFile string) Configuration.Config {

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("Unable to open config file: ", err)
	}

	var config Configuration.Config
	if _, err := toml.Decode(string(b), &config); err != nil {
		log.Fatal(err)
	}

	return config

}

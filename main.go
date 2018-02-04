package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"encoding/json"
	"runtime"
	Configuration "Capstone/Configuration"
)



func main() {

	configPathPtr := flag.String("config", "Configuration/config.txt", "Loction of the configuration file")
	flag.Parse()
	var config Configuration.Config = parseConfig(*configPathPtr)
	
	if runtime.GOOS == "windows" {
		windowsTools(config)
	} else if runtime.GOOS == "linux" {
		fmt.Println("GNU/Linux compatibility coming soon!")
	} else if runtime.GOOS == "darwin" {
		fmt.Println("OSX compatibility coming soon!")
	}

}


func parseConfig(configFile string) Configuration.Config {

	b, err := ioutil.ReadFile(configFile) 
    if err != nil {
        fmt.Println("Unable to open config file: ", err)
    }

    myConfig := Configuration.Config{}

    if err := json.Unmarshal(b, &myConfig); err != nil {
        fmt.Println("Error!\n", err)
    }

    return myConfig

}

func windowsTools(config Configuration.Config) {
	
}
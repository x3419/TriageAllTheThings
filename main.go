package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"encoding/json"
)

    

func main() {

	configPathPtr := flag.String("config", "config.txt", "Loction of the configuration file")
	flag.Parse()
	config Config = parseConfig(*configPathPtr)
	
	if runtime.GOOS == "windows" {
		windowsTools(config)
	} else if untime.GOOS == "linux" {
		fmt.Println("GNU/Linux compatibility coming soon!")
	} else if untime.GOOS == "darwin" {
		fmt.Println("OSX compatibility coming soon!")
	}

}


func parseConfig(configFile string) Config {

	b, err := ioutil.ReadFile(configFile) 
    if err != nil {
        fmt.Println("Unable to open config file: ", err)
    }

    myConfig := Config{}

    if err := json.Unmarshal(b, &myConfig); err != nil {
        fmt.Println("Error!\n", err)
    }

    return myConfig

}

func windowsTools(config Config) {
	
}
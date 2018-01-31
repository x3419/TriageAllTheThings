package main

import (
	"fmt"
	"flag"
	"io/ioutil"
)

    

func main() {

	configPtr := flag.String("config", "config.txt", "Loction of the configuration file")
	flag.Parse()
	parseConfig(*configPtr)

}


func parseConfig(configFile string) {

	b, err := ioutil.ReadFile(configFile) 
    if err != nil {
        fmt.Println("Unable to open config file: ", err)
    }

    fmt.Println(string(b))



}
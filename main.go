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
	"github.com/GeertJohan/go.rice"
	"os"
	"path/filepath"
)

func main() {

	configPathPtr := flag.String("config", "Configuration/config.txt", "Location of the configuration file")
	portable := flag.Bool("portable", true, "Enable the bundling process that embeds tools within the executable")
	flag.Parse()

	var config Configuration.Config

	fmt.Println(*portable)

	if *portable {

		// bundling tools together
		conf := rice.Config{
			LocateOrder: []rice.LocateMethod{rice.LocateEmbedded, rice.LocateAppended, rice.LocateFS},
		}
		box, err := conf.FindBox("Tools")
		configBox, err := conf.FindBox("Configuration") // I'd use the config path ptr dir but rice needs a string literal... :-/


		if err != nil {
			log.Fatalf("error opening rice.Box: %s\n", err)
		}

		configString, err := configBox.String(filepath.Base(*configPathPtr))
		if err != nil {
			log.Fatal("Error unbundling config")
		}

		config = TomlParseConfig(configString)

		//box.Walk("Tools", dumpTools)
		CreateDirIfNotExist("Tools")
		for _,t := range(config.Tool){
			if t.Enabled {

				data, err := box.Bytes(t.Path)
				if err != nil {
					fmt.Println(t.Path + " not bundled properly. Skipping")
					continue
				}

				var path string
				if(config.RelativePath){
					path = "Tools/" + t.Path
				} else {
					path = "Tools/" + filepath.Base(t.Path)
				}
				err = ioutil.WriteFile(path, data, 755)
				if(err != nil){
					log.Fatalf("error unbundling file: %s\n", err)
				}
			}

		}

		// bundling tools together

	} else {
		b, err := ioutil.ReadFile(*configPathPtr)
		if err != nil {
			fmt.Println("Unable to open config file: ", err)
		}
		config = TomlParseConfig(string(b))
	}


	fmt.Println(strings.Title(runtime.GOOS) + " OS detected\nEnabled tools will begin to run in parallel. This may take some time and will slow the system down, so please be patient.")

	var os Osutil.ToolRunner = Osutil.Util{}
	os.MakeGUI(config)

}

func dumpTools(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	if !info.IsDir(){
		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Print(err)
			return nil
		}
		CreateDirIfNotExist("Tools")
																			// Read Exec
		err = ioutil.WriteFile("Tools/" + info.Name() + ".exe", data, 755)
		fmt.Print("err? ")
		log.Print(err)
	}
	return err
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		fmt.Println(os.Getwd())
		if err != nil {
			panic(err)
		}
	}
}

func TomlParseConfig(configString string) Configuration.Config {


	var config Configuration.Config
	if _, err := toml.Decode(configString, &config); err != nil {
		log.Fatal(err)
	}

	return config

}

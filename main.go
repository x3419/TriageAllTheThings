package main

import (
	Configuration "Capstone/Configuration"
	"Capstone/Osutil"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/GeertJohan/go.rice"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// main is the main function, first to run, that deals with the config file, boxing/unboxing for delployment on all operating systems
func main() {

	// parse arguments
	configPathPtr := flag.String("config", "Configuration/config.txt", "Location of the configuration file")
	portable := flag.Bool("portable", true, "Enable the bundling process that embeds tools within the executable")
	flag.Parse()

	var config Configuration.Config

	if *portable {

		// bundling tools together
		conf := rice.Config{
			LocateOrder: []rice.LocateMethod{rice.LocateEmbedded, rice.LocateAppended, rice.LocateFS},
		}
		
		// box it up
		box, err := conf.FindBox("Tools")
		configBox, err := conf.FindBox("Configuration") // I'd use the config path ptr dir but rice needs a string literal... :-/

		if err != nil {
			log.Fatalf("error opening rice.Box: %s\n", err)
		}

		// unbox it
		configString, err := configBox.String(filepath.Base(*configPathPtr))
		if err != nil {
			log.Fatal("Error unbundling config")
		}

		config = TomlParseConfig(configString)

		if checkConfig(config) {
			log.Fatal("Directories or files different than specified in configuration file")
		}

		//box.Walk("Tools", dumpTools) // not sure why this isn't working as expected :-/ well here's the backup plan...
		CreateDirIfNotExist("Tools")
		for _, t := range config.Tool {
			if t.Enabled {

				data, err := box.Bytes(t.Path)
				if err != nil {
					fmt.Println(t.Path + " not bundled properly. Skipping")
					continue
				}

				var path string
				if config.RelativePath {
					path = "Tools/" + t.Path
				} else {
					path = "Tools/" + filepath.Base(t.Path)
				}
				err = ioutil.WriteFile(path, data, 755)
				if err != nil {
					log.Fatalf("error unbundling file: %s\n", err)
				}
			}

		}

		// bundling tools together

	} else {
		
		// read config from disk and parse
		b, err := ioutil.ReadFile(*configPathPtr)
		if err != nil {
			fmt.Println("Unable to open config file: ", err)
		}
		config = TomlParseConfig(string(b))
		if checkConfig(config) {
			log.Fatal("Directories or files different than specified in configuration file")
		}
	}

	fmt.Println(strings.Title(runtime.GOOS) + " OS detected\nEnabled tools will begin to run in parallel. This may take some time and will slow the system down, so please be patient.")

	var os Osutil.ToolRunner = Osutil.Util{}
	os.MakeGUI(config)

}

// checkConfig return whether the config works as expected
func checkConfig(config Configuration.Config) bool {
	okay := "abcdefghijklmnopqrstuvwxyz1234567890_-.\\:"
	for _, t := range config.Tool {
		for _, char := range t.Path {
			if !strings.Contains(okay, strings.ToLower(string(char))) {
				return false
			}
		}
		if _, err := os.Stat(t.Path); os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// dumpTools is supposed to be a function used on every file in a given folder. 
// isn't working for go.rice as expected
func dumpTools(path string, info os.FileInfo, err error) error {
	if err != nil || info == nil {
		log.Print(err)
		return nil
	}
	if !info.IsDir() {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Print(err)
			return nil
		}
		CreateDirIfNotExist("Tools")
		// Read Exec
		err = ioutil.WriteFile("Tools/"+info.Name()+".exe", data, 755)
		fmt.Print("err? ")
		log.Print(err)
	}
	return err
}

// CreateDirIfNotExist checks whether a directory exists. If it does not, it makes the directory.
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0644)
		if err != nil {
			panic(err)
		}
	}
}

// TomlParseConfig parses a config from string -> Configuration struct
func TomlParseConfig(configString string) Configuration.Config {

	var config Configuration.Config
	if _, err := toml.Decode(configString, &config); err != nil {
		log.Fatal(err)
	}

	return config
}

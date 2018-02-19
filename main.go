package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"encoding/json"
	"runtime"
	Configuration "Capstone/Configuration"
	Windows "Capstone/Windows"
	Linux "Capstone/Linux"
	"strings"
)



func main() {

	configPathPtr := flag.String("config", "Configuration/config.txt", "Location of the configuration file")
	flag.Parse()
	var config Configuration.Config = ParseConfig(*configPathPtr)

	fmt.Println(strings.Title(runtime.GOOS) + " OS detected\nEnabled tools will begin to run in parallel. This may take some time and will slow the system down, so please be patient.")

	if runtime.GOOS == "windows" {
		windowsTools(config)
	} else if runtime.GOOS == "linux" {
		fmt.Println("GNU/Linux compatibility coming soon!")
	} else if runtime.GOOS == "darwin" {
		fmt.Println("OSX compatibility coming soon!")
	} else {
		fmt.Println(strings.Title(runtime.GOOS) + " OS is not supported in this project.")
	}

}


func ParseConfig(configFile string) Configuration.Config {

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("Unable to open config file: ", err)
	}

	myConfig := Configuration.Config{}

	if err := json.Unmarshal(b, &myConfig); err != nil {
		fmt.Println("Error! Problem parsing the configuration file - please ensure that it reflects the example on Github.\n", err)
		return Configuration.Config{}
	}

	return myConfig

}

func windowsTools(config Configuration.Config) {

	win := config.WinTools
	nix := config.NixTools

	// Windows tools

	if win.BulkExtractor.Enabled {
		Windows.BulkExtractor(win.BulkExtractor.Args)
	}
	if win.Fiwalk.Enabled {
		Windows.Fiwalk(win.Fiwalk.Args)
	}
	if win.Blkcalc.Enabled {
		Windows.Blkcalc(win.Blkcalc.Args)
	}
	if win.Blkcat.Enabled {
		Windows.Blkcat(win.Blkcat.Args)
	}
	if win.Blkls.Enabled {
		Windows.Blkls(win.Blkls.Args)
	}
	if win.Blkstat.Enabled {
		Windows.Blkstat(win.Blkstat.Args)
	}
	if win.Fcat.Enabled {
		Windows.Fcat(win.Fcat.Args)
	}
	if win.Ffind.Enabled {
		Windows.Ffind(win.Ffind.Args)
	}
	if win.Fls.Enabled {
		Windows.Fls(win.Fls.Args)
	}
	if win.Fsstat.Enabled {
		Windows.Fsstat(win.Fsstat.Args)
	}
	if win.Hfind.Enabled {
		Windows.Hfind(win.Hfind.Args)
	}
	if win.Icat.Enabled {
		Windows.Icat(win.Icat.Args)
	}
	if win.Ifind.Enabled {
		Windows.Ifind(win.Ifind.Args)
	}
	if win.Ils.Enabled {
		Windows.Ils(win.Ils.Args)
	}
	if win.Imgcat.Enabled {
		Windows.Img_cat(win.Imgcat.Args)
	}
	if win.Imgstat.Enabled {
		Windows.Img_stat(win.Imgstat.Args)
	}
	if win.Istat.Enabled {
		Windows.Istat(win.Istat.Args)
	}
	if win.Jcat.Enabled {
		Windows.Jcat(win.Jcat.Args)
	}
	if win.Jls.Enabled {
		Windows.Img_cat(win.Jls.Args)
	}
	if win.Mmcat.Enabled {
		Windows.Mmcat(win.Mmcat.Args)
	}
	if win.Mmls.Enabled {
		Windows.Mmls(win.Mmls.Args)
	}
	if win.Mmstat.Enabled {
		Windows.Mmstat(win.Mmstat.Args)
	}
	if win.TskCompareDir.Enabled {
		Windows.Tsk_comparedir(win.TskCompareDir.Args)
	}
	if win.TskGetTimes.Enabled {
		Windows.Tsk_gettimes(win.TskCompareDir.Args)
	}
	if win.TskLoaddb.Enabled {
		Windows.Tsk_loaddb(win.TskCompareDir.Args)
	}
	if win.TskRecover.Enabled {
		Windows.Tsk_recover(win.TskRecover.Args)
	}
	if win.Tcpflow.Enabled {
		Windows.Tcpflow(win.Tcpflow.Args)
	}
	if win.WinPrefetch.Enabled {
		Windows.WinPrefetch(win.WinPrefetch.Args)
	}

	// GNU/Linux tools

	if nix.Ps.Enabled {
		Linux.Ps(nix.Ps.Args)
	}

}


package Osutil

import (
	"Capstone/Windows"
	//"Capstone/Linux"
	"github.com/ProtonMail/ui"
	"Capstone/Configuration"
	"Capstone/Structs"
)

type Util struct {}

func (u Util) BuildUi(myBox *ui.Box, uiCompMap map[string]Structs.UIComp, config Configuration.Config, tsks chan <- Structs.Result) {

	win := config.WinTools
	//nix := config.NixTools

	// Windows tools
	if win.BulkExtractor.Enabled {
		Windows.BulkExtractor(win.BulkExtractor.Args, tsks)
	}
	if win.Fiwalk.Enabled {
		Windows.Fiwalk(win.Fiwalk.Args, tsks)
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
		uiComp := uiCompMap["tcpFlow"]
		//addToolToUI(myBox, "tcpflow", uiComp.Label, uiComp.Output)
		Windows.Tcpflow(win.Tcpflow.Args, uiComp.Label, uiComp.Output)
	}
	if win.WinPrefetch.Enabled {

		uiComp := uiCompMap["winprefetch"]
		//addToolToUI(myBox, "WinPrefetch", uiComp.Label, uiComp.Output)
		Windows.WinPrefetch(win.WinPrefetch.Args, uiComp.Label, uiComp.Output)

	}
	if win.MFTDump.Enabled {

		uiComp := uiCompMap["mftdump"]
		//addToolToUI(myBox, "MFTDump", uiComp.Label, uiComp.Output)
		Windows.MftDump(win.MFTDump.Args, uiComp.Label, uiComp.Output)
	}

}

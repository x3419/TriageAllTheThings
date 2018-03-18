package Osutil

import (
	//"Capstone/Windows"
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
		BulkExtractor(win.BulkExtractor.Args, tsks)
	}
	if win.Fiwalk.Enabled {
		Fiwalk(win.Fiwalk.Args, tsks)
	}
	if win.Blkcalc.Enabled {
		Blkcalc(win.Blkcalc.Args)
	}
	if win.Blkcat.Enabled {
		Blkcat(win.Blkcat.Args)
	}
	if win.Blkls.Enabled {
		Blkls(win.Blkls.Args)
	}
	if win.Blkstat.Enabled {
		Blkstat(win.Blkstat.Args)
	}
	if win.Fcat.Enabled {
		Fcat(win.Fcat.Args)
	}
	if win.Ffind.Enabled {
		Ffind(win.Ffind.Args)
	}
	if win.Fls.Enabled {
		Fls(win.Fls.Args)
	}
	if win.Fsstat.Enabled {
		Fsstat(win.Fsstat.Args)
	}
	if win.Hfind.Enabled {
		Hfind(win.Hfind.Args)
	}
	if win.Icat.Enabled {
		Icat(win.Icat.Args)
	}
	if win.Ifind.Enabled {
		Ifind(win.Ifind.Args)
	}
	if win.Ils.Enabled {
		Ils(win.Ils.Args)
	}
	if win.Imgcat.Enabled {
		Img_cat(win.Imgcat.Args)
	}
	if win.Imgstat.Enabled {
		Img_stat(win.Imgstat.Args)
	}
	if win.Istat.Enabled {
		Istat(win.Istat.Args)
	}
	if win.Jcat.Enabled {
		Jcat(win.Jcat.Args)
	}
	if win.Jls.Enabled {
		Img_cat(win.Jls.Args)
	}
	if win.Mmcat.Enabled {
		Mmcat(win.Mmcat.Args)
	}
	if win.Mmls.Enabled {
		Mmls(win.Mmls.Args)
	}
	if win.Mmstat.Enabled {
		Mmstat(win.Mmstat.Args)
	}
	if win.TskCompareDir.Enabled {
		Tsk_comparedir(win.TskCompareDir.Args)
	}
	if win.TskGetTimes.Enabled {
		Tsk_gettimes(win.TskCompareDir.Args)
	}
	if win.TskLoaddb.Enabled {
		Tsk_loaddb(win.TskCompareDir.Args)
	}
	if win.TskRecover.Enabled {
		Tsk_recover(win.TskRecover.Args)
	}
	if win.Tcpflow.Enabled {
		uiComp := uiCompMap["tcpFlow"]
		//addToolToUI(myBox, "tcpflow", uiComp.Label, uiComp.Output)
		Tcpflow(win.Tcpflow.Args, uiComp.Label, uiComp.Output)
	}
	if win.WinPrefetch.Enabled {

		uiComp := uiCompMap["winprefetch"]
		WinPrefetch(win.WinPrefetch.Args, uiComp)

	}
	if win.MFTDump.Enabled {

		uiComp := uiCompMap["mftdump"]
		MftDump(win.MFTDump.Args, uiComp.Label, uiComp.Output)
	}

	// Copying the windows file




}

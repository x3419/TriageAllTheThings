package Configuration

type Config struct {
	WinTools WindowsTools
	NixTools LinuxTools
	MacTools OSXTools
}

type WindowsTools struct {
	BulkExtractor Tool
	Fiwalk Tool
	Frag_find Tool
	Tcpflow Tool
	Afxml Tool
	MFTDump Tool
	WinPrefetch Tool
	RegRipper Tool
	Blkcalc Tool
	Blkcat Tool
	Blkls Tool
	Blkstat Tool
	Fcat Tool
	Ffind Tool
	Fls Tool
	Fsstat Tool
	Hfind Tool
	Icat Tool
	Ifind Tool
	Ils Tool
	Imgcat Tool
	Imgstat Tool
	Istat Tool
	Jcat Tool
	Jls Tool
	Mmcat Tool
	Mmls Tool
	Mmstat Tool
	TskCompareDir Tool
	TskGetTimes Tool
	TskLoaddb Tool
	TskRecover Tool


}

type LinuxTools struct {
	Ps Tool
	Mrutools Tool
	LinuxForensicsTool Tool
	BulkExtractor Tool
	ENT Tool
	Fdupes Tool
	Foremost Tool
}

type OSXTools struct {
	OSXAudiotr Tool
	KnockKnock Tool
	Pac4Mac Tool
	KeychainForensicTool Tool
	OSXCollector Tool
}

type Tool struct {
	Enabled bool
	Args string
}

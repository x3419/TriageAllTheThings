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
	SleuthKit Tool
	MFTDump Tool
	WinPrefetch Tool
	RegRipper Tool
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

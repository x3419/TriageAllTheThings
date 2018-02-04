package Configuration

type Config struct {
	BulkExtractor Tool
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
	Ps bool
	Mrutools bool
	LinuxForensicsTool bool
	BulkExtractor bool
	ENT bool
	Fdupes bool
	Foremost bool
}

type OSXTools struct {
	OSXAudiotr bool
	KnockKnock bool
	Pac4Mac bool
	KeychainForensicTool bool
	OSXCollector bool
}

type Tool struct {
	Enabled bool
	Args string
}


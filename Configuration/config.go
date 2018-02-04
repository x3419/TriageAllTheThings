package Configuration

type Config struct {
	winTools WindowsTools
	nixTools LinuxTools
	macTools OSXTools
}

type WindowsTools struct {
	bulk_extractor bool
	fiwalk bool
	frag_find bool
	tcpflow bool
	afxml bool
	SleuthKit bool
	MFTDump bool
	WinPrefetch bool
	RegRipper bool
}

type LinuxTools struct {
	ps bool
	mrutools bool
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

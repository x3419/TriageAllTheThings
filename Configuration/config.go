package Configuration

type Config struct {
	Tool         []Tool
	RelativePath bool
}

type Tool struct {
	Name    string
	Enabled bool
	Args    string
	Path    string
}

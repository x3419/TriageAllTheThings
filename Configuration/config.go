// Package Configuration contains the things necessary for configuration.
package Configuration

// Config struct is the configuration object that is seralized.
type Config struct {
	Tool         []Tool
	RelativePath bool
}

// Tool contains a given tools settings.
type Tool struct {
	Name    string
	Enabled bool
	Args    string
	Path    string
}

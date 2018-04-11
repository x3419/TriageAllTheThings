package Configuration


type DynamicConfig struct {
	Tools []DynamicTool
}

type DynamicTool struct {
	Name string
	Enabled bool
	Args string
	Location string
}

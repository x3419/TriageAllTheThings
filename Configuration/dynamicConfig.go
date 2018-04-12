package Configuration


type DynamicConfig struct {
	Tool []DynamicTool
}

type DynamicTool struct {
	Name string
	Enabled bool
	Args string
	Location string
}

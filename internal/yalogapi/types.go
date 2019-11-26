package yalogapi

type Types struct {
	Items []*TypeItem `mapstructure:"items"`
}
type TypeItem struct {
	Name  string `mapstructure:"name"`
	Value string `mapstructure:"value"`
}

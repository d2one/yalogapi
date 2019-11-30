package yalogapi

type Types struct {
	Items []*TypeItem `mapstructure:"items"`
}
type TypeItem struct {
	Name  string `mapstructure:"name"`
	Value string `mapstructure:"value"`
}

func (types *Types) Build(fields []string) map[string]string {
	result := make(map[string]string)
	for _, field := range fields {
		for _, typeItem := range types.Items {
			if field == typeItem.Name {
				result[field] = typeItem.Value
			}
		}
	}
	return result
}

package yalogapi

import (
	"fmt"
)

type YaLogApi struct {
	config *Config
}

func main() {

}

func NewYaLogApi(config *Config) *YaLogApi {
	fmt.Println(config)
	return &YaLogApi{config: config}
}

func (yalogapi *YaLogApi) Run() {
	fmt.Println(yalogapi.config)
	fmt.Println("YaLogApi")
}

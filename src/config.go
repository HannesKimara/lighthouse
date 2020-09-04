package lighthouse

import (
	"io/ioutil"
	"os"
	"log"

	"github.com/go-yaml/yaml"
)

// Config ...
type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Mongo struct {
		URL string `yaml:"CLUSTER_URL"`
		Database string `yaml:"DATABASE"`
	} `yaml:"mongo"`
}

// Load ...
func (c *Config) LoadConf() *Config {
	data, err := ioutil.ReadFile("conf.yml")
	if err != nil {
		log.Fatal(err)
	}

	data = []byte(os.ExpandEnv(string(data)))

	err = yaml.Unmarshal([]byte(data), c)

	if err != nil {
		log.Fatal(err)
	}

	return c
}

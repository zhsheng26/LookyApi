package conf

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

/**
* @author zhangsheng
* @date 2019/6/12
 */

type DbConf struct {
	Oracle Oracle
	Redis  Redis
}

type Oracle struct {
	Host     string
	Port     string
	Username string
	Password string
}
type Redis struct {
	Host     string
	Port     string
	Database int
	Password string
}

func (redis Redis) Address() string {
	return redis.Host + ":" + redis.Port
}

type Consul struct {
	Host       string
	Port       string
	HealthPath string `yaml:"health-check-path"`
	Tags       []string
}

type ServerConf struct {
	Host   string
	Port   int
	Name   string
	Env    string
	Consul Consul
}
type Setting struct {
	Db       DbConf
	Server   ServerConf
	LogLevel string
}

func defaultSetting() *Setting {
	return &Setting{
		Db: DbConf{
			Oracle: Oracle{
				Host:     "192.168.1.4",
				Port:     "1521",
				Username: "us_db_dev",
				Password: "123456",
			},
			Redis: Redis{
				Host:     "192.168.1.2",
				Port:     "18160",
				Database: 2,
				Password: "",
			},
		},
		Server: ServerConf{
			Host: "",
			Port: 9080,
			Name: "us_diagram_go_zsheng",
			Env:  "dev",
			Consul: Consul{
				Host:       "192.168.1.2",
				Port:       "18500",
				HealthPath: "/health",
				Tags:       []string{},
			},
		},
		LogLevel: "debug",
	}
}

var parseYmlError = errors.New("error while trying to decode configuration")

func YAML(filename ...string) (*Setting, error) {
	setting := defaultSetting()
	if len(filename) == 0 {
		return setting, nil
	}
	yamlAbsPath, err := filepath.Abs(filename[0])
	if err != nil {
		return setting, parseYmlError
	}

	data, err := ioutil.ReadFile(yamlAbsPath)
	if err != nil {
		return setting, parseYmlError
	}

	if err := yaml.Unmarshal(data, &setting); err != nil {
		return setting, parseYmlError
	}
	return setting, nil
}

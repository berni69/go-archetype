package consul

import (
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/berni69/go-archetype/utils"
	consulapi "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
)

func LookupService(serviceName string) (string, error) {
	config := consulapi.DefaultConfig()
	co, err := consulapi.NewClient(config)
	if err != nil {
		log.Debug(err)
		return "", err
	}
	service, _, err := co.Catalog().Service(serviceName, "", nil)
	if err != nil {
		log.Debug(err)
		return "", err
	}

	r := utils.GetRandomInt(len(service) - 1)
	address := service[r].Address
	port := service[r].ServicePort
	return fmt.Sprintf("http://%s:%v", address, port), nil
}

func LoadConsulConfig(path string, consulConfig interface{}) error {
	log.Info("_______________________________________________________________________________________________________________________")
	log.Infof("Consul Properties ================<< %s/v1/kv/%s >>======================================================",
		utils.GetEnv("CONSUL_HTTP_ADDR", "127.0.0.1"), path)
	log.Info("_______________________________________________________________________________________________________________________")

	config := consulapi.DefaultConfig()
	co, err := consulapi.NewClient(config)
	if err != nil {
		log.Debug(err)
		return err
	}
	pair, _, err := co.KV().Get(path, nil)
	if err != nil {
		log.Debug(err)
		return err
	}

	log.Debug(path)
	log.Debug(string(pair.Value))

	err = yaml.Unmarshal(pair.Value, consulConfig)
	if err != nil {
		log.Debug(err)
		return err
	}
	log.Debug(consulConfig)
	return nil
}

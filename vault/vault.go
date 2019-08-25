// Package vault allows a user interact with vault (retrieving stored secrets)
package vault

import (
	"encoding/json"

	"github.com/berni69/go-archetype/utils"
	vaultapi "github.com/hashicorp/vault/api"
	log "github.com/sirupsen/logrus"
)

// LoadVaultConfig Given a relative path to vault secret, this function will download
// the the secrets from vault and it will fill the structure vaultConfig.
func LoadVaultConfig(path string, vaultConfig interface{}) error {
	log.Info("_______________________________________________________________________________________________________________________")
	log.Infof("Vault Properties ================<< %s/v1/kv/%s >>======================================================",
		utils.GetEnv("VAULT_ADDR", "127.0.0.1"), path)
	log.Info("_______________________________________________________________________________________________________________________")

	config := vaultapi.DefaultConfig()

	va, err := vaultapi.NewClient(config)
	if err != nil {
		log.Debug(err)
		return err
	}

	secrets, err := va.Logical().Read(path)
	if err != nil {
		log.Debug(err)
		return err
	}

	b, _ := json.Marshal(secrets.Data)
	json.Unmarshal(b, vaultConfig)

	return nil
}

package consul

import (
	"crypto/tls"
	"log"
)

func Create_tls_config(crt string, key string, intermca string) tls.Config {

	cert, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}

	return tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
}

func InitConsul() {
	config := Create_tls_config("", "", "")
	config.BuildNameToCertificate()
}

package config

import (
	"log"
	"os"
)

type configSender struct {
	clientIDSender string
}

type configHandler struct {
	clientIDHandler string
	dbName          string
	dbUser          string
	dbPassword      string
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func InitSender() configSender {
	cS := configSender{
		clientIDSender: getEnv("NATS_CLUSTER_ID", "default-cluster-id"),
	}
	log.Printf("Cluster ID: %s\n", cS.clientIDSender)
	return cS
}

func InitHandler() configHandler {
	cH := configHandler{
		clientIDHandler: getEnv("NATS_CLUSTER_ID", "default-cluster-id"),
		dbName:          getEnv("POSTGRES_DB", "default-db"),
		dbUser:          getEnv("POSTGRES_USER", "default-user"),
		dbPassword:      getEnv("POSTGRES_PASSWORD", "default-password"),
	}
	log.Printf("Cluster ID: %s\n", cH.clientIDHandler)
	log.Printf("POSTGRES DB: %s\n", cH.dbName)
	log.Printf("POSTGRES USER: %s\n", cH.dbUser)
	log.Printf("POSTGRES PASSWORD: %s\n", cH.dbPassword)
	return cH
}

func (cS configSender) GetClientID() string {
	return cS.clientIDSender
}

func (cH configHandler) GetClientID() string {
	return cH.clientIDHandler
}

func (cH configHandler) GetDBName() string {
	return cH.dbName
}

func (cH configHandler) GetDBUser() string {
	return cH.dbUser
}

func (cH configHandler) GetDBPassword() string {
	return cH.dbPassword
}

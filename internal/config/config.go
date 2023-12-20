package config

import (
	"log"
	"os"
)

type NATSConfig struct {
	clusterID string
	clientID  string
	subject   string
	URL       string
}

type DBConfig struct {
	dbName     string
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
}

type SenderConfig struct {
	NATSCfg NATSConfig
}

type HandlerConfig struct {
	NATSCfg NATSConfig
	DBCfg   DBConfig
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func InitSender() *SenderConfig {
	cS := &SenderConfig{
		NATSConfig{
			clusterID: getEnv("NATS_CLUSTER_ID", "default-cluster-id"),
			clientID:  "Sender",
			subject:   "subject",
			URL:       "nats://nats-streaming:4222",
		},
	}
	log.Printf("Cluster ID: %s\n", cS.NATSCfg.clusterID)
	log.Printf("Cluster ID: %s\n", cS.NATSCfg.clientID)
	log.Printf("Cluster ID: %s\n", cS.NATSCfg.subject)
	return cS
}

func InitHandler() *HandlerConfig {
	cH := &HandlerConfig{
		NATSConfig{
			clusterID: getEnv("NATS_CLUSTER_ID", "default-cluster-id"),
			clientID:  "Handler",
			subject:   "subject",
			URL:       "nats://nats-streaming:4222",
		},
		DBConfig{
			dbName:     getEnv("POSTGRES_DB", "default-db"),
			dbUser:     getEnv("POSTGRES_USER", "default-user"),
			dbPassword: getEnv("POSTGRES_PASSWORD", "default-password"),
			dbHost:     "db",
			dbPort:     "5432",
		},
	}
	log.Printf("Cluster ID: %s\n", cH.NATSCfg.clusterID)
	log.Printf("Cluster ID: %s\n", cH.NATSCfg.clientID)
	log.Printf("Cluster ID: %s\n", cH.NATSCfg.subject)
	log.Printf("POSTGRES DB: %s\n", cH.DBCfg.dbName)
	log.Printf("POSTGRES USER: %s\n", cH.DBCfg.dbUser)
	log.Printf("POSTGRES PASSWORD: %s\n", cH.DBCfg.dbPassword)
	log.Printf("POSTGRES HOST: %s\n", cH.DBCfg.dbHost)
	log.Printf("POSTGRES PORT: %s\n", cH.DBCfg.dbPort)
	return cH
}

func (cS NATSConfig) GetClientID() string {
	return cS.clientID
}

func (cS NATSConfig) GetClusterID() string {
	return cS.clusterID
}

func (cS NATSConfig) GetSubject() string {
	return cS.subject
}

func (cS NATSConfig) GetUrl() string {
	return cS.URL
}

func (cH DBConfig) GetDBName() string {
	return cH.dbName
}

func (cH DBConfig) GetDBUser() string {
	return cH.dbUser
}

func (cH DBConfig) GetDBPassword() string {
	return cH.dbPassword
}

func (cH DBConfig) GetDBHost() string {
	return cH.dbHost
}

func (cH DBConfig) GetDBPort() string {
	return cH.dbPort
}

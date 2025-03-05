package models

type Config struct {
	DbName    string `json:"dbName"`
	DbUser    string `json:"dbUser"`
	DbTcp     string `json:"dbTcp"`
	SecretKey string `json:"secretKey"`
}
